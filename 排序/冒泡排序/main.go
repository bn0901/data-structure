package main

import "fmt"

/*
冒泡排序的三种方法
*/
func main(){
	data := []int{9,1,5,8,3,7,4,6,2}
	fmt.Println(BubbleLow(data))

	data2 := []int{9,1,5,8,3,7,4,6,2}
	fmt.Println(Bubble(data2))

	data3 := []int{9,1,5,8,3,7,4,6,2}
	fmt.Println(BubbleOptimize(data3))

}

// BubbleLow 不算特别规范的简单冒泡
// 做法是：
// 1、第一位数字跟所有其他位置相比较，如果第一位数字大，就互换位置。  第一轮循环保证了第一位是整个数组的最小值。
// 2、第二位数字与 第三位开始的所有位置相比较，如果第二位数字大，就互换位置。。。。
// ....
// 总结：这种方法效率过低。
func BubbleLow(data []int)[]int{
	for i:=0;i<len(data)-1;i++{    // 注意这里i不用等于最后一位。只要是倒数第二位，能跟j比较即可。
		for j:=i+1; j<=len(data)-1;j++{    // 注意：这里要是j<len(data)-1的话，j就不可能等于最后一位数字了，肯定是错的。 所以<=len-1或< len都可以
			if data[i] > data[j]{
				swap(data, i, j)
			}
		}
	}
	return data
}

// Bubble 正规冒泡
// 做法是：
// 1 、第一个位置，要全数组最小。 这个i的作用。  然后j从倒数第二位与倒数第一位进行比较。 依次把最小的数冒泡上来。 结束
// 2、 然后是第二个位置要剩下里面数组里最小， i=1。 同理
func Bubble(data []int)[]int{
	for i:=0;i<len(data)-1;i++{  // 剩下的最后一个位置不需要排
		for j:= len(data)-2; j >=i; j--{
			if data[j] > data[j+1]{
				swap(data,j,j+1)   // 注意：这里不是i和j交换了。 跟i已经没关系了
			}
		}
	}
	return data
}

// BubbleOptimize 冒泡优化
func BubbleOptimize(data []int)[]int{
	flag := true
	for i:=0;i<len(data)-1 &&flag;i++{
		flag=false    // 如果底下一直没修改flag，说明没有排序需求了，那么是可以退出的。
		for j:= len(data)-2; j >=i; j--{
			if data[j] > data[j+1]{
				swap(data,j,j+1)
				flag=true   // 如果有数据互换，说明数组里仍有要排序的需求，继续。
			}
		}
	}
	return data
}


func swap(data []int,i,j int)[]int{
	data[i],data[j] = data[j],data[i]
	return data
}
