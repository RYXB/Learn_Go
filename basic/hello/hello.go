package main

import "fmt"

func main()  {
	s := "good bye"
	fmt.Printf("Here is the string s:%s\n", s)
	fmt.Printf("Here is the string s:%s\n", &s)
	var p *string = &s  /* &s是 s的地址给了p *p就是s了 */
	*p = "hello"
	fmt.Printf("Here is the pointer p:%p\n", p)
	fmt.Printf("Here is the string *p:%s\n", *p)
	fmt.Printf("Here is the string s:%s\n", s)
}