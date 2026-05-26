package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Valor int
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	sentinel := &Node{}

	sentinel.next = sentinel
	sentinel.prev = sentinel

	lista := &LList{
		root: sentinel,
		size: 0,
	}

	return lista
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next

}

func (x *Node) Prev() *Node {
	if x.prev == x.root {
		return nil
	}
	return x.prev
}

func (l *LList) Front() *Node {
	if l.root.next == l.root {
		return nil
	}
	return l.root.next
}

func (l *LList) Back() *Node {
	if l.root.prev == l.root {
		return nil
	}
	return l.root.prev
}

func (l *LList) PushBack(value int) {
	novoNo := &Node{
		Valor: value,
		root:  l.root,
	}

	ultimoAtual := l.root.prev

	novoNo.prev = ultimoAtual
	novoNo.next = l.root

	ultimoAtual.next = novoNo
	l.root.prev = novoNo

	l.size++
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	// 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			// fmt.Println(ll.String())
		case "size":
			// fmt.Println(ll.Size())
		case "push_back":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushBack(num)
			// }
		case "push_front":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushFront(num)
			// }
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			// ll.Clear()
		case "walk":
			// fmt.Print("[ ")
			// for node := ll.Front(); node != nil; node = node.Next() {
			// 	fmt.Printf("%v ", node.Value)
			// }
			// fmt.Print("]\n[ ")
			// for node := ll.Back(); node != nil; node = node.Prev() {
			// 	fmt.Printf("%v ", node.Value)
			// }
			// fmt.Println("]")
		case "replace":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	node.Value = newvalue
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "insert":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Insert(node, newvalue)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "remove":
			// oldvalue, _ := strconv.Atoi(args[1])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Remove(node)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
