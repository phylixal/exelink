/* Public domain.
 * inspired by callexec.c Originally written by Akira Kakuto  <kakuto@fuk.kindai.ac.jp>
 *
 * WIN32 wrapper program replacing Unix symlinks such as,
 * e.g., vi -> vim.
 *
 */
//缩小编译后体积
//go build -ldflags "-s -w"

package main

import (
	"os"
	"os/exec"
)

func main() {
	exeProg := "D:\\Green\\Arms\\MSYS32\\usr\\bin\\git.exe"
	os.Setenv("PATH", os.Getenv("PATH") + ";D:\\Green\\Arms\\MSYS32\\usr\\bin")
	cmd := exec.Command(exeProg, os.Args[1:]...)
	//传入数组，后接三个点使得此函数可接收非确定个数的参数，意义同直接传入多个字符串
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Start()
	cmd.Run()
	cmd.Wait()
}
