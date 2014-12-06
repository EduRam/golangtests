package main

import (
	"fmt"
)

func byval(q *int) {
	fmt.Printf("[func] q *int       | type=%T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
	*q = 4143
	fmt.Printf("[func] *q = 4143    | type=%T: &q=%p q=&i=%p  *q=i=%v\n", q, &q, q, *q)
	q = nil
}

func main() {
	i := int(42)
	fmt.Printf("[main] i := int(42) | type=%T: &i=%p i=%v\n", i, &i, i)
	p := &i
	fmt.Printf("[main] p := &i      | type=%T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
	byval(p)
	fmt.Printf("[main] p %T: &p=%p p=&i=%p  *p=i=%v\n", p, &p, p, *p)
	fmt.Printf("[main] i %T: &i=%p i=%v\n", i, &i, i)
}
