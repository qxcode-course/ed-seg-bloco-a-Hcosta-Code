package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	val  int
	next *Node
	prev *Node
}

type LList struct {
	root *Node
	size int
}

// Inicializa a lista criando o nó sentinela que aponta para si mesmo
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

// Retorna a representação da lista no formato [1, 2, 3]
func (ll *LList) String() string {
	var elementos []string
	node := ll.root.next

	for node != ll.root {
		elementos = append(elementos, strconv.Itoa(node.val))
		node = node.next
	}

	return "[" + strings.Join(elementos, ", ") + "]"
}

// Retorna a quantidade de elementos válidos na lista
func (ll *LList) Size() int {
	return ll.size
}

// Insere um novo nó no início da lista (logo após o sentinela)
func (ll *LList) PushFront(val int) {
	novoNo := &Node{val: val}
	primeiroAtual := ll.root.next

	novoNo.next = primeiroAtual
	novoNo.prev = ll.root

	primeiroAtual.prev = novoNo
	ll.root.next = novoNo

	ll.size++
}

// Insere um novo nó no fim da lista (logo antes do sentinela)
func (ll *LList) PushBack(val int) {
	novoNo := &Node{val: val}
	ultimoAtual := ll.root.prev

	// Configura as conexões do novo nó
	novoNo.next = ll.root
	novoNo.prev = ultimoAtual

	// Atualiza os nós vizinhos para apontarem para o novo nó
	ultimoAtual.next = novoNo
	ll.root.prev = novoNo

	ll.size++
}

// Remove o primeiro elemento válido da lista (se existir)
func (ll *LList) PopFront() {
	if ll.size == 0 {
		return
	}

	noParaRemover := ll.root.next
	proximoNo := noParaRemover.next

	// Isola o nó removido ajustando as pontas
	ll.root.next = proximoNo
	proximoNo.prev = ll.root

	ll.size--
}

// Remove o último elemento válido da lista (se existir)
func (ll *LList) PopBack() {
	if ll.size == 0 {
		return
	}

	noParaRemover := ll.root.prev
	anteriorNo := noParaRemover.prev

	// Isola o nó removido ajustando as pontas
	anteriorNo.next = ll.root
	ll.root.prev = anteriorNo

	ll.size--
}

// Esvazia a lista, voltando ao estado inicial onde o sentinela aponta para si mesmo
func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()

		// AJUSTE FINAL: Adicionamos o "$" grudado com o comando lido
		fmt.Println("$" + line)

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
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
