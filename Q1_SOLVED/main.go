package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

func main() {
	content, err := os.ReadFile("./hard.json")
	if err != nil {
		log.Fatal("Error when opening file: ", err)
	}
	var payload [][]int
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	// ex := [][]int{
	// 	{59},
	// 	{73, 41},
	// 	{52, 40, 53},
	// 	{26, 53, 6, 34},
	// }
	maximumSum := maxSumPath(payload)
	fmt.Println(maximumSum)
}

func maxSumPath(tree [][]int) int {
	for i := 1; i < len(tree); i++ {
		for j := 0; j < len(tree[i]); j++ {
			if j == 0 {
				// รวม node ซ้ายสุด
				tree[i][j] += tree[i-1][j]
			} else if j == len(tree[i])-1 {
				// รวม node ขวาสุด
				tree[i][j] += tree[i-1][j-1]
			} else {
				// current level และเป็นพวกที่อยู่ตรงกลาง จะหาว่า parent ของมัน ซ้ายหรือขวาที่มากกว่ากัน
				tree[i][j] += max(tree[i-1][j-1], tree[i-1][j])
			}
		}
		// fmt.Println(tree)
	}

	// find maxsum of deepest level of root
	maxSum := tree[len(tree)-1][0]
	for _, val := range tree[len(tree)-1] {
		if val > maxSum {
			maxSum = val
		}
	}
	return maxSum
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
