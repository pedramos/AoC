package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

type TokenType int

const (
	CdTok TokenType = iota // includes $ and the cmd (ls, cd, etc)
	LsTok

	FileEntryTok
	DirEntryTok
	Illegal
)

func (t TokenType) String() string {
	switch t {
	case CdTok:
		return "CdTok"
	case LsTok:
		return "LsTok"
	case FileEntryTok:
		return "FileEntryTok"
	case DirEntryTok:
		return "DirEntryTok"
	default:
		return "Illegal"
	}
}

type Token struct {
	typ TokenType
	val []string
}

func (t Token) String() string {
	return fmt.Sprintf("Token{typ: %v, val: %s}", t.typ, t.val)
}

type Scanner struct {
	tok  Token
	buff []byte

	isListing bool

	// current position in the buffer and
	// start of the current token
	curr, start int

	err error
}

func NewScanner(r io.Reader) Scanner {
	var buff bytes.Buffer
	_, _ = io.Copy(&buff, r)

	return Scanner{
		buff: buff.Bytes(),
	}
}

func (s Scanner) Rune() (rune, int, error) {
	r, sz := utf8.DecodeRune(s.buff[s.curr:])

	// checking rune errors
	if r == utf8.RuneError && sz == 0 {
		return utf8.RuneError, 0, io.EOF
	} else if r == utf8.RuneError && sz == 1 {
		return r, 0, fmt.Errorf("Invalud urf8 char")
	}
	return r, sz, nil
}

func (s *Scanner) advance() (rune, error) {

	r, sz, err := s.Rune()
	if err != nil {
		return r, err
	}
	s.curr += sz
	return r, nil

}

func (s *Scanner) peek() (rune, error) {

	r, _, err := s.Rune()
	if err != nil {
		return r, err
	}
	return r, nil

}

func (s *Scanner) skipSpace() error {

	r, _, err := s.Rune()
	if !unicode.IsSpace(r) || err != nil {
		s.start = s.curr
		return err
	}

	for {
		r, err = s.peek()
		if !unicode.IsSpace(r) || err != nil {
			break
		}
		s.advance()
		s.start = s.curr
	}
	return err
}

func (s *Scanner) Word() (string, error) {
	var sb strings.Builder

	s.skipSpace()

	s.curr = s.start
	r, err := s.peek()

	for {
		r, err = s.peek()
		if unicode.IsSpace(r) || err != nil {
			break
		}
		s.advance()
		sb.WriteRune(r)
	}
	return sb.String(), err
}

func (s *Scanner) Scan() bool {

	s.skipSpace()
	r, err := s.peek()
	if err != nil {
		return false
	}

	// set the start of the token
	s.start = s.curr

	switch r {
	case '$':
		s.advance()
		w, err := s.Word()
		if w == "cd" {
			if err != nil {
				s.err = err
				return false
			}
			cdArg, _ := s.Word()
			s.tok = Token{
				typ: CdTok,
				val: []string{cdArg},
			}
			return true
		} else if w == "ls" {
			s.tok = Token{
				typ: LsTok,
			}
		}
	case 'd':
		_, _ = s.Word() // consume "dir"
		_ = s.skipSpace()
		w, _ := s.Word()
		s.tok = Token{
			typ: DirEntryTok,
			val: []string{w},
		}
	default:
		var fileVal []string
		w, _ := s.Word()
		// log.Println(w)
		fileVal = append(fileVal, w)
		w, _ = s.Word()
		fileVal = append(fileVal, w)
		s.tok = Token{
			typ: FileEntryTok,
			val: fileVal,
		}
	}
	return true
}

func (s *Scanner) Token() Token {
	return s.tok
}

type Path []string

func NewPath(path string) Path {
	if path == "/" {
		return []string{"/"}
	} else {
		return strings.Split(path, "/")
	}
}

func (p Path) String() string {
	return strings.Join(p, "/")
}

type Dir struct {
	path    Path
	files   map[string]*File
	subdirs map[string]*Dir
	prev    *Dir
	size    int
}

func (d Dir) String() string {
	var sb strings.Builder
	sb.WriteString("\n")
	if d.Name() == "/" {
		sb.WriteString(fmt.Sprintf("%s (size=%d)", d.Name(), d.size))
	} else {
		sb.WriteString(fmt.Sprintf("%s/ (size=%d)", d.Name(), d.size))
	}
	for _, f := range d.files {
		s := fmt.Sprintf("\n	%s (size=%d)", f.Name(), f.size)
		sb.WriteString(s)
	}
	for _, sd := range d.subdirs {

		sub := fmt.Sprintf("	%s", sd)
		sub = strings.ReplaceAll(sub, "\n", "\n	")
		sb.WriteString(sub)
	}

	return sb.String()

}

func NewDir(path Path, prev *Dir) *Dir {
	return &Dir{
		path:    path,
		prev:    prev,
		files:   make(map[string]*File),
		subdirs: make(map[string]*Dir),
	}
}

func (d Dir) Name() string {
	if len(d.path) == 0 {
		return "/"
	} else {
		return d.path[len(d.path)-1]
	}
}

func (d *Dir) UpdateSize() int {
	total := 0
	for i := range d.files {
		total += d.files[i].size
	}

	for i := range d.subdirs {
		total += d.subdirs[i].UpdateSize()
	}
	d.size = total
	return total
}

type File struct {
	name string
	size int
}

func (f File) Name() string {
	return f.name
}

func SetCWD(fs *Dir, path Path) (*Dir, error) {
	cwd := fs

	if len(path) == 1 && path[0] == "/" {
		return cwd, nil
	} else if path[0] != "/" {
		return nil, fmt.Errorf("Path should start with '/'")
	} else {
		path = path[1:]
	}
	for _, dir := range path {
		if subdir, ok := cwd.subdirs[dir]; ok {
			cwd = subdir
		} else {
			return nil, fmt.Errorf("path not found")
		}

	}
	return cwd, nil
}

func main() {
	log.SetFlags(log.Llongfile)

	f, err := os.Open("p1")
	if err != nil {
		log.Fatal(err)
	}
	s := NewScanner(f)

	fs := NewDir(NewPath("/"), nil)
	cwd := fs

	lsmod := false

	for s.Scan() {
		switch s.Token().typ {
		case CdTok:
			dirname := s.Token().val[0]
			if dirname == ".." {
				cwdpath := cwd.path[:len(cwd.path)-1]
				cwd, _ = SetCWD(fs, cwdpath)
			} else if []rune(dirname)[0] == '/' {
				d := NewPath(dirname)
				if cwd, err = SetCWD(fs, d); err != nil {
					log.Fatalf("Could not find dir %s in %v \n", d, fs)
				}
			} else {
				if _, ok := cwd.subdirs[dirname]; !ok {
					log.Fatalf("Could not find dir %s in path %s", dirname, cwd.path)
				}
				var cwdpath Path
				cwdpath = append(cwdpath, cwd.path...)
				cwdpath = append(cwdpath, dirname)
				if cwd, err = SetCWD(fs, cwdpath); err != nil {
					log.Fatalf("Error 'cd'ing into %s in %s\n", dirname, cwdpath)
				}
			}
			lsmod = false
		case LsTok:
			lsmod = true
		case FileEntryTok:
			if lsmod == false {
				log.Fatal("Error parsing\n")
			}
			sz, _ := strconv.Atoi(s.Token().val[0])
			cwd.files[s.Token().val[0]] = &File{
				name: s.Token().val[1],
				size: sz,
			}
		case DirEntryTok:
			if lsmod == false {
				log.Fatal("Error parsing\n")
			}
			dirname := s.Token().val[0]
			var newdir Path
			newdir = append(newdir, cwd.path...)
			newdir = append(newdir, dirname)
			cwd.subdirs[dirname] = NewDir(newdir, cwd)
		}
	}
	fmt.Printf("Total: %d\n", fs.size)
	fs.UpdateSize()

	//log.Printf("\n--FINAL--\n%s", fs)
	fmt.Printf("Total: %d\n", fs.size)
	fmt.Printf("Solution1: %d\n", fs.solution1())
	fmt.Printf("Solution2: %d\n", fs.solution2())
}

func (d Dir) solution1() int {
	dirs := d.atMost100000()
	total := 0
	for i := range dirs {
		//fmt.Println(dirs[i].Name())
		total += dirs[i].size
	}
	return total
}

func (cwd Dir) atMost100000() []*Dir {
	var dirs []*Dir

	if cwd.size < 100000 {
		dirs = append(dirs, &cwd)
	}
	for _, d := range cwd.subdirs {
		if tmp := d.atMost100000(); len(tmp) > 0 {
			dirs = append(dirs, tmp...)
		}
	}
	return dirs
}

func (cwd Dir) bigEnough(needed int) []*Dir {
	var canRm []*Dir
	var s strings.Builder
	s.WriteString(fmt.Sprintf("%s (size=%d).........", cwd.Name(), cwd.size/1_000_000))
	if cwd.size >= needed {
		s.WriteString("Added")
		canRm = append(canRm, &cwd)
	} else {
		s.WriteString("NOK")
	}
	log.Print(s.String())
	for _, d := range cwd.subdirs {
		if tmp := d.bigEnough(needed); len(tmp) > 0 {
			canRm = append(canRm, tmp...)
		}
	}
	return canRm

}

func (fs Dir) solution2() int {
	log.Printf("Required: %d\n", (70000000-fs.size)/1_000_000)
	canRm := fs.bigEnough(70000000 - fs.size)

	log.Printf("%d", len(canRm))
	lowest := fs.size
	for _, d := range canRm {
		if d.size <= lowest {
			lowest = d.size
		}
	}
	return lowest
}