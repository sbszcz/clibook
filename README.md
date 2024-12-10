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
