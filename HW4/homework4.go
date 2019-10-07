package main

import "fmt"

type Node struct {
	Value int
}

// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
	var i int
	for i = len(nodes) - 1; i >= 0; i--{
		down(nodes, i, len(nodes))
	}

}

// 需要down（下沉）的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
	// fmt.Printf(":%d %d %d\n", i, nodes[i].Value, len(nodes))
	if (2 * (i + 1) < len(nodes)){
		if	(nodes[2 * (i + 1)].Value < nodes[2 * (i + 1) - 1].Value){
			if nodes[2 * (i + 1)].Value < nodes[i].Value{
				nodes[2 * (i + 1)],nodes[i] = nodes[i], nodes[2 * (i + 1)]
			}
		}else{
			if nodes[2 * (i + 1) - 1].Value < nodes[i].Value{
				nodes[2 * (i + 1) - 1],nodes[i] = nodes[i], nodes[2 * (i + 1) - 1]
			}
		}
		down(nodes, 2 * (i + 1), n)
		down(nodes, 2 * (i + 1) - 1, n)
	}else{
		if 2 * (i + 1) - 1 < len(nodes) {
			if nodes[2 * (i + 1) - 1].Value < nodes[i].Value{
				nodes[2 * (i + 1) - 1],nodes[i] = nodes[i], nodes[2 * (i + 1) - 1]

			}
			down(nodes, 2 * (i + 1) - 1, n)
		}

	}
}

// 用于保证插入新元素(j为元素的索引,切片末尾插入，堆底插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
	nodes[0], nodes[j] = nodes[j], nodes[0]
	down(nodes, 0, len(nodes))
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	var top Node = nodes[0]
	nodes[0] = nodes[len(nodes) - 1]
	nodes = nodes[:len(nodes) - 1]
	down(nodes, 0, len(nodes))
	return top, nodes
}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes, node)
	up(nodes, len(nodes) - 1)
	return nodes
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
	var i int
	var n Node
	for i, n = range nodes{
		if (n.Value == node.Value){
			nodes[i] = nodes[len(nodes) - 1]
			nodes = nodes[:len(nodes) - 1]
			break
		}
	}
	up(nodes, i)
	down(nodes, i, len(nodes))
	return nodes
}

func main() {
	//test Init
    fmt.Printf("Test Init\nInit with array[1 3 2]\nExpecting:1 3 2\nResult:\n")
	var arr []Node
	Init(arr)
	var n Node
	n.Value = 1;
	arr = Push(n, arr)
	n.Value = 3;
	arr = Push(n, arr)
	n.Value = 2;
	arr = Push(n, arr)
	var i int
	for i, n = range arr{
		fmt.Printf("%d %d\n",i , n.Value)
	}
	fmt.Printf("\n")

	//test push
    fmt.Printf("Test Push\nPush 0 into heap\nExpecting:0 1 2 3\nResult:\n")
	n.Value = 0;
	arr = Push(n, arr)
	for i, n = range arr{
		fmt.Printf("%d %d\n",i , n.Value)
	}
    fmt.Printf("\n")

	//test pop
    fmt.Printf("Test Pop\nPop 0 out of heap\nExpecting:1 3 2\nResult:\n")
	n, arr = Pop(arr)
	for i, n = range arr{
		fmt.Printf("%d %d\n",i , n.Value)
	}
    fmt.Printf("\n")

    //test remove
    fmt.Printf("Test Remove\nRemove the first node whose value equal to 1\nExpecting:2 3\nResult:\n")
	n.Value = 1;
	arr = Remove(arr, n)
	for i, n = range arr{
		fmt.Printf("%d %d\n",i , n.Value)
	}
    fmt.Printf("\n")
}