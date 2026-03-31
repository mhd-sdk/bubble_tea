---
theme: default
colorSchema: dark
class: text-center
highlighter: shiki
lineNumbers: false
info: |
  ## Bubble Tea
  Framework TUI en Go basé sur l'architecture Elm
drawings:
  persist: false
transition: slide-left
title: Bubble Tea, TUI en Go
mdc: true
---

# 🧋 Bubble Tea

<div>
  <span>
        Rendez vos CLI interactives !
  </span>
</div>

<div class="abs-bl m-6 flex flex-col items-start gap-1">
  <span class="text-xs opacity-50">Mehdi Seddik — 2026 Meetup Go Lyon #32</span>

</div>

---

# Qu'est-ce que Bubble Tea ?

Une lib Go pour construire des **interfaces en ligne de commande interactives** (TUI)

<div class="grid grid-cols-2 gap-8 mt-8">

<div>

### Caractéristiques
- Basé sur l'**architecture Elm**
- Gestion clavier et souris
- Composable et modulaire
- Écosystème Charm (Lip Gloss, Bubbles…)

</div>

<div>
<img src="https://camo.githubusercontent.com/2978ecb9e7d89397213797fb538e92d8083b94eac6b516d2fb03189e9db06999/68747470733a2f2f73747566662e636861726d2e73682f627562626c657465612f627562626c657465612d6578616d706c652e676966" class="w-full h-full object-contain" />
</div>

</div>

<div class="mt-8 text-center">

```bash
go get github.com/charmbracelet/bubbletea
```

</div>

---

# Les Bubbles disponibles

Composants prêts à l'emploi — [`github.com/charmbracelet/bubbles`](https://github.com/charmbracelet/bubbles)

<div class="flex justify-center mt-6">
<img src="https://camo.githubusercontent.com/3a59a4885c8f93fbc9b19b1d2437b46a076002ff82fbfe2125ad98b5736bf5b7/68747470733a2f2f73747566662e636861726d2e73682f627562626c65732d6578616d706c65732f6c6973742e676966" class="h-80" />
</div>

---
layout: center
class: text-center
---

# L'Architecture Elm


> Un pattern fonctionnel et réactif pour gérer l'état ⚡

---

# L'Architecture Elm : Vue d'ensemble

<div class="flex gap-4" style="height: calc(100% - 4rem);">

  <div class="flex items-start" style="width: fit-content;">
    <img src="https://mentalbreaktown.blog/wp-content/uploads/2020/11/elm.png" style="max-height: 100%; width: auto;" />
  </div>

  <div class="flex flex-col gap-4 text-sm  self-start">
    <div class="bg-blue-500 bg-opacity-20 p-3"><b>Model</b> : structure de données qui représente tout l'état de l'app</div>
    <div class="bg-green-500 bg-opacity-20 p-3"><b>View</b> : rendu pur du modèle en string</div>
    <div class="bg-red-500 bg-opacity-20 p-3"><b>Update</b> : réagit aux événements, retourne un nouveau modèle</div>
    <p class="text-xs opacity-50 ">Même système que useReducer ou Redux pour avec React</p>
  </div>

</div>

---
layout: image
image: /coding-time.gif
---

# Coding Time

<style>
h1 {
  color: white !important;
  @apply !text-shadow-lg;
  @apply !text-center;
  @apply !text-8xl
}
</style>

---

# Le Modèle

**Le modèle** est une `struct` Go qui contient **tout l'état** de l'application

```go
package main

// La structure du modèle qui contiendra tout le "state"
type model struct {
    choices  []string
    cursor   int
    selected map[int]struct{}
}

// L'état initial
func initialModel() model {
    return model{
        choices:  []string{"Bubble Tea", "Lipgloss", "Glamour"},
        selected: make(map[int]struct{}),
    }
}
```

<div class="mt-4 text-sm opacity-70">

> Pas de mutation globale, l'état est **immuable** et remplacé à chaque update

</div>

---

# Init

**Init** est appelé une seule fois au démarrage. Retourne un `Cmd` optionnel.

```go 
// tea.Cmd représente une commande asynchrone (I/O, timer…)
// tea.Msg est le résultat qui reviendra dans Update

func (m model) Init() tea.Cmd {
    // Pas d'I/O au démarrage ici → nil
    return nil
}
```

<div class="mt-6">

Exemples de `Cmd` utiles :

```go
tea.Tick(time.Second, func(t time.Time) tea.Msg { … }) // timer
tea.ExecProcess(cmd, func(err error) tea.Msg { … })    // process
```

</div>

---

# Update

**Update** reçoit un `Msg`, met à jour le modèle et retourne un nouveau modèle + une commande optionnelle


```go {*}{maxHeight:'350px'}
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
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
```

<style>
.slidev-code, .slidev-code .shiki, pre { font-size: 0.52rem; }
</style>

---

# View

**View** transforme le modèle en `string` : fonction pure, aucun effet de bord

```go {*}{maxHeight:'350px'}
func (m model) View() string {
    s := "Quelle lib Charm utilises-tu ?\n\n"

    for i, choice := range m.choices {

        // Curseur
        cursor := " "
        if m.cursor == i {
            cursor = ">"
        }

        // Item sélectionné
        checked := " "
        if _, ok := m.selected[i]; ok {
            checked = "x"
        }

        s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
    }

    s += "\nAppuie sur q pour quitter.\n"
    return s
}
```

<style>
.slidev-code, .slidev-code .shiki, pre { font-size: 0.52rem; }
</style>

---

# Assembler le tout

```go {all|3-7}
func main() {
    p := tea.NewProgram(initialModel())

    if _, err := p.Run(); err != nil {
        fmt.Printf("Erreur: %v", err)
        os.Exit(1)
    }
}
```

<div class="mt-6 grid grid-cols-2 gap-6">

<div>

**Le résultat :**
```
Quelle lib Charm utilises-tu ?

> [x] Bubble Tea
  [ ] Lipgloss
  [ ] Glamour

Appuie sur q pour quitter.
```

</div>

<div class="text-sm mt-2">

**3 méthodes suffisent** pour une TUI complète :

| Méthode | Signature |
|---------|-----------|
| `Init` | `() → Cmd` |
| `Update` | `Msg → (Model, Cmd)` |
| `View` | `() → string` |

</div>

</div>

---

# Les Commandes (Cmd)

Les `Cmd` permettent de gérer les **effets de bord** de façon contrôlée

```go
// Un Cmd est juste une fonction qui retourne un Msg
type Cmd func() Msg

// Exemple : appel HTTP asynchrone
func fetchData() tea.Cmd {
    return func() tea.Msg {
        resp, err := http.Get("https://api.example.com/data")
        if err != nil {
            return errMsg{err}
        }
        return dataMsg{resp}
    }
}

// Dans Update, on lance la commande
case tea.KeyMsg:
    if msg.String() == "f" {
        return m, fetchData()
    }
```

<div class="mt-4 text-sm opacity-70">
→ L'I/O est isolé hors du cycle Model/View, la logique reste testable
</div>

---

# L'écosystème Charm

<div class="grid grid-cols-3 gap-4 mt-6">

<div class="border border-pink-500 border-opacity-50 p-4">

### 💄 Lip Gloss
Styling CSS-like pour le terminal

```go
style := lipgloss.NewStyle().
    Bold(true).
    Foreground(lipgloss.Color("#FF06B7")).
    PaddingLeft(4)

fmt.Println(style.Render("Hello!"))
```

</div>

<div class="border border-blue-500 border-opacity-50 p-4">

### 🫧 Bubbles
Composants prêts à l'emploi

- `textinput`
- `textarea`
- `list`
- `table`
- `progress`
- `spinner`
- `viewport`

</div>

<div class="border border-green-500 border-opacity-50 p-4">

### ✨ Glamour
Rendu Markdown dans le terminal

```go
out, _ := glamour.Render(
    "# Hello **world**",
    "dark",
)
fmt.Print(out)
```

</div>

</div>

---
layout: center
class: text-center
---

# Pourquoi Bubble Tea ?

<div class="grid grid-cols-2 gap-12 mt-8 text-left">

<div>

### ✅ Avantages
- Architecture **prévisible** et testable
- Séparation claire état / rendu / effets
- Composants **réutilisables** (Bubbles)
- Style puissant avec Lip Gloss
- Communauté active (Charm)

</div>

<div>

### 🎯 Idéal pour
- CLIs interactifs
- Dashboards terminal
- Outils de dev (git, k8s, docker…)
- Jeux en ASCII
- Apps SSH avec Wish

</div>

</div>

---
layout: center
class: text-center
---

# 🧋 Merci !

<div class="text-2xl mt-4 opacity-80">

`github.com/charmbracelet/bubbletea`

</div>

<div class="mt-8 grid grid-cols-3 gap-4 text-sm opacity-60">
  <div>github.com/charmbracelet/lipgloss</div>
  <div>github.com/charmbracelet/bubbles</div>
  <div>github.com/charmbracelet/glamour</div>
</div>
