---
title: Helix tutorial for Vim plebs
date: 2025-07-21
description: A walk through vim and helix differences and the various helix modes for people that are familiar with vim motions
---

(Neo)Vim has been one of the most popular editors for decades now and a lot of people have muscle memory for it.
So every time I suggest someone tries Helix, if they've used or use vim or vim mode on other editors, they almost alway complain about 
the keybindings being different being an impediment to them using the editor. So this is my attempt at trying to make the transition less painful
by drawing parallels between Helix and Vim bindings.  

## The low-hanging fruits
#### Basic Motions
The main difference between Helix and Vim is that Helix is selection based and reverses the motions. A lot of the actions are actually the same.
So for example `de` in Vim is `ed` in Helix, `ce` is `ec` and so on. The reason for this is that Helix selects text before doing anything with it.

You can think of the Helix binding as just Vim but entering visual mode before doing anything. So `de` would be `ved` and `ce` would be `vec`, the end result on
Vim would be the exact same as what would happen on Helix.

\

It's very common to move to the beginning or ending of the line or the file, this is slightly different in Helix.
Helix has a "goto mode" which is under the `g` key.

| Vim   | Helix   |
| :---: | :---:   |
| gg    | gg      |
| G     | ge      | 
| $     | gl      |
| 0     | gh      |

#### Doing selections
A cursor in Helix is essentially a one character selection, so if you attempt to use `dd` or `cc` or `yy` or any other double action to apply it to the entire line,
it just won't work. You have to do a selection before. Here's the most common selections:

- `x`: select the current line (if you keep pressing it it'll keep adding more lines to the selection)
- `%`: select the entire file
- `e`, `E`, `b`, `B`, `w`, `W`: essentially the same as in vim, select from the current cursor position until where they stop

## Helix modes and macros
#### Macros
Macros can be recorded by pressing `Q`, then after you've done whatever you wanted to record, you can press `Q` again to stop.
By default macros are recorded into the `@` register, you can change where your macros are recorded if you want to store and use multiple macros.
For example you could do `"mQ` to record a macro in a custom register `m`. `"<register>` is the syntax for changing registers.

You can also delete all macros by doing `:clear-register`.
Registers can also be used to save your selection (text) so if you use `d`, `y` and so on, they're also saved in registers.
\
Helix has a few special registers. If you want to know more about registers you can [check the documentation](https://docs.helix-editor.com/registers.html).

#### Space mode
Helix heavily relies on pickers to provide a lot of interactive features all these can be accessed by pressing space:

- `<space>b`: Buffer picker
- `<space>e`: File explorer
- `<space>s`: Symbol search
- `<space>/`: Global search
- `<space>d`: LSP diagnostics
- `<space>a`: Code actions

It has a few other features as well. You can check what if can do [here](https://docs.helix-editor.com/keymap.html?highlight=space%20mode#space-mode).

#### Window mode
Window mode allows you to manage splits in  Helix, you can enter window mode with `<space>w`.
Once in window mode, to create the splits you can use `s` for vertical splits and `v` for vertical splits.
Then you can use window mode to manage your splits. Such as moving between them (using hjkl in window mode), swapping them (using HJKL), 
\
You can change windows between vertical and horizontal using `t` (t stands for transpose), and you can quit the window with `q`.

#### Match mode
Match mode allows you to select text, in Vim you would usually use `gi` and `ga`, match mode is the Helix equivalent.

For example `giw` would be `miw` in helix (though for this case going back and forth or the other way around with the `w`, `b` and `e` keys is one less keystroke).
If you're inside of a `()` or a `{}` or some other pair and you have an LSP configured you can do `mim` or `mam` to select everything inside them.
If you don't have an LSP or you want to match something else you can just type them for example you could do `mi(` to match everything inside of a parenthesis.
\
You can also swap a `()` for a `{}` or vice verse with `mr({`. Or you could delete the surrounding parenthesis with `md(`

#### 

#### Command mode
Basic commands are mostly the same as in Vim. You can enter the command mode with `:`.

Common commands shared by Helix and Vim:

- `:q`: quit
- `:q!`: force quit
- `:w`: save
- `:w!`: force save
- `:x`: save and quit
- `:|`: pipe selection to command and substitute with command output 
- `:o`/`:open`: open file

Common commands that are different: 

- `:!`: required a space in helix as it's an alias to `:sh`
- `:lsp-restart`: (re)start LSP
- `:lsp-stop`: stop LSP
- `:log-open`: read logs
- `set-language`: equivalent to `:lang` in Vim
- `:format`: format text with language formatter

#### Multicursors
Helix supports multiple cursors. These can be created by using `<shift>C` the cursors can be disabled with `,`.

But the most powerful thing about it is being to spawn multiple cursors on certain locations based on their content.
To do this you can do a selection of the region where you want the cursors to spawn then you press `s` and type what you're looking for,
after you've matched the content you can press enter and continue editing.

To disable a selection you can use `;`.

<hr/>

If you want to learn more give the [official docs](https://docs.helix-editor.com/) a try.
You might also want to try the helix tutor by running `helix --tutor`.
