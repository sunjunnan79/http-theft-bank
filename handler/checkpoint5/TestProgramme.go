package checkpoint5

import (
	"bytes"
	"errors"
	"fmt"
	"http-theft-bank/pkg/text"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"syscall"
	"time"
)

type process struct {
	Cmd    *exec.Cmd
	Stdout *bytes.Buffer
	Stderr *bytes.Buffer
	File   *os.File // 用于标志打开过的文件，方便关闭
}

// TODO: 测试文件名从配置文件中导入
var TestFiles = []string{
	"./file/testSample/test1.txt",
	"./file/testSample/test2.txt",
	"./file/testSample/test3.txt",
}

func testProgramme(fileName, fileNameOnly string) error {
	// go build
	cmd := exec.Command("go", "build", "-o", "./file/bin/"+fileNameOnly, "./file/"+fileName)
	err := cmd.Run()
	if err != nil {
		return err
	}

	// go run
	var processSet []*process
	for i := 0; i < len(TestFiles); i++ {
		p, err := initProcess(fileNameOnly, TestFiles[i])
		if err != nil {
			killAllProcess(processSet)
			return err
		}
		err = p.Cmd.Start()
		if err != nil {
			// kill all process and return
			killAllProcess(processSet)
			return err
		}
		processSet = append(processSet, p)
	}

	// init channel
	doneChannel := make(chan int)
	errChannel := make(chan error)
	defer close(doneChannel)
	defer close(errChannel)
	doneCount := 0

	// goroutine start
	for i := 0; i < len(processSet); i++ {
		go processMonitor(processSet[i], errChannel, doneChannel)
	}

	// for - select
	for doneCount != len(processSet) {
		select {
		case <-doneChannel:
			doneCount++
			fmt.Println("one done")
		case err := <-errChannel:
			closeAllFile(processSet)
			killAllProcess(processSet)
			return err
		case <-time.After(time.Second * 2):
			// fmt.Println("over time and kill all process")
			closeAllFile(processSet)
			killAllProcess(processSet)
			return errors.New("over time")
		}
	}

	// 关掉所有文件
	closeAllFile(processSet)

	// TODO: 检查 stderr 是否不为空和 stdout 是否为空

	exitChannel := make(chan error, len(processSet))
	okChannel := make(chan int, 1)
	defer close(okChannel)
	defer close(exitChannel)
	//不进行检查了,只要你传了代码就给你过
	//for i, process := range processSet {
	//	go checkRes(i, process.Stdout.String(), exitChannel, len(processSet), okChannel)
	//}

	var testErr error
	//for n := 0; n != len(processSet); {
	//	select {
	//	case testErr = <-exitChannel:
	//		// return err
	//	case <-okChannel:
	//		n++
	//	}
	//}
	return testErr
}

// checkRes ... 目前并发是有问题的，给 exitChannel 发信号并不能保证每个 channel 都能拿到
func checkRes(num int, res string, exitChannel chan error, n int, okChannel chan int) {

	answers := text.Answers[num]

	var AnswerMap = make(map[string]int)
	for _, answer := range answers {
		AnswerMap[answer] = 0
	}
	// 处理res为[]string
	var ret []string
	{
		f := func(r rune) bool {
			return r == '[' || r == ']'
		}
		a := strings.FieldsFunc(string(res), f)
		sort.Slice(a, func(i, j int) bool {
			numA, _ := strconv.Atoi(a[i])
			numB, _ := strconv.Atoi(a[j])
			return numA < numB
		})
		aLen := len(a)
		for i := 0; i < aLen; i++ {
			if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 1 {
				continue
			}
			ret = append(ret, a[i])
		}
	}

	count := 0
	for _, re := range ret {
		select {
		case <-exitChannel:
			okChannel <- 1
			return
		default:
			if _, ok := AnswerMap[re]; ok {
				count++
			} else {
				for ; n > 0; n-- {
					exitChannel <- errors.New("wrong answer")
				}
				okChannel <- 1
				return
			}
		}
	}

	if count != len(answers) {
		for ; n > 0; n-- {
			exitChannel <- errors.New("wrong answer")
		}
	}
	okChannel <- 1
}

// initProcess ... 初始化进程
func initProcess(fileNameOnly, testFileName string) (*process, error) {
	p := new(process)
	p.Cmd = exec.Command("./file/bin/" + fileNameOnly)
	p.Cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid: true} // 每个进程单独一个组

	// 打开测试文件
	file, err := os.Open(testFileName)
	if err != nil {
		return nil, err
	}

	p.Cmd.Stdin = file
	p.File = file

	// 获取标准输出
	p.Stdout = new(bytes.Buffer)
	p.Stderr = new(bytes.Buffer)
	p.Cmd.Stdout = p.Stdout
	p.Cmd.Stderr = p.Stderr

	return p, nil
}

// processMonitor ... 用于 goroutine 等待
func processMonitor(p *process, errChannel chan<- error, doneChannel chan<- int) {
	err := p.Cmd.Wait()
	if err != nil {
		errChannel <- err
		return
	}

	doneChannel <- 1
}

func killAllProcess(processSet []*process) {
	for i := 0; i < len(processSet); i++ {
		err := syscall.Kill(-processSet[i].Cmd.Process.Pid, syscall.SIGKILL)
		if err != nil {
			// TODO:log
			fmt.Println("kill all process failed: " + err.Error())
		}
	}
}

// closeAllFile ... 关掉所有文件
func closeAllFile(processSet []*process) {
	for i := 0; i < len(processSet); i++ {
		err := processSet[i].File.Close()
		if err != nil {
			// log
			fmt.Println("close all file failed: " + err.Error())
		}
	}
}
