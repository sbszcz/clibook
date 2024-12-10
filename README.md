# Motivation

While preparing for the CKAD exam recently, I often felt the need to document and annotate the kubectl commands I used. My process involved copying the commands, pasting them into my note-taking tool (Obsidian), and adding comments. Over time, this manual process became cumbersome, and I began wondering if there might be an easier way to integrate this into my command-line-focused workflow.

This led me to explore zsh widgets more deeply, and I developed a newfound appreciation for the Z shell. However, I still needed a small backend tool to simplify the management of my notes. Being a fan of Go, I decided to implement the tool in Go.

I finally came up with this solution which works pretty well for me.

# Prerequisites

- [Z shell (zsh)](https://www.zsh.org) for a convenient shell experience
- [fzf](https://github.com/junegunn/fzf) as fuzzy finder for searching inside recorded cli notes
- Ideally vim or neovim as editor

# Installation

1. Install `clibook` binary

```bash
go install github.com/sbszcz/clibook@latest
```

2. Source zsh widgets

```bash
curl https://raw.githubusercontent.com/sbszcz/clibook/refs/heads/main/clibook-zsh-widgets.zsh -o clibook.zsh && source clibook.zsh
```

# Workflow

My current workflow with `clibook` looks like this.

1. Issuing a command!
1. Pressing `ctrl + e` opens neovim to take notes for the previously executed command.
1. Fuzzy searching through my _clibook_ by pressing `ctrl + t` and adding the selection to the command prompt for repeated execution.

## Example

![](clibook.gif)
