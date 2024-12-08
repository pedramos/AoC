#!/bin/sh

. $HOME/lib/aoc-session

curl 'https://adventofcode.com/2024/day/$1/input'  --cookie "session=$SESSION"
