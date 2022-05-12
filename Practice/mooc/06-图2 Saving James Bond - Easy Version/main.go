package main

import (
	"fmt"
	"math"
)

/*
This time let us consider the situation in the movie "Live and Let Die" in which James Bond,
the world's most famous spy, was captured by a group of drug dealers. He was sent to a small piece
of land at the center of a lake filled with crocodiles. There he performed the most daring action
to escape -- he jumped onto the head of the nearest crocodile! Before the animal realized what was
happening, James jumped again onto the next big head... Finally, he reached the bank before the last
crocodile could bite him (actually the stuntman was caught by the big mouth and barely escaped with his extra thick boot).
Assume that the lake is a 100 by 100 square one. Assume that the center of the lake is at (0,0)
and the northeast corner at (50,50). The central island is a disk centered at (0,0) with the
diameter of 15. A number of crocodiles are in the lake at various positions. Given the coordinates
of each crocodile and the distance that James could jump, you must tell him whether he can escape.
Input Specification:
Each input file contains one test case. Each case starts with a line containing two positive integers
N (≤100), the number of crocodiles, and D, the maximum distance that James could jump. Then N lines
follow, each containing the (x,y) location of a crocodile. Note that no two crocodiles are staying at the same position.
Output Specification:
For each test case, print in a line "Yes" if James can escape, or "No" if not.
Sample Input 1:
14 20
25 -15
-25 28
8 49
29 15
-35 -2
5 28
27 -29
-8 -28
-20 -35
-25 -20
-13 29
-30 15
-35 40
12 12
Sample Output 1:
Yes
Sample Input 2:
4 13
-12 12
12 12
-12 -12
12 -12
Sample Output 2:
No
*/

func main() {
	// 鳄鱼数量, 可以跳跃的距离
	var num int
	var distance float64
	fmt.Scan(&num, &distance)
	// 生成图 graph[x][y] 存鳄鱼坐标
	graph := make([][2]int, 0)
	// 读入鳄鱼(转换一下坐标系,中心岛在50,50)
	var x, y int
	for i := 0; i < num; i++ {
		fmt.Scan(&x, &y)
		graph = append(graph, [2]int{x + 50, y + 50})
	}
	// 路径栈
	routeStack := make([][2]int, 0)
	// 访问过的节点
	reachedNode := make([][2]int, 0)
	// 中心岛,直径是15
	routeStack = append(routeStack, [2]int{50, 50})
	for len(routeStack) != 0 {
		// 取最新的节点作为当前节点
		nowPos := routeStack[len(routeStack)-1]
		// 判断能不能上岸
		if canReachBrink(nowPos, distance) {
			// 可以上岸
			fmt.Println("Yes")
			return
		}
		// 不能上岸 判断能不能向前走
		var canGo [][2]int
		if nowPos[0] == 50 && nowPos[1] == 50 {
			// 如果在中心岛, 加跳跃距离 15/2
			canGo = getCanGoPos(nowPos, distance+7.5, reachedNode, graph)
		} else {
			// 不是中心岛
			canGo = getCanGoPos(nowPos, distance, reachedNode, graph)
		}
		if len(canGo) == 0 {
			// 没地方可以去 后退
			routeStack = routeStack[0 : len(routeStack)-1]
			reachedNode = append(reachedNode, nowPos)
		} else {
			// 有地方可以向前 继续
			routeStack = append(routeStack, canGo[0])
		}
	}
	fmt.Println("No")
}

// 判断能否到达岸边
func canReachBrink(pos [2]int, distance float64) bool {
	// [0, 0-99]
	for y := 0; y < 100; y++ {
		if getDistance(pos, [2]int{0, y}) <= distance || getDistance(pos, [2]int{99, y}) <= distance {
			return true
		}
	}
	for x := 0; x < 100; x++ {
		if getDistance(pos, [2]int{x, 0}) <= distance || getDistance(pos, [2]int{x, 99}) <= distance {
			return true
		}
	}

	return false
}

// 获得可以去的鳄鱼
func getCanGoPos(nowPos [2]int, distance float64, reachedNode [][2]int, graph [][2]int) [][2]int {
	canGo := make([][2]int, 0)
	for _, pos := range graph {
		// 检查是否去过这个鳄鱼
		if isReached(pos, reachedNode) == true || pos == nowPos {
			// 去过这个鳄鱼 跳过
			continue
		} else if getDistance(nowPos, pos) <= distance {
			// 没去过 加入到canGo, 可以到达
			canGo = append(canGo, pos)
		}
	}
	return canGo
}

// 检查是否去过这个鳄鱼
func isReached(nowPos [2]int, reachedNode [][2]int) bool {
	for _, pos := range reachedNode {
		if pos == nowPos {
			return true
		}
	}
	return false
}

// 获得两点之间的距离
func getDistance(pos1, pos2 [2]int) float64 {
	return math.Sqrt(float64((pos1[0]-pos2[0])*(pos1[0]-pos2[0]) + (pos1[1]-pos2[1])*(pos1[1]-pos2[1])))
}
