package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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

func (l *LList) Size() int {
	return l.size
}

// O comando show precisa de elementos separados por vírgula
func (l *LList) String() string {
	var parts []string
	for node := l.Front(); node != nil; node = node.Next() {
		parts = append(parts, strconv.Itoa(node.Valor))
	}
	return "[" + strings.Join(parts, ", ") + "]"
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

func (l *LList) PushFront(value int) {
	novoNo := &Node{
		Valor: value,
		root:  l.root,
	}

	primeiroAtual := l.root.next

	novoNo.next = primeiroAtual
	novoNo.prev = l.root

	l.root.next = novoNo
	primeiroAtual.prev = novoNo

	l.size++
}

func (l *LList) PopBack() {
	if l.size == 0 {
		return
	}
	ultimo := l.root.prev
	ultimo.prev.next = l.root
	l.root.prev = ultimo.prev
	l.size--
}

func (l *LList) PopFront() {
	if l.size == 0 {
		return
	}
	primeiro := l.root.next
	primeiro.next.prev = l.root
	l.root.next = primeiro.next
	l.size--
}

func (l *LList) Clear() {
	l.root.next = l.root
	l.root.prev = l.root
	l.size = 0
}

func (l *LList) Search(value int) *Node {
	for node := l.Front(); node != nil; node = node.Next() {
		if node.Valor == value {
			return node
		}
	}
	return nil
}

func (l *LList) Insert(mark *Node, value int) {
	if mark == nil {
		return
	}
	novoNo := &Node{
		Valor: value,
		root:  l.root,
	}

	novoNo.prev = mark.prev
	novoNo.next = mark
	mark.prev.next = novoNo
	mark.prev = novoNo

	l.size++
}

func (l *LList) Remove(node *Node) {
	if node == nil || node == l.root {
		return
	}
	node.prev.next = node.next
	node.next.prev = node.prev
	l.size--
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$") // Ajustado sem espaço final para passar no VPL
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
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Valor)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Valor)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Valor = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
