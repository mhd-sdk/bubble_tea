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

Composants prêts à l'emploi -> [`github.com/charmbracelet/bubbles`](https://github.com/charmbracelet/bubbles)

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
layout: center
class: text-center
---

# Partie 1 : l'API

---

# `tea.Model`

L'interface que toute application Bubble Tea doit implémenter.

```go
type Model interface {
    Init()             tea.Cmd
    Update(tea.Msg)   (tea.Model, tea.Cmd)
    View()             string
}
```

<div class="grid grid-cols-3 gap-4 mt-6 text-sm">
<div class="bg-blue-500 bg-opacity-20 p-3">

**Init**
Appelé une fois au démarrage. Retourne un `Cmd` optionnel pour lancer les effets initiaux.

</div>
<div class="bg-red-500 bg-opacity-20 p-3">

**Update**
Reçoit un `Msg`, retourne le nouvel état et un `Cmd` optionnel.

</div>
<div class="bg-green-500 bg-opacity-20 p-3">

**View**
Retourne une `string` à afficher. Fonction pure, aucun effet de bord.

</div>
</div>

---

# `tea.Msg`

`Msg` est une **interface vide**. N'importe quel type Go peut devenir un message.

```go
type Msg interface{}
```

<div class="grid grid-cols-2 gap-6 mt-4 text-sm">
<div>

**Messages built-in**
```go
tea.KeyMsg        // touche clavier
tea.WindowSizeMsg // redimensionnement
tea.MouseMsg      // événement souris
```

</div>
<div>

**Messages custom**
```go
type choicesLoadedMsg struct {
    choices []string
    err     error
}
```

</div>
</div>

---

# `tea.Cmd`

`Cmd` est une **fonction** que le runtime exécute en goroutine et dont le résultat revient dans Update.

```go
type Cmd func() Msg
```

<div class="grid grid-cols-2 gap-6 mt-6 text-sm">
<div>

**Pourquoi pas du synchrone ?**

`Update` tourne sur la boucle principale. Un appel bloquant gèle toute l'UI, plus aucune touche ne répond, le terminal ne se redessine plus.

Un `Cmd` délègue le travail à une goroutine, **la boucle reste fluide** pendant l'I/O.

</div>
<div>

**Définir et lancer une commande**
```go
func loadChoices() tea.Cmd {
    return func() tea.Msg {
        choices, err := fetchFromAPI()
        return choicesLoadedMsg{choices, err}
    }
}

// dans Update :
case "r":
    m.loading = true          // appliqué immédiatement
    return m, loadChoices()   // lancé en goroutine
```

</div>
</div>


---
layout: center
class: text-center
---

# Partie 2 : mini démo

---

# Le Modèle

```go
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
```

<div class="mt-4 text-sm opacity-70">

> `choices` est vide au démarrage il sera rempli par un `Cmd` dans `Init`

</div>

---

# Init

`Init` lance le chargement des données au démarrage via un `Cmd`.

```go
type choicesLoadedMsg struct {
    choices []string
}

func (m model) Init() tea.Cmd {
    return loadChoices()
}

func loadChoices() tea.Cmd {
    return func() tea.Msg {
        // simulation d'un appel externe (base de données, API…)
        time.Sleep(2 * time.Second)
        languages := []string{"Go", "Rust", "TypeScript", "Python", "C"}
        return choicesLoadedMsg{languages}
    }
}
```

---

# Update

```go {*}{maxHeight:'320px'}
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
```

<style>
.slidev-code, .slidev-code .shiki, pre { font-size: 0.52rem; }
</style>

---

# View

```go {*}{maxHeight:'320px'}
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
```

<style>
.slidev-code, .slidev-code .shiki, pre { font-size: 0.52rem; }
</style>

---

# Assembler le tout

```go
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
Quel est ton langage de programmation favori ?

> [x] Go
  [ ] Rust
  [ ] TypeScript
  [ ] Python
  [ ] C

q pour quitter
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

# 🧋 Merci !

<div class="text-2xl mt-4 opacity-80">

`github.com/charmbracelet/bubbletea`

</div>

