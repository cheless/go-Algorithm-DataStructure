package main

import (
	"DataStructure-described-by-Go/skiplist"
)

func main() {
	sl := skiplist.NewSkiplist()
	for i := 0; i < 20; i++ {
		sl.Add(i)
	}
	sl.PrintStruct()
	for i := 1; i < 20; i++ {
		sl.Delete(20-i)
		sl.PrintStruct()
	}
}