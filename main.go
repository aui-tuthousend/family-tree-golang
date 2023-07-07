package main

import "fmt"

type Node struct {
	nama     string
	pasangan string
	status   bool
	left     *Node
	mid      *Node
	right    *Node
	par      *Node
	next     *Node
	prev     *Node
}

func input(n *Node) {
	var namas string
	if n.right != nil {
		fmt.Println("Anak Beliau sudah tiga :D")
	} else {
		fmt.Print("Masukkan nama anak: ")
		fmt.Scanln(&namas)
		nod := &Node{nama: namas}
		if n.left == nil {
			n.left = nod
			nod.par = n
		} else if n.mid == nil {
			n.mid = nod
			nod.par = n
			n.left.next = nod
			nod.prev = n.left
		} else {
			n.right = nod
			nod.par = n
			n.mid.next = nod
			nod.prev = n.mid
		}
	}
}

func printAnak(n *Node) {
	if n.left == nil {
		fmt.Println("Belum punya anak :D")
	} else {
		fmt.Print(n.left.nama + " ")
		if n.mid != nil {
			fmt.Println(n.mid.nama + " ")
		}
		if n.right != nil {
			fmt.Println(n.right.nama + " \n")
		}
	}
}

func printDulur(n *Node) {
	temp := n.par
	if n.nama == temp.left.nama {
		for n != nil {
			fmt.Print(n.nama + " ")
			n = n.next
		}
	} else if n.nama == temp.mid.nama {
		fmt.Print(n.prev.nama + " ")
		fmt.Print(n.next.nama)
	} else {
		for n != nil {
			fmt.Print(n.nama + " ")
			n = n.prev
		}
	}
}

func inputPasangan(n *Node) {
	var np string
	if n.status == false {
		fmt.Print("Masukkan nama pasangan: ")
		fmt.Scanln(&np)
		n.pasangan = np
		n.status = true
	} else {
		fmt.Println("Beliau sudah punya pasangan :D")
	}
}

func search(root *Node, val string) *Node {
	if root == nil {
		return root
	}

	if root.nama == val || root.pasangan == val {
		return root
	}
	leftResult := search(root.left, val)
	if leftResult != nil {
		return leftResult
	}

	midResult := search(root.mid, val)
	if midResult != nil {
		return midResult
	}

	rigtResult := search(root.right, val)
	if rigtResult != nil {
		return rigtResult
	}
	return nil
}

func pr(nn *Node) {
	fmt.Print(" -> ")
	for nn != nil {
		fmt.Print(nn.nama + " ")
		nn = nn.next
	}
}

func printAll(n *Node, depth int) {
	if n == nil {
		return
	}
	fmt.Printf("%*s%s\n", depth*4, "", n.nama)
	printAll(n.left, depth+1)
	printAll(n.mid, depth+1)
	printAll(n.right, depth+1)
}

//func printAll(root *Node) {
//	fmt.Print(root.nama)
//	for root != nil {
//		nn := root.left
//		fmt.Print(" -> ")
//		for nn != nil {
//			fmt.Print(nn.nama + " ")
//			nn = nn.next
//		}
//		root = root.left
//	}
//
//	ro := root.mid
//	for ro != nil {
//		nn := root.left
//		fmt.Print(" -> ")
//		for nn != nil {
//			fmt.Print(nn.nama + " ")
//			nn = nn.next
//		}
//		root = root.left
//	}
//
//}

func menu() {
	fmt.Print(
		"\nTindakan: " +
			"\n[1]. Input pasangan" +
			"\n[2]. Input anak" +
			"\n[3]. Daftar anak" +
			"\n[4]. Daftar Saudara" +
			"\n[5]. Siapa Parentku" +
			"\n[6]. Siapa Kakek-ku" +
			"\n[7]. Siapa Kakek Buyut-ku" +
			"\nMasukkan opsi: ")
}

func main() {
	tree := &Node{nama: "adam", pasangan: "eve", status: true}
	var lanjut string
	//newNode := &Node{}
	for {
		var tem string
		var opsi int
		fmt.Print("\nMasukkan nama orang atau ketik 1 untuk print all: ")
		fmt.Scanln(&tem)
		if tem == "1" {
			printAll(tree, 0)
			continue
		}
		newNode := search(tree, tem)
		if newNode == nil {
			fmt.Println("orang tak ditemukan :( ")
			continue
		}
		fmt.Print("Halo Tuan ")
		fmt.Print(newNode.nama)
		if newNode.status == true {
			fmt.Print(" dan Nyonya ")
			fmt.Print(newNode.pasangan)
		}
		menu()
		fmt.Scanln(&opsi)
		switch opsi {
		case 1:
			inputPasangan(newNode)
		case 2:
			if newNode.status != true {
				fmt.Println("Beliau belum punya pasangan T_T")
			} else {
				input(newNode)
			}
		case 3:
			printAnak(newNode)
		case 4:
			printDulur(newNode)
		case 5:
			fmt.Print("Anda adalah anak dari Tuan " + newNode.par.nama + " Dan Nyonya " + newNode.par.pasangan + "\n")
		case 6:
			kak := newNode.par.par
			if kak != nil {
				fmt.Print("Anda adalah Cucu dari Tuan " + kak.nama + " Dan Nyonya " + kak.pasangan + "\n")
			} else {
				fmt.Println("Anda belum punya kakek || kakek anda masih belum diketahui :D")
			}
		case 7:
			kak := newNode.par.par.par
			if kak != nil {
				fmt.Print("Anda adalah Buyut dari Tuan " + kak.nama + " Dan Nyonya " + kak.pasangan + "\n")
			} else {
				fmt.Println("Anda belum punya kakek buyut || kakek buyut anda masih belum diketahui :D")
			}
		}

		fmt.Print("lanjut? : ")
		fmt.Scanln(&lanjut)
		if lanjut != "y" {
			break
		}
	}
}
