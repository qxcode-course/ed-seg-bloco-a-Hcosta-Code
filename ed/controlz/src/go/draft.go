package main

import (
	"container/list"
	"fmt"
)

func main() {
	var entrada string
	fmt.Scan(&entrada)

	l := list.New()
	l.PushFront('#') 
	cursor := l.Front()

	for _, char := range entrada {
		if char == 'R' {
			l.InsertAfter('\n', cursor)
			cursor = cursor.Next()
			
		} else if char == 'B' {
			if cursor != l.Front() { 
				prev := cursor.Prev()
				l.Remove(cursor)
				cursor = prev
			}
			
		} else if char == 'D' {
			if cursor.Next() != nil {
				l.Remove(cursor.Next())
			}
			
		} else if char == '<' {
			if cursor != l.Front() { 
				cursor = cursor.Prev()
			}
			
		} else if char == '>' {

			if cursor.Next() != nil {
				cursor = cursor.Next()
			}
			
		} else {
			l.InsertAfter(char, cursor)
			cursor = cursor.Next()
		}
	}

	texto := ""

	for it := l.Front().Next(); it != nil; it = it.Next() {
		texto += string(it.Value.(rune))
		
		if it == cursor {
			texto += "|"
		}
	}

	fmt.Println(texto)
}