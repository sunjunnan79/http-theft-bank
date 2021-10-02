package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func checkFileIsExist(filename string) bool {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return false
	}
	return true
}

func main() {
	n := 0
	fmt.Scanf("%d", &n)
	test2 := make([]int, n)

	for ; n > 0; n-- {
		tmp := 0
		fmt.Scanf("%d", &tmp)
		test2[n-1] = tmp
	}

	// test1 := []int{3, 1, 2, 5, 10, 8, 12}
	res := permute(test2)
	// var buff string
	// for _, Res := range res {
	// 	// fmt.Sprintf(strings.Join(Res, " "))
	// 	// buff += fmt.Sprintf(strings.Join(Res, " "))

	// 	buff += strings.Replace(strings.Trim(fmt.Sprint(Res), "[]"), " ", " ", -1)
	// 	buff += "\n"
	// }
	fmt.Println(res)
	// fmt.Println(buff)
	// slice := strings.Split(ss, ",")
}

func permute(nums []int) [][]int {
	var (
		result  [][]int
		res     = make([]int, len(nums))
		visited = make([]bool, len(nums))
		dfs     func(int)
	)
	dfs = func(depth int) {
		if depth == len(nums) {
			tmp := make([]int, len(res))
			copy(tmp, res)
			result = append(result, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if !visited[i] {
				visited[i] = true
				res[depth] = nums[i]
				dfs(depth + 1)
				visited[i] = false
			}
		}
	}
	dfs(0)
	return result
}
func CommandBash() {
	cmd := exec.Command("/bin/bash", "python3 ../robclass/robclass.py")
	fmt.Printf("in all")
	cmd.Stdin = strings.NewReader("some input")
	cmd.Run()
	// fmt.Println(cmd.Dir)
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("in all caps: %q\n", out.String())
}
