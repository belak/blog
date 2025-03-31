---
title: "An Intro to zsh-utils"
date: 2019-03-12T11:15:00-07:00
tags: ["dotfiles", "zsh", "zsh-utils"]
hide_updated: true
---

ZSH is an extremely powerful tool which can be extremely useful but hard to configure. [zsh-utils](https://github.com/belak/zsh-utils) aims to fix this. It's a small configuration framework which aims to provide a base to work off of without getting in your way if you want to go deeper later.

<!--more-->

## Why Use a Config Framework?

If you've ever started zsh without a config file before, you've probably seen the following mass of text before:

```
Please pick one of the following options:

(1)  Configure settings for history, i.e. command lines remembered
     and saved by the shell.  (Recommended.)

(2)  Configure the new completion system.  (Recommended.)

(3)  Configure how keys behave when editing command lines.  (Recommended.)

(4)  Pick some of the more common shell options.  These are simple "on"
     or "off" switches controlling the shell's features.

(0)  Exit, leaving the existing ~/.zshrc alone.

(a)  Abort all settings and start from scratch.  Note this will overwrite
     any settings from zsh-newuser-install already in the startup file.
     It will not alter any of your other settings, however.

(q)  Quit and do nothing else.
--- Type one of the keys in parentheses ---
```

Even after using ZSH for multiple years, if I see this, it's rare that I take the time to go through this process every time: it's daunting and almost every time you'll end up with a config file that's different on each of your computers, leading to tons of frustration.

That's where I see the main advantages of a config framework: providing a much better out of the box ZSH experience and providing a base to build a personal config off of.

## What Is This?

Taken from the project README, [zsh-utils](https://github.com/belak/zsh-utils) is "a minimal set of ZSH plugins designed to be low-friction and low-complexity."

After getting frustrated with the maintenance of prezto and oh-my-zsh, I decided to write my own small config framework, in a similar vein to some of the starter kits.

## Why Another Framework?

There were a number of main problems I noticed during my time maintaining prezto:

- Most ZSH config frameworks contain a bundled plugin manager along with a number of modules. It would make more sense to officially support an external plugin manager and focus development on the plugins.
- It takes a ton of work to get a module ready for prezto, even if it's small.
- A number of large frameworks set up strange defaults.
- Many plugins end up being nothing more than a small wrapper around an external repo. Most of this work could be handled by plugin managers.
- Some plugin managers build non-standard ways to configure ZSH (prezto puts a *ton* of config values in zstyle and oh-my-zsh has their own custom theme format).
- They just do too much. All I'm really looking for is a small wrapper around the built-in ZSH features.
- So many modules have way too many options, often because they're trying to do too many things.

That being said, there are a number of strengths to large configuration frameworks as well.

- Plugins can work well together. As an example, prezto's git plugin provides a git-info command which can be used by prompts to display git information without the prompts having to implement this themselves.
- Having a number of plugins bundled together in one place makes it simple to enable the ones you want without having to do a bunch of additional work.

These are not really something I’m looking for. I want to use a shell as a shell and keep everything as simple as possible so it doesn’t get overwhelming.

## Ok, So Why This One?

This project aims to solve some of the problems faced by larger frameworks by focusing on the following clear goals:

- Loadable by any compatible plugin manager
- Plugins will be kept small, relatively inflexible, well documented, and well organized
- No external dependencies without good reason
- Focus on improving the existing experience rather than expanding it

Each plugin has a focused purpose:

- `completion` configures and loads basic tab completion (imagine bash tab completion on steroids). It's also recommended to use an additional set of [zsh-completions](https://github.com/zsh-users/zsh-completions) because this plugin only sets up the completion system and doesn't provide completions for common utils.
- `editor` sets up some keybinds to fix some rough edges. It adds a few missing default key-binds such as Home/End and aims to add a few Vim features to the vi keybind set.
- `history` configures and loads the shell history systems.
- `prompt` runs initializes the built-in zsh theming system and provides a place for maintainers to put the prompt they use.
- `utility` makes it easier to switch between operating systems by adding aliases for common operations so you don't need to remember different commands across platforms.

That's it! There are no plugins for programming languages. There are no custom formats. And there is no configuration outside loading the plugins and optionally overriding ZSH settings.

## Sounds Great! How Do I Start?

If this is something that sounds useful to you, it's fairly easy to get started.

Simply replace your `.zshrc` with the following snippet. This is [copied directly from the setup](https://github.com/belak/zsh-utils/#recommended-installation) in the zsh-utils README.

It downloads [antigen](https://github.com/zsh-users/antigen.git), a simple plugin manager, straight from the source if it doesn't exist and loads all the zsh-utils plugins along with a few other commonly used external plugins.

```sh
[[ ! -d "$HOME/.antigen" ]] && git clone https://github.com/zsh-users/antigen.git "$HOME/.antigen"
source "$HOME/.antigen/antigen.zsh"

# Set the default plugin repo to be zsh-utils
antigen use belak/zsh-utils

# Specify completions we want before the completion module
antigen bundle zsh-users/zsh-completions

# Specify plugins we want
antigen bundle editor
antigen bundle history
antigen bundle prompt
antigen bundle utility
antigen bundle completion

# Specify additional external plugins we want
antigen bundle zsh-users/zsh-syntax-highlighting

# Load everything
antigen apply

# Set any settings or overrides here
prompt belak
bindkey -e
```

Alternatively, [my dotfiles](https://github.com/belak/dotfiles/blob/master/zshrc) are a good resource and starting point.

## Now What?

Just use ZSH! If you have any problems, feel free to [file an issue](https://github.com/belak/zsh-utils/issues/new) or open a pull request if you're feeling up for it.

## Discussion

Want to talk more about this? There are discussion threads at the following locations:

- [lobste.rs](https://lobste.rs/s/6hlx5f/intro_zsh_utils)
- [Hacker News](https://news.ycombinator.com/item?id=19371463)

Thanks for reading!
