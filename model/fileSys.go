// Package model 数据模型
package model

import (
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	Reset   = "\033[0m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Purple  = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"
	RedWine = "\033[38;2;139;0;0m"
)

type Path = string

// Package File 文件夹.
type File struct {
	absolute Path
	relative Path
	files    []*File
	open     bool
	isDir    bool
	space    int
	FilesLen int
}

func InitRoot(root Path) *File {
	if len(os.Args) < 2 {
		fmt.Println("需要设置音乐文件夹,如下:\nmusic -root D:\\音频")
		os.Exit(0)
	}
	f := &File{}
	f.absolute = root
	paths := strings.Split(root, "\\")
	f.relative = paths[len(paths)-1]
	f.space = 1
	f.isDir = true
	f.FilesLen = len(f.files) - 1
	f.InitFiles()
	return f
}
func (f *File) InitFiles() {
	f.open = true
	// 前面初始化过，只修改状态即可
	if f.files != nil {
		return
	}
	files, err := os.ReadDir(f.absolute)
	if err != nil {
		return
	}
	for _, file := range files {
		if file.IsDir() {
			f.files = append(f.files, &File{
				absolute: f.absolute + "\\" + file.Name(),
				relative: file.Name(),
				open:     false,
				isDir:    true,
				space:    f.space + 2,
			})
		} else {
			f.files = append(f.files, &File{
				absolute: f.absolute + "\\" + file.Name(),
				relative: file.Name(),
				open:     false,
				isDir:    false,
				space:    f.space + 2,
			})
		}
	}
}
func (f *File) View(index int, enter *bool) string {
	var str = strings.Builder{}
	stack := make([]*File, 0)
	files := make([]*File, 0)
	stack = append(stack, f)
	var j = 0
	for len(stack) != 0 {
		head := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		files = append(files, head)
		if j == index && *enter {
			if head.isDir {
				head.open = !head.open
			} else {
				go func() {
					ch <- struct{}{}
					ch = make(chan struct{}, 1)
					Play(head.absolute)
					for {
						select {
						case <-ch:
							return
						default:
							// time.Sleep(time.Second)
						}
					}
				}()
				time.Sleep(time.Millisecond * 100)
			}
		}
		if head.open {
			head.InitFiles()
			stack = append(stack, head.files...)
		}
		j++
	}
	f.FilesLen = len(files) - 1
	for i, file := range files {
		//鼠标指针在哪一行
		cur := " "
		if i == index {
			cur = RedWine + "→"
		}
		// 空格
		space := strings.Repeat(" ", file.space)
		// 是否打开
		isOpen := Green + "> "
		if file.open {
			isOpen = Red + "- "
		}
		str.WriteString(cur)
		str.WriteString(space)
		if file.isDir {
			str.WriteString(isOpen)
		}
		str.WriteString(Cyan + file.relative)
		str.WriteByte('\n')
	}
	*enter = false
	return str.String()
}
