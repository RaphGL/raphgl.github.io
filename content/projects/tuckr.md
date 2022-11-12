---
title: "Tuckr"
description: "A dotfile manager heavily inspired by GNU Stow"
date: 2022-11-08T20:42:21-01:00
tags: [Rust, Automation, Dotfiles]
---

## [Find it on GitHub](https://github.com/RaphGL/Tuckr)

<!-- ABOUT THE PROJECT -->

Tuckr is a dotfile manager inspired by Stow and Git. Tuckr aims to make dotfile management less painful. It follows the same model as Stow, symlinking files onto $HOME. It works on all the major OSes (Windows, MacOS, Linux).  

Most dotfile managers out there rely on some sort of configuration file to be able manage your dotfiles, this project came about because I couldn't find any dotfile manager that was simple enough that you could just jump into it and start using it, with no need for reading lengthy documentation and dotfile manager specific things. 

A lot of people have been using Stow + Git to manage their dotfiles, while this approach is fine, Stow was not made for this, so it's not a perfect solution and it lacks features that are dotfile management specific, thus this project was born.

**What makes tuckr different?**

- No additional configuration required, everything that is needed comes setup by default
- You can manage your files from any directory
- Symlinks are tracked, the manager is smart enough to be able to manage them without conflicting with the rest of the symlinks in the system
- Hooks, write small scripts that will be run when you set up programs from your dotfiles
- Encrypted files for sensitive information  
  
  
Built With:
- [Rust](https://www.rust-lang.org/)
- [Clap](https://github.com/clap-rs/clap)

<!-- USAGE EXAMPLES -->

## Usage
```sh
$ tuckr add \* # adds all dotfiles to the system
$ tuckr add neovim zsh # adds the neovim and zsh dotfiles only
$ tuckr set \* # adds all the dotfiles and runs their hooks (scripts)
$ tuckr rm \* # removes all dotfiles from your system
```

```sh
Super powered GNU Stow replacement

Usage: tuckr <COMMAND>

Commands:
  set        Setup a program and run their hooks hooks
  add        Deploy dotfiles for PROGRAM
  rm         Remove configuration for a program on the system
  status     Check symlink status
  init       Initialize a dotfile folder
  from-stow  Converts a stow repo into a tuckr one
  help       Print this message or the help of the given subcommand(s)

Options:
  -h, --help     Print help information
  -V, --version  Print version information
```

### How it works

Tuckr works without having to use a configuration file by making a few minor choices for you. As long as you follow the file structure for tuckr repos it will do everything else for you automatically.

```sh
.
├── Configs # Dotfiles go here
├── Encrypts # Encrypted files go here
└── Hooks # Setup scripts go here
```

Your dotfiles should be one folder by program, the folder name will become how that program is named by tuckr.
```sh
.
├── Configs
│   ├── Program1
│   ├── Program2
├── Encrypts
└── Hooks
    ├── Program1
    └── Program2
```
As long as the names align between Configs, Hooks and Encrypts, they will work together.

### Using Hooks
Hooks are run before and after adding every program. Hooks that run before the program addition are prefixed with `pre`, scripts that run afterward are prefixed with `post`, as long as this is true you can name the file whatever you want.

```sh
Hooks
├── Program1
│   ├── post.sh
│   └── pre.sh
└── Program2
    ├── post.sh
    └── pre.sh
```
To run scripts for a program run `tuckr set <program_name>` or alternatively use a wildcard like so: `tuckr set \*` to run all hooks.

