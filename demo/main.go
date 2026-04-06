package main

import (
	"fmt"
	"os"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

// --- Model ---

type model struct {
	choices  []string
	cursor   int
	selected map[int]struct{}
	loading  bool
}

func initialModel() model {
	return model{
		loading:  true,
		selected: make(map[int]struct{}),
	}
}

// --- Messages ---

type choicesLoadedMsg struct {
	choices []string
}

// --- Commands ---

func loadChoices() tea.Cmd {
	return func() tea.Msg {
		// simulation d'un appel externe (base de données, API…)
		time.Sleep(2 * time.Second)
		languages := []string{"Go", "Rust", "TypeScript", "Python", "C"}
		return choicesLoadedMsg{languages}
	}
}

// --- Init ---

func (m model) Init() tea.Cmd {
	return loadChoices()
}

// --- Update ---

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case choicesLoadedMsg:
		m.loading = false
		m.choices = msg.choices

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		case "down", "j":
			if m.cursor < len(m.choices)-1 {
				m.cursor++
			}
		case "enter", " ":
			if _, ok := m.selected[m.cursor]; ok {
				delete(m.selected, m.cursor)
			} else {
				m.selected[m.cursor] = struct{}{}
			}
		}
	}
	return m, nil
}

// --- View ---

func (m model) View() string {
	if m.loading {
		return "Chargement…\n"
	}

	s := "Quel est ton langage de programmation favori ?\n\n"

	for i, choice := range m.choices {
		cursor := " "
		if m.cursor == i {
			cursor = ">"
		}

		checked := " "
		if _, ok := m.selected[i]; ok {
			checked = "x"
		}

		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	s += "\nq pour quitter\n"
	return s
}

// --- Main ---

func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Erreur: %v", err)
		os.Exit(1)
	}
}
