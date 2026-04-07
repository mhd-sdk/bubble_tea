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

# Qui suis-je ?

<div class="grid grid-cols-2 gap-8 mt-2">

<div>

### Mehdi Seddik
<div class="mt-[-0.5rem]">

Dev Backend chez **Batch** depuis septembre 2025

</div>

### Formation
<div class="mt-[-0.5rem]">

- Mastère expert développement logiciel Ynov Lyon
- DUT informatique — IUT Lyon 1

</div>

### Contact
<div class="mt-[-0.5rem] text-sm opacity-70">

🔗 linkedin.com/in/mehdi-seddik

</div>

</div>

<div class="flex items-center justify-center">
  <img src="/photo.jpg" class="rounded-xl h-64 object-cover" />
</div>

</div>

---

# Qu'est-ce que Bubble Tea ?

Une lib Go pour construire des **interfaces en ligne de commande interactives** (TUI)

<div class="grid grid-cols-2 gap-8 mt-2">

<div>

### Caractéristiques
<div class="mt-[-0.5rem]">

- Basé sur l'**architecture Elm**
- Gestion clavier et souris
- Composable et modulaire

</div>

### Créé par [Charm](https://charm.sh)
<div class="mt-[-0.5rem]">

Charm est une société open source qui développe des outils pour améliorer l'expérience terminal. 

</div>

</div>

<div>
<img src="https://camo.githubusercontent.com/2978ecb9e7d89397213797fb538e92d8083b94eac6b516d2fb03189e9db06999/68747470733a2f2f73747566662e636861726d2e73682f627562626c657465612f627562626c657465612d6578616d706c652e676966" class="w-full h-full object-contain" />
</div>

</div>

<div class="mt-4 text-center">

```bash
go get github.com/charmbracelet/bubbletea
```

</div>

---

# L'écosystème Charm

<div class="flex flex-col gap-2 mt-4 text-sm">

<div class="border border-purple-500 border-opacity-50 px-4 py-2 flex items-center gap-4">
<span class="font-bold whitespace-nowrap">🫧 Bubbles</span>
<span class="opacity-60">Composants prêts à l'emploi : inputs, viewports, spinners…</span>
<a href="https://github.com/charmbracelet/bubbles" class="ml-auto font-mono opacity-50 text-xs">github.com/charmbracelet/bubbles</a>
</div>

<div class="border border-pink-500 border-opacity-50 px-4 py-2 flex items-center gap-4">
<span class="font-bold whitespace-nowrap">💄 Lip Gloss</span>
<span class="opacity-60">Style, mise en page et formatage pour le terminal</span>
<a href="https://github.com/charmbracelet/lipgloss" class="ml-auto font-mono opacity-50 text-xs">github.com/charmbracelet/lipgloss</a>
</div>

<div class="border border-blue-500 border-opacity-50 px-4 py-2 flex items-center gap-4">
<span class="font-bold whitespace-nowrap">🎯 Harmonica</span>
<span class="opacity-60">Animations spring pour des mouvements fluides et naturels</span>
<a href="https://github.com/charmbracelet/harmonica" class="ml-auto font-mono opacity-50 text-xs">github.com/charmbracelet/harmonica</a>
</div>

<div class="border border-yellow-500 border-opacity-50 px-4 py-2 flex items-center gap-4">
<span class="font-bold whitespace-nowrap">🖱️ BubbleZone</span>
<span class="opacity-60">Tracking des événements souris par zone</span>
<a href="https://github.com/lrstanley/bubblezone" class="ml-auto font-mono opacity-50 text-xs">github.com/lrstanley/bubblezone</a>
</div>

<div class="border border-green-500 border-opacity-50 px-4 py-2 flex items-center gap-4">
<span class="font-bold whitespace-nowrap">📊 ntcharts</span>
<span class="opacity-60">Graphiques dans le terminal pour Bubble Tea et Lip Gloss</span>
<a href="https://github.com/NimbleMarkets/ntcharts" class="ml-auto font-mono opacity-50 text-xs">github.com/NimbleMarkets/ntcharts</a>
</div>

<div class="border border-green-500 border-opacity-50 px-4 py-2 flex items-center gap-4">
<span class="font-bold whitespace-nowrap">✨ Glamour</span>
<span class="opacity-60">Rendu Markdown dans le terminal</span>
<a href="https://github.com/charmbracelet/glamour" class="ml-auto font-mono opacity-50 text-xs">github.com/charmbracelet/glamour</a>
</div>

</div>

---

# Ils utilisent Bubble Tea

<div class="flex flex-col gap-2 mt-6 text-sm">

<div class="border border-white border-opacity-10 px-4 py-2 flex items-center gap-4">
<span>🏠</span>
<a href="https://github.com/twpayne/chezmoi" class="font-bold">chezmoi</a>
<span class="opacity-60">gestionnaire de dotfiles</span>
<a href="https://github.com/twpayne/chezmoi" class="ml-auto font-mono opacity-50 text-xs">github.com/twpayne/chezmoi</a>
</div>

<div class="border border-white border-opacity-10 px-4 py-2 flex items-center gap-4">
<span>🐙</span>
<a href="https://github.com/jesseduffield/lazygit" class="font-bold">lazygit</a>
<span class="opacity-60">TUI Git</span>
<a href="https://github.com/jesseduffield/lazygit" class="ml-auto font-mono opacity-50 text-xs">github.com/jesseduffield/lazygit</a>
</div>

<div class="border border-white border-opacity-10 px-4 py-2 flex items-center gap-4">
<span>✨</span>
<a href="https://github.com/charmbracelet/gum" class="font-bold">gum</a>
<span class="opacity-60">scripts shell interactifs</span>
<a href="https://github.com/charmbracelet/gum" class="ml-auto font-mono opacity-50 text-xs">github.com/charmbracelet/gum</a>
</div>

</div>

<div class="mt-4 text-xs opacity-50 text-right">
<a href="https://github.com/charmbracelet/bubbletea?tab=readme-ov-file#staff-favourites">et bien d'autres</a>
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

# L'Architecture Elm

> Un pattern fonctionnel et réactif pour gérer l'état ⚡

---

# L'Architecture Elm : Vue d'ensemble

<div class="flex gap-4" style="height: calc(100% - 4rem);">

  <div class="flex items-start" style="width: fit-content;">
    <img src="https://mentalbreaktown.blog/wp-content/uploads/2020/11/elm.png" style="max-height: 100%; width: auto;" />
  </div>

  <div class="flex flex-col gap-4 text-sm self-start">
    <div class="bg-blue-500 bg-opacity-20 p-3"><b>Model</b> : structure de données qui représente tout l'état de l'app</div>
    <div class="bg-green-500 bg-opacity-20 p-3"><b>View</b> : rendu pur du modèle en string</div>
    <div class="bg-red-500 bg-opacity-20 p-3"><b>Update</b> : réagit aux événements, retourne un nouveau modèle</div>
    <div class="text-xs opacity-50 border-l-2 border-white border-opacity-20 pl-3">
      <b>Elm</b> est le langage qui a inventé ce pattern.<br/>
      <b>MVU</b> (Model–View–Update) en est la généralisation.
    </div>
    <p class="text-xs opacity-50">Même système que useReducer ou Redux avec React</p>
  </div>

</div>

---
layout: center
class: text-center
---

# L'API Bubble Tea

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

# Démo

<div class="text-left inline-block">

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
        case "r":
            m.loading = true
            m.choices = nil
            m.selected = make(map[int]struct{})
            m.cursor = 0
            return m, loadChoices()
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

<div class="mt-2 text-sm">

**3 méthodes suffisent** pour une TUI complète :

| Méthode | Signature |
|---------|-----------|
| `Init` | `() → Cmd` |
| `Update` | `Msg → (Model, Cmd)` |
| `View` | `() → string` |

</div>


---
layout: center
class: text-center
---

# Démo un peu plus avancée 🚀

---
layout: center
class: text-center
---

# 🧋 Merci !

<div class="text-2xl mt-4 opacity-80">

`github.com/charmbracelet/bubbletea`

</div>

<div class="text-xl mt-4 opacity-60">

Slides : `github.com/mhd-sdk/bubble_tea`

</div>

