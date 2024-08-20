package gomemory

import (
"fmt"
//"math"
)
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

//-128-127 int

func Optimize(slice*[]int) []int{
	for i:=0;i<len(*slice)-1;i++{
		if (*slice)[i]!=0 && i!=0{
			if (*slice)[0]+(*slice)[i]<127 && (*slice)[0]+(*slice)[i]>-128{
				(*slice)[0]+=(*slice)[i]
				(*slice)[i]=0
			}
		}
	}
	var placeholder int
	for i:=0;i<len(*slice)-1;i++{
		if (*slice)[i]==0{
			placeholder=i
			break
		}

	}
 	*slice=(*slice)[:(len(*slice)-(len(*slice)-placeholder))]
	return *slice
}

func Realloc(slice*[]int, size int, customId string) []int{
	if size%4!=0{
		fmt.Println("Error at "+customId+": Desired size not divisible by 4")
		return *slice
	}else if size==len(*slice)*4{
		fmt.Println("Error at "+customId+": Reallocation size and current size are the same. If you want to optimize the memory, use Optimize(slice*[]int)")
		return *slice
	}
	if size>len(*slice)*4{
		for i:=0;i<(size-len(*slice)*4)/4;i++{
			*slice=append(*slice, 0)
		}
		return *slice
	}
	if size<len(*slice)*4{
		Optimize(slice)
		if len(*slice)*4!=size{
                	for i:=0;i<(size-len(*slice)*4)/4;i++{                                                                                                                       
                        	*slice=append(*slice, 0)                                                                                                                             
                	}      
		}
		if len(*slice)*4!=size{
			fmt.Println("Note at "+customId+": Size did not reach realloc in end result most likely due to optimize not having enough space to compress slots")
		}
		return *slice

	}
	return *slice
}

func Add(slice*[]int, num int) []int{
	if num+(*slice)[len(*slice)-1]<127 && num+(*slice)[len(*slice)-1]>-128{
		(*slice)[len(*slice)-1]+=num
		return *slice
	}else{
		outScopeBreak:=false
		numref:=num
		negative:=false
		if numref<0{negative=true}
		for{
			if !negative{
				*slice=append(*slice, 0)
				for i:=0;i<127;i++{
					(*slice)[len(*slice)-1]+=1
					numref-=1
				
					if numref==0{
						outScopeBreak=true
						break
					}
				}
			}else{
				//
				*slice =append(*slice, 0)
                                for i:=0;i>-128;i-=1{
                                        (*slice)[len(*slice)-1]-=1
                                        numref+=1
                                                          
                                        if numref==0{     
                                                outScopeBreak=true
                                                break     
                                        }
                                } 

			
		
			}
			if outScopeBreak{
				break
			}
		}
	}
	return *slice
}



