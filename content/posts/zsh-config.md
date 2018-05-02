---
title: "ZSH Config"
date: 2018-05-02T13:49:22-07:00
tags: ["zsh", "dotfiles"]
draft: true
---

# Why Use a Config Framework?

Full disclosure: I am currently one of the main maintainers of prezto and the
sole maintainer of zsh-utils, so this article *will* have a bias. However, if
you see any incorrect points in here, please let me know and I'll happily
amend them.

If you've ever started zsh without a config file before, you were probably
greeted with something like the following:

```
Please pick one of the following options:

(1)  Configure settings for history, i.e. command lines remembered
     and saved by the shell.  (Recommended.)

(2)  Configure the new completion system.  (Recommended.)

(3)  Configure how keys behave when editing command lines.  (Recommended.)

(4)  Pick some of the more common shell options.  These are simple "on"
     or "off" switches controlling the shell's features.

(0)  Exit, creating a blank ~/.zshrc file.

(a)  Abort all settings and start from scratch.  Note this will overwrite
     any settings from zsh-newuser-install already in the startup file.
     It will not alter any of your other settings, however.

(q)  Quit and do nothing else.
--- Type one of the keys in parentheses ---
```

I don't know about you, but even after using zsh for multiple years, if I ever
see this, it's very rare that I take the time to go through this process every
time: it's daunting and easy to end up with a config file that's just a tiny
bit different on each of your computers leading to tons of frustration.

That's where I see the main advantages of a config framework: providing a much
better out of the box experience for zsh and providing a base to build a
personal config off of.

## Frameworks

### Oh My Zsh

[oh-my-zsh](https://github.com/robbyrussell/oh-my-zsh) is the original (and
probably best-known) ZSH config framework. It provides a core with some basic
config and utilities, tons of plugins which do everything from configure parts
of zsh to add new functionality and tons of themes. It's meant to be as low-
friction as possible and easy for anyone to use.

### Prezto

[Prezto](https://github.com/sorin-ionescu/prezto) was forked from oh-my-zsh a
long time ago as a result of differences in opinion between how the framework
should be managed and what it should include. Prezto includes a minimal core
to load modules, a number of modules meant to provide sane defaults and
conveniences, and a reasonable number of built-in themes.

### zsh-utils

[zsh-utils](https://github.com/belak/zsh-utils) was recently born out of the
frustration of prezto maintenence, and is meant to be a clean slate with only
a few modules designed to make the out of the box zsh experience much better.

### Others

Please note that I've only provided details on the above frameworks because
they're what I'm most familiar with and provide the history behind zsh-utils.
There are a number of other config frameworks which I would be remiss to avoid
mentioning such as [zim](https://github.com/zimfw/zimfw) (a solid framework
originally forked from prezto which aims to be faster and more consistently
maintained).
