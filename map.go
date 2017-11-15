package main

import (
	"fmt"
	"sort"
)

func main() {
	ages := make(map[string]int)

	ages["bob"] = 12;
	ages["alice"] = 100
	ages["jack"] = 35

	fmt.Printf("%v\n",ages)

	for k,v :=range ages{
		fmt.Printf("%s---%d\n",k,v)
	}


	ages2 := map[string]int{
		"bob":12,
		"alice":100,
	}
	fmt.Println(ages2)

	fmt.Println(len(ages2))

	ages2["alice"]++
	fmt.Println(ages2["alice"])

	delete(ages2,"alice")
	fmt.Println(ages2)

	for name, age := range ages{
		fmt.Printf("%s\t%d\n",name,age)
	}

	//var names []string
	names := make([]string,0,len(ages))
	for name := range ages{
		names = append(names,name)
	}
	sort.Strings(names)
	for _,name := range names{
		fmt.Printf("%s\t%d\n",name,ages[name])
	}
	fmt.Println(ages)

	//map 零值nil
	var ages3 map[string]int
	fmt.Println(ages3 == nil) //map不能比较相等，只有nil例外
	fmt.Println(len(ages) == 0)

	//ages3["alice"] = 12 //panic: assignment to entry in nil map
	ages3 = make(map[string]int)
	ages3["slice"] = 12
	fmt.Println(ages3)

	ages3["john"] = 100
	if age,ok := ages3["john"]; !ok{
		fmt.Println("john not exist!")
	}else{
		fmt.Println(age)
	}

	e := equal(ages3,ages)
	fmt.Println(e)

	e = equal(map[string]int{"A":0},map[string]int{"B":42})
	fmt.Println(e)


	string_receive := fmt.Sprintf("%q", ages3)
	fmt.Println(string_receive)
	fmt.Printf("%q", ages3)

}

func equal(x,y map[string]int) bool{
	if(len(x)!=len(y)){
		return false
	}
	for k,xv := range x{
		if yv,ok:= y[k];!ok||xv!=yv{
			return false
		}
	}
	return true
}


