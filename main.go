package main

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {
	m := initialModel()
	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Printf("发生错误退出: %v", err)
		os.Exit(1)
	}

	if m.quit {
		return
	}

	LoadCPU(m.selected)
}
