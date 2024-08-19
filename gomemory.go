package gomemory

import ("fmt")

func Malloc(slice*[]int, allocsize int, customId string) []int{
	if allocsize%4!=0{
		fmt.Println("Error in "+customId+": Allocation size not a multiple of 4!")
		return nil
	}
	for i:=0;i<allocsize/4;i++{
		*slice=append(*slice, 0)
	}
	return *slice

}

func Checksize(elementSize int,slice []int) int{
	return elementSize*len(slice)
}

func getSum(slice []int) int{
	var placeholder int
	for i:=0;i<len(slice);i++{
		placeholder+=slice[i]
	}
	return placeholder
}

func Dealloc(slice*[]int, deallocsize int, ioWarning bool,customId string) []int{
	var placeholder string
	if deallocsize%4!=0{
		fmt.Println("Error on "+customId+": Deallocation size not divisible by 4")
		return *slice
	}
	if len(*slice)==1 && ioWarning{
		fmt.Println("Warning on "+ customId+": Only 1 memory slot. Empty it?")
		fmt.Scanln(&placeholder)
		if placeholder!="y"{
			return *slice
		}
		

	}
	for i:=0;i<len(*slice);i++{
		if i==len(*slice)-1{
			if ioWarning && (*slice)[i]!=0{
				fmt.Println("Warning on "+customId+": Deallocating memory will lower variable value as a value is already signed to that/those memory slot(s), proceed?")
				fmt.Scanln(&placeholder)
				if placeholder=="y"{
					*slice = (*slice)[:len(*slice)-1]
				}else{
					return *slice
				}
			}else{
				 *slice = (*slice)[:len(*slice)-1]

			}
		}
	}
	return *slice

}



