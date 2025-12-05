package main

import "fmt"

type Queue struct {
	s1 []int
	s2 []int
}

func (q *Queue) Push(num int) {
	q.s1 = append(q.s1, num)
}

func (q *Queue) Remove() int {
	if len(q.s2) == 0 {
		for len(q.s1) > 0 {
			top := len(q.s1) - 1
			q.s2 = append(q.s2, q.s1[top])
			q.s1 = q.s1[:top]
		}
	}
	if len(q.s2) == 0 {
		return 0
	}
	n := len(q.s2) - 1
	val := q.s2[n]
	q.s2 = q.s2[:n]
	return val
}
func (q *Queue) Peek() int {
	if len(q.s2) == 0 {
		for len(q.s1) > 0 {
			top := len(q.s1) - 1
			q.s2 = append(q.s2, q.s1[top])
			q.s1 = q.s1[:top]
		}
	}
	if len(q.s2) == 0 {
		return 0
	}
	val := q.s2[len(q.s2)-1]
	return val
}

func main() {
	q := &Queue{}
	q.Push(10)
	q.Push(20)
	q.Push(30)
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	fmt.Println(q.Remove())
	q.Push(40)
	fmt.Println(q.Remove())
	fmt.Println(q.Peek())
}
