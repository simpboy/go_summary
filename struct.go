package main

import (
	"time"
	"fmt"
)

func main() {
	type Employee struct {
		ID 		int
		Name 	string
		Address string
		DoB 	time.Time
		Position	string
		Salary		int
		ManagerID	int
	}
	var dilbert Employee
	dilbert = Employee{
		ID : 1,
		Name:"john",
		Address:"zhangtao",
		DoB:time.Now(),
		Position:"progamer",
		Salary:10,
		ManagerID:19,
	}
	dilbert.Salary *= 2
	fmt.Printf("%v",dilbert);

	t := time.Now()
	ts := t.Unix()
	tf := t.Format("2006-01-02 15:04:05")
	fmt.Printf("%v,%v,%v",t,ts,tf)
	fmt.Println("")
	fmt.Println(tf)

	position := &dilbert.Position
	*position = "Senior " + *position

	fmt.Println(dilbert)

	var employeeOfTheMonth *Employee = &dilbert
	employeeOfTheMonth.Position +=  " (proactive team player)"
	fmt.Println(dilbert)

	type Point struct{ X, Y int }
	//1.1
	p := Point{1,2}
	fmt.Println(p)

	pig := Point{1,2}
	fmt.Println(pig)
	//1.2
	pp := &Point{1,3}
	fmt.Println(pp)

	//2、
	var point Point
	point.X = 1
	point.Y = 2

	//3、
	ppp := new(Point)
	*ppp = Point{1, 2}

	p1 := Point{1, 2}
	q1 := Point{2, 1}

	fmt.Println(p1 == q1);


	type address struct {
		hostname string
		port     int
	}
	hits := make(map[address]int)
	hits[address{"www.baidu.com",443}]++


	type Circle struct {
		X,Y,Radius int
	}
	type Wheel struct {
		X,Y,Radius,Spokes int
	}

	type Point_s struct{
		X,Y int
	}
	type Circle_s struct {
		Point_s
		Radius int
	}
	type Wheel_s struct {
		Circle_s
		Spokes int
	}

	var wheel Wheel_s
	wheel.X = 1
	wheel.Y = 2
	wheel.Radius = 10
	wheel.Spokes = 25

	fmt.Println(wheel)

	wheel = Wheel_s{Circle_s{Point_s{10,25},20},35}
	fmt.Println(wheel)






}
