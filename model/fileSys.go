// Package model 数据模型
package model

import (
	"os"
	"strings"
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
}

func InitRoot(root Path) *File {
	f := &File{}
	f.absolute = root
	paths := strings.Split(root, "\\")
	f.relative = paths[len(paths)-1]
	f.space = 1
	f.isDir = true
	f.InitFiles()
	return f
}
func (f *File) InitFiles() {
	files, err := os.ReadDir(f.absolute)
	if err != nil {
		return
	}
	f.open = true
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
func (f *File) Open() {
	f.open = true
	if f.isDir {
		f.InitFiles()
	}
}
func (f *File) Close() {
	f.open = false
	// if f.isDir {
	// 	f.files = nil
	// }
}

func (f *File) View(index int) string {
	var str = strings.Builder{}
	stack := make([]*File, 0)
	files := make([]*File, 0)
	stack = append(stack, f)
	for len(stack) != 0 {
		head := stack[0]
		stack = stack[1:]
		files = append(files, head)
		stack = append(stack, head.files...)
	}
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
	return str.String()
}
