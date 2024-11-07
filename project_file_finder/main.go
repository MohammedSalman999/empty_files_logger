package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Provide a Directory")
		return
	}

	files, err := os.ReadDir(args[0])
	if err !=nil{
		fmt.Println(err)
		return
	}

	var total int 

	for _ ,file := range files {
		info,err := file.Info()
		if err != nil {
			fmt.Println(err)
			return
		}
		if info.Size() == 0 {
			total += len(info.Name()) +1 
		}
	}

	fmt.Printf("total required space :  %d bytes \n",total)
	
	var names []byte 
	for _, file := range files{
		info,err := file.Info()
		if err!= nil {
			fmt.Println(err)
			return
		}
		if info.Size() == 0 {
			name := info.Name()
			names = append(names, name... )
			names = append(names, '\n')
			
		}
	}

	fmt.Printf("%s",names)
	err = os.WriteFile("out.txt",names,0644)
	if err != nil {
		fmt.Println(err)
		return
	}

}