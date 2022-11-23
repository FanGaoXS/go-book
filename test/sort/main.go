package main

import (
	"fmt"
	"sort"
)

func main() {
	ps := Peoples{
		&People{Name: "test1", Age: 30},
		&People{Name: "test2", Age: 20},
		&People{Name: "test3", Age: 45},
	}
	sort.Sort(ps)

	for _, p := range ps {
		fmt.Println(p)
	}
}

type People struct {
	Name string
	Age  int
}

type Peoples []*People

func (ps Peoples) Len() int {
	return len(ps)
}

func (ps Peoples) Less(i, j int) bool {
	return ps[i].Age < ps[j].Age
}

func (ps Peoples) Swap(i, j int) {
	temp := People{
		Name: ps[i].Name,
		Age:  ps[i].Age,
	}

	ps[i].Name, ps[i].Age = ps[j].Name, ps[j].Age
	ps[j].Name, ps[j].Age = temp.Name, temp.Age
}
