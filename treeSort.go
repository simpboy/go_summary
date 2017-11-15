package main

import "fmt"

type tree struct{
	value int
	left,right *tree
}

func main() {
	var slice = []int{1,10,2,9,3,8,5,6,8,-1,-5,-3,}
	slice = Sort(slice)
	fmt.Println(slice)

}

func Sort(values []int) []int{
	var root *tree
	for _,v := range values{
		root = add(root,v)
		fmt.Println(root)
	}
	appendValues(values[:0],root)
	return values
}

func appendValues(values []int,t *tree) []int  {
	if(t!=nil){
		values = appendValues(values[:0],t.left)
		values = append(values, t.value)
		values = appendValues(values[len(values):],t.right)
	}
	return values
}


func add(t *tree,value int) *tree{
	if(t == nil){
		t := new(tree)
		t.value = value
		return t
	}
	if value<t.value{
		t.left = add(t.left,value)
	}else if value>=t.value{
		t.right = add(t.right,value)
	}
	return t
}

