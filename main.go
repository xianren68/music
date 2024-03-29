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
	// 是否点击enter
	isEnter bool
	music   *model.Music
}

func (m *Model) View() string {
	f := m.file.View(&m.isEnter)
	return m.music.View() + f
}
func initModel() *Model {
	m := &Model{}
	m.file = model.InitRoot(root)
	m.music = model.DefaultMusic
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
		case "w":
			m.file.Index--
			if m.file.Index < 0 {
				m.file.Index = 0
			}
		case "s":
			m.file.Index++
			if m.file.Index >= len(m.file.AllFiles) {
				m.file.Index = len(m.file.AllFiles) - 1
			}
		case "ctrl+s":
			model.CLose()
		case "enter":
			m.isEnter = true
		}
	}
	return m, nil
}
