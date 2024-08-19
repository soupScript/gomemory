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

func Getsum(slice []int) int{
	var placeholder int
	for i:=0;i<len(slice);i++{
		placeholder+=slice[i]
	}
	return placeholder
}

func Shift(slice*[]int, slot1 int, slot2 int) []int{
          (*slice)[slot1]+=(*slice)[slot2]
          (*slice)[slot2]=0
          return *slice
}

func Dealloc(slice*[]int, deallocsize int, ioWarning bool, autoSolution bool,customId string) []int{
	var placeholder string
	if deallocsize%4!=0{
		fmt.Println("Error on "+customId+": Deallocation size not divisible by 4")
		return *slice
	}
	for i:=0;i<len(*slice);i++{
		if i==len(*slice)-1{
			if ioWarning && (*slice)[i]!=0{
				fmt.Println("Warning on "+customId+": Deallocating memory will lower variable value as a value is already signed to that/those memory slot(s), proceed? (Possible solution: b to shift problematic slot to safe slot)")
				fmt.Scanln(&placeholder)
				switch placeholder{
					
				case "y":
					*slice = (*slice)[:len(*slice)-1]
				case "b":
                                        Shift(slice, len(*slice)-2,len(*slice)-1)
                                        *slice = (*slice)[:len(*slice)-1]
				default:
					return *slice
				}
				
			}else if autoSolution && (*slice)[i]!=0{
				Shift(slice, len(*slice)-2, len(*slice)-1)
				*slice = (*slice)[:len(*slice)-1]

			}else if (*slice)[i]!=0{
				*slice = (*slice)[:len(*slice)-1]

			}
		}
	}
	return *slice

}





