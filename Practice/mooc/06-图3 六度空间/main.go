package main

import "fmt"

/*
“六度空间”理论又称作“六度分隔（Six Degrees of Separation）”理论。这个理论可以通俗地阐述为：
“你和任何一个陌生人之间所间隔的人不会超过六个，也就是说，最多通过五个人你就能够认识任何一个陌生人。”如图1所示。
图1 六度空间示意图
“六度空间”理论虽然得到广泛的认同，并且正在得到越来越多的应用。但是数十年来，
试图验证这个理论始终是许多社会学家努力追求的目标。然而由于历史的原因，这样的研究具有太大的局限性和困难。
随着当代人的联络主要依赖于电话、短信、微信以及因特网上即时通信等工具，
能够体现社交网络关系的一手数据已经逐渐使得“六度空间”理论的验证成为可能。
假如给你一个社交网络图，请你对每个节点计算符合“六度空间”理论的结点占结点总数的百分比。
输入格式:
输入第1行给出两个正整数，分别表示社交网络图的结点数N、边数M（≤33×N，表示社交关系数）。
随后的M行对应M条边，每行给出一对正整数，分别是该条边直接连通的两个结点的编号（节点从1到N编号）。
输出格式:
对每个结点输出与该结点距离不超过6的结点数占结点总数的百分比，精确到小数点后2位。每个结节点输出一行，格式为“结点编号:（空格）百分比%”。
输入样例:
10 9
1 2
2 3
3 4
4 5
5 6
6 7
7 8
8 9
9 10
输出样例:
1: 70.00%
2: 80.00%
3: 90.00%
4: 100.00%
5: 100.00%
6: 100.00%
7: 100.00%
8: 90.00%
9: 80.00%
10: 70.00%
*/

func main() {
	var nodeCount, edgeCount int
	fmt.Scan(&nodeCount, &edgeCount)
	// 初始化图
	graph := make([][]bool, nodeCount)
	for i, _ := range graph {
		graph[i] = make([]bool, nodeCount)
	}
	// 读入边
	var t1, t2 int
	for i := 0; i < edgeCount; i++ {
		fmt.Scan(&t1, &t2)
		graph[t1-1][t2-1] = true
		graph[t2-1][t1-1] = true
	}
	// 计算每个边
	for node := 0; node < nodeCount; node++ {
		// 计算这个边的六度比例(BFS)
		// 存放访问到的所有节点名(题目的名称-1 从0开始)
		queue := make([]int, 0, 1000)
		// 能够访问的节点
		visited := make(map[int]bool, 1000)
		visited[node] = true
		queue = append(queue, node)
		// 这一层在queue中的开始节点的下标
		startIndex := 0
		for layer := 1; layer <= 6; layer++ {
			if len(queue) == nodeCount {
				break
			}
			// 遍历6层 把遍历到的节点加入到queue中
			// 下一层的开始下标
			nextStartIndex := len(queue)
			for index := startIndex; index < nextStartIndex; index++ {
				// 遍历node这一层的节点,把可以访问的节点加入到queue中
				for sonNode, can := range graph[queue[index]] {
					if _, ok := visited[sonNode]; !ok && can {
						// 之前没存过这个节点 存入queue和map
						queue = append(queue, sonNode)
						visited[sonNode] = true
					}
				}
			}
			// 更新下一层
			startIndex = nextStartIndex
		}
		fmt.Printf("%v: %.2f%%\n", node+1, float64(len(queue))/float64(nodeCount)*100)
	}
}
