# project-jump

I was tired of typing:

```bash
cd ~/dev/some-project
```

and then realizing the project was actually in:

```bash
~/Desktop/random-shit/final-final-actually-final/some-project
```

and then spending 30 seconds pretending I remembered where I put it.

So I built this.

`project-jump` is a tiny Go utility that scans my projects, throws them into fzf, and lets me jump/open them without playing hide and seek with my filesystem.

## What it does

```bash
pjump
```

Find project.

Hit enter.

Choose:

```text
nvim
code
finder
github
```

Or skip the menu entirely:

```text
Ctrl-N -> Neovim
Ctrl-V -> VS Code
Ctrl-F -> Finder
Ctrl-G -> GitHub
```

Because life is short and my attention span is shorter.

## Features

* Fuzzy project search
* Frecency-based ranking
* Git status preview
* Project tree preview
* GitHub shortcuts
* Mildly weaponized fzf

## Why not use...

### cd

Because it's 2026.

### find

Because I enjoy happiness.

### Raycast

Because I wanted something in my terminal.

### zoxide

Actually zoxide is cool.

This is just my problem now.

## Philosophy

Every feature must answer one question:

> Does this get me into a project faster?

If the answer is no, it probably doesn't belong here.

That's why this project currently contains:

```text
Go
fzf
eza
git
```

and not:

```text
microservices
kubernetes
blockchain
AI
```

## Installation

If you're reading this, you probably know how to build a Go binary.

If not:

```bash
go build -o ~/.local/bin/pj
```

Then wire `pj` into your shell however you like.

## Future

Whatever annoys me next.
