awk '
BEGIN	{
			finalMax=0
			finalElf=1 
			elf=1
			max=0
		}
/^$/ 	{
			if(finalMax<max) {finalElf=elf; finalMax=max}
			elf++
			max=0
			next
		}
// 		{
			max += $1
		}
END		{
			printf("%d: %d\n", finalElf, finalMax)
		}
'

