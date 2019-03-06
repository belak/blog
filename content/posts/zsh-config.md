---
title: "ZSH Config Frameworks"
date: 2018-05-02T13:49:22-07:00
tags: ["zsh", "dotfiles", "zsh-utils", "oh-my-zsh", "presto"]
draft: true
---

ZSH config frameworks are a good way to get a better shell experience without spending a ton of time configuring your shell manually. If you're not using one already, you should check them out.

<!--more-->

# Why Use a Config Framework?

Full disclosure: I am currently one of the main maintainers of prezto and the sole maintainer of zsh-utils, so this article *will* have a bias. However, if you see any incorrect points in here, please let me know and I'll happily amend them.

If you've ever started zsh without a config file before, you've probably seen the following mass of text before:

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

Even after using zsh for multiple years, if I see this, it's very rare that I take the time to go through this process every time: it's daunting and almost every time you'll end up with a config file that's just a tiny bit different on each of your computers, leading to tons of frustration.

That's where I see the main advantages of a config framework: providing a much better out of the box experience for zsh and providing a base to build a personal config off of.

## Frameworks

### Oh My Zsh

[oh-my-zsh](https://github.com/robbyrussell/oh-my-zsh) is the original (and probably best-known) ZSH config framework. It provides a core with some basic config and utilities, tons of plugins which do everything from configure parts of zsh to add new functionality and tons of themes. It's meant to be as low-friction as possible and easy for anyone to use.

Oh My Zsh excels at providing a good out of the box experience with tons and tons of plugins and themes. Almost every ZSH plugin manager integrates with it, making it the easiest to get started with.

However it's known for being relatively heavy-weight, not as configurable (outside choosing modules and a theme) and there was a long period where the maintenance status wasn't clear. The creator has written a blog post called [d'Oh My Zsh](https://medium.freecodecamp.org/d-oh-my-zsh-af99ca54212c) outlining why: this is an open source project coded in his free time as *fun*, not as a job. I completely respect and agree with the reasoning behind this decision, but it was frustrating when I opened a number of PRs against OMZ and had to wait months for them to be looked at, let alone merged.

### Prezto

[Prezto](https://github.com/sorin-ionescu/prezto) was forked from oh-my-zsh a long time ago as a result of differences in opinion between how the framework should be managed and what it should include. Prezto includes a minimal core to load modules, a number of modules meant to provide sane defaults and conveniences, and a reasonable number of built-in themes.

### zsh-utils

[zsh-utils](https://github.com/belak/zsh-utils) was recently born out of the frustration of prezto maintenance, and is meant to be a clean slate with only a few modules designed to make the out of the box zsh experience much better.

### Others

There are other config frameworks which I would be remiss to avoid mentioning such as [zim](https://github.com/zimfw/zimfw) (a solid framework originally forked from prezto which aims to be faster and more consistently maintained). However I've only provided details on the above frameworks because they're what I'm most familiar with and provide the history behind zsh-utils. 

## What's Hard About Maintaining a ZSH Framework

There are a number of frustrations I've run into with maintaining prezto. Some of these cannot be avoided, some can. None of the frameworks currently mentioned solve these in a consistent way.

Problems which can be solved:

- Interconnected modules make it hard to make big changes. This can be solved by clearly defining interfaces between modules. As an example, the git module in prezto tries really hard to do this.
- Module dependencies aren't well defined. This can be solved by explicitly defining relationships between modules and building infrastructure to handle those relationships. This can unfortunately add a huge layer of complexity.

Problems which don't have a clear solution:

- Varying Module quality and maintenance status
- How to handle conflicting feature requests
- Repo owners going MIA with conflicting ideas about the future

## So, What's the Best?

At least for me, zsh-utils is the right choice. I know it will be maintained as long as I'm using it (and I'd like to add some additional contributors in the future), it's lightweight/fast, and it's simple.

However, that doesn't mean it's right for everyone. I don't plan on adding tons of features, modules, or themes.

Oh My Zsh is another good option. It's easy to set up, there are tons and tons of modules and the new maintainer [@mcornella](https://github.com/mcornella) has been doing a very good job of working through the backlog of PRs and issues that built up over the last few years.

I personally view prezto as a better option for the more technically minded. It's more configurable, still has a reasonable number of modules, and is actively maintained by a number of volunteers. However, some of it is currently in limbo because the owner made a number of decisions a while ago and isn't around as much any more. A large portion of prezto is saner defaults but some of these defaults are fairly opinionated.