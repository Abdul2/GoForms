package main

import "fmt"

type Asset struct {

	Location string
	Titledeeds string
	Type string

}

type Unit struct  {

	TitleDeeds string
	Unitnumber string
	Usage string
	Tenant string
	Contractref string
	Vacant bool
}


type Addnew interface {

	Addnew()

}



func main()  {

	fmt.Println("test")
}