	for i:=0;  i<len(line);i++{
		fmt.Printf("%c\n",line[i])
	}

	kebalik:=len(line)
	for j := kebalik; j >0; j-- {
		fmt.Printf("%c\n",line[j-1])
	}