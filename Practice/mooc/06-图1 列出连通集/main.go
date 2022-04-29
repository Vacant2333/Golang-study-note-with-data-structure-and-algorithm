package main

import "fmt"

/*
https://pintia.cn/problem-sets/1497448825169559552/problems/1510878544933298176
给定一个有N个顶点和E条边地无向图，请用DFS和BFS分别列出其所有的连通集。假设顶点从0到N−1编号。进行搜索时，
假设我们总是从编号最小的顶点出发，按编号递增的顺序访问邻接点。
输入格式:
输入第1行给出2个整数N(0<N≤10)和E，分别是图的顶点数和边数。随后E行，每行给出一条边的两个端点。每行中的数字之间用空格分隔。
输出格式:
每行输出一个连通集。先输出DFS的结果，再输出BFS的结果。
输入样例:
8 6
0 7
0 1
2 0
4 1
2 4
3 5
输出样例:
{ 0 1 4 2 7 }
{ 3 5 }
{ 6 }
{ 0 1 2 7 4 }
{ 3 5 }
{ 6 }
*/

func main() {
	var nodeCount, edgeCount int
	fmt.Scan(&nodeCount, &edgeCount)
	graph := make([][]bool, nodeCount)
	for i, _ := range graph {
		graph[i] = make([]bool, nodeCount)
	}
	var n1, n2 int
	for i := 0; i < edgeCount; i++ {
		fmt.Scan(&n1, &n2)
		graph[n1][n2] = true
		graph[n2][n1] = true
	}
	// DFS
	visited := make(map[int]bool, 0)
	for node := 0; node < nodeCount; node++ {
		if _, ok := visited[node]; ok {
			continue
		}
		fmt.Print("{ ")
		stack := make([]int, 0)
		stack = append(stack, node)
		visited[node] = true
		fmt.Print(node, " ")
		var flag bool
		for len(stack) > 0 {
			// 现在的节点是最后一个节点
			now := stack[len(stack)-1]
			flag = false
			// 检查有没有可以访问的节点
			for i, can := range graph[now] {
				if _, ok := visited[i]; !ok && can {
					// 没访问过
					stack = append(stack, i)
					visited[i] = true
					flag = true
					fmt.Print(i, " ")
					break
				}
			}
			if flag == false {
				// 没有可以访问的节点 回退
				stack = stack[:len(stack)-1]
			}
		}
		fmt.Print("}\n")
	}
	// BFS
	visited = make(map[int]bool, 0)
	for node := 0; node < nodeCount; node++ {
		if _, ok := visited[node]; ok {
			continue
		}
		fmt.Print("{ ")
		queue := make([]int, 0)
		queue = append(queue, node)
		visited[node] = true
		fmt.Print(node, " ")
		index := 0
		for index < len(queue) {
			// 现在的节点是最后一个节点
			now := queue[index]
			// 检查有没有可以访问的节点
			for i, can := range graph[now] {
				if _, ok := visited[i]; !ok && can {
					// 没访问过
					queue = append(queue, i)
					visited[i] = true
					fmt.Print(i, " ")
				}
			}
			index++
		}
		fmt.Print("}\n")
	}
}
