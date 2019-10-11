package main

import (
	"fmt"
	"os/exec"
	//"bytes"
	//"io"
	//"bufio"
	"bytes"
)

func main() {
	//cmd0 := exec.Command("echo", "-n", "My first command comes from golang.")
	//stdout0, err := cmd0.StdoutPipe()
	//if err != nil {
	//	fmt.Printf("Error: Couldn't obtain the stdout pipe for command no.0: %s\n", err)
	//	return
	//}
	//if err := cmd0.Start(); err != nil {
	//	fmt.Printf("Error: The command No.0 can not be startup: %s\n", err)
	//	return
	//}
	//buffer output
	//var outputBuf0 bytes.Buffer
	//for{
	//	output0 := make([]byte, 5)
	//	n, err := stdout0.Read(output0)
	//	if err != nil {
	//		if err == io.EOF{
	//			break
	//		}else{
	//			fmt.Printf("Error: Couldn't read data from the pipe: %s\n", err)
	//			return
	//		}
	//	}
	//	if n>0{
	//		outputBuf0.Write(output0)
	//	}
	//}
	//fmt.Printf("%s\n", outputBuf0.String())
	//outputBuf0 := bufio.NewReader(stdout0)
	//output0, _, err := outputBuf0.ReadLine()
	//if err != nil {
	//	fmt.Printf("Error: Couldn`t read data from the pipe: %s\n", err)
	//	return
	//}
	//fmt.Printf("%s\n", string(output0))
	cmd1 := exec.Command("ps", "aux")
	cmd2 := exec.Command("grep", "apipe")
	var outputBuf1 bytes.Buffer
	cmd1.Stdout = &outputBuf1
	if err := cmd1.Start(); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	if err := cmd1.Wait(); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	cmd2.Stdin = &outputBuf1
	var outputBuf2 bytes.Buffer
	cmd2.Stdout = &outputBuf2
	if err := cmd2.Start(); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	if err := cmd2.Wait(); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return
	}
	fmt.Printf("%s\n", outputBuf2.Bytes())
}
