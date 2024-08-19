package gomemory

import ("fmt")

func Malloc(slice*[]int, allocsize int, customId string){
	if allocsize%4!=0{
		fmt.Println("Error in "+customId+": Allocation size not a multiple of 4!")
		return
	}
	for i:=0;i<allocsize/4;i++{
		*slice=append(*slice, 0)
	}
}

func Checksize(elementSize int,slice []int) int{
	return elementSize*len(slice)
}


