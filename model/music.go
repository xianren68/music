package model

import (
	"os"
	"strings"
	"time"

	"github.com/faiface/beep/mp3"
	"github.com/faiface/beep/speaker"
)

type Music struct {
	name        string
	duration    int
	currenttime int
}

// 关闭计时器
var ch chan struct{} = make(chan struct{}, 1)
var DefaultMusic *Music = &Music{
	name: "没有音乐播放",
}

func Play(path string) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0666)
	splits := strings.Split(path, "\\")
	DefaultMusic.name = "正在播放：" + splits[len(splits)-1]
	if err != nil {
		DefaultMusic.name = err.Error()
	}
	defer file.Close()
	// 创建解码器
	streamer, format, err := mp3.Decode(file)
	if err != nil {
		DefaultMusic.name = err.Error()
	}
	defer streamer.Close()

	speaker.Init(format.SampleRate, format.SampleRate.N(time.Second/10))
	speaker.Play(streamer)

	select {}
}
func CLose() {
	DefaultMusic.name = "没有音乐播放"
	ch <- struct{}{}
	speaker.Close()
}

func (m *Music) View() string {
	return "    " + m.name + "\n"
}
