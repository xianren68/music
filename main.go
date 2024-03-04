package main

import (
	"flag"
	"music/model"

	tea "github.com/charmbracelet/bubbletea"
)

// 音乐文件夹根目录
var root string

func main() {
	// 命令行参数获取文件根目录
	flag.StringVar(&root, "root", ".", "音乐文件夹根目录")
	flag.Parse()
	p := tea.NewProgram(initModel())
	if _, err := p.Run(); err != nil {
		panic(err)
	}
}

type Model struct {
	file *model.File
}

func (m *Model) View() string {
	return m.file.View(0)
}
func initModel() *Model {
	m := &Model{}
	m.file = model.InitRoot(root)
	return m
}
func (m *Model) Init() tea.Cmd {
	// Just return `nil`, which means "no I/O right now, please."
	return nil
}
func (m *Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}
	return m, nil
}
