---
permalink: "dark-mode-ubuntu-i3"
layout: post
title: "Toggling light/dark color scheme for i3 on Ubuntu"
date: "2024-07-16 00:30"

tags: project
---

I didn't find an easy way to switch between light and dark mode in
[i3](https://i3wm.org) on Ubuntu using Gnome.  I really prefer light schemes in
a well-lit office, and a darker screen when sitting at night. On that note, I
really appreciate [redshift](https://github.com/jonls/redshift).

This is how I approached **easily changing between light and dark mode**:

- Figure out how to toggle the light/dark preference programmatically
- Make sure all software reacts to it
  - Preferably: supporting the system settings
  - Fallback: hook into the toggling and adapt the specific programs one-by-one
    to change themes etc.

## Changing theme programmatically

Creating a clickable button in the
[i3blocks](https://github.com/vivien/i3blocks) menu is easy:

```
# in ~/.i3blocks.conf

[colorscheme]
# font awesome sun (f185)
full_text=
command=~/bin/color-scheme>/dev/null
```

A left click toggles between light and dark, here's the [`~/bin/color-scheme` script](https://github.com/chelmertz/dotfiles/tree/master/bin/color-scheme):

```sh
#!/bin/sh

current=$(gsettings get org.gnome.desktop.interface color-scheme)
if [ $? -ne 0 ]; then echo "Could not fetch current color-scheme" >&2; exit 1; fi
echo "Current color-scheme: $current"

case "$current" in
	"'prefer-dark'") new_color_scheme="prefer-light" ;;
	"'prefer-light'"|"'default'") new_color_scheme="prefer-dark" ;;
esac

echo "New color-scheme: $new_color_scheme"

if ! $(gsettings set org.gnome.desktop.interface color-scheme "$new_color_scheme"); then
	echo "Could not set new color-scheme" >&2
	exit 1
fi
```

## Make software react to the system setting

Software that can be configured to adapt to the system wide settings:

- Firefox
- [gnome-console](https://gitlab.gnome.org/GNOME/console)
  - The default terminal, gnome-terminal, doesn't support listening in on the system wide color scheme.
  - The apt package is called `gnome-console` but the terminal's program is called `kgx`.
  - After installing it, do a `sudo update-alternatives --install /usr/bin/x-terminal-emulator x-terminal-emulator /usr/bin/kgx 1` and `sudo update-alternatives --set x-terminal-emulator /usr/bin/kgx`, to make it the default terminal.
  - Also: since I'm using i3 and it doesn't have kgx in its allowlist in /usr/bin/i3-sensible-terminal, I also `export TERMINAL=kgx` in my .zshrc
- [bat](https://github.com/sharkdp/bat)
  - Use the `base16` or `ansi` theme (either in a configuration file, or `alias bat="bat --theme=base16"`)
- [doom emacs](https://github.com/doomemacs/doomemacs)
  - Using the [auto-dark-emacs](https://github.com/LionyxML/auto-dark-emacs) plugin, [configured with a few lines](https://github.com/chelmertz/dotfiles/commit/21944f4daa5fb185e5724020748f20b5e7d1b603)
- Slack
- [element](https://element.io/), a matrix chat client
- [elly](https://github.com/chelmertz/elly), my tool for keeping track of Github pull requests (yes, this is a shameless plug; no, it doesn't need to be configured, because media queries)
- VS Code
  - Searching for "preferred" in settings, allows you to apply these settings (ends up in ~/.config/Code/User/settings.json):
    ```
    "workbench.preferredDarkColorTheme": "No Happiness in Colors Theme",
    "workbench.preferredLightColorTheme": "Subtle Monochrome (Light)",
    "window.autoDetectColorScheme": true
    ```
  - I've found that using monochrome themes makes me skip less of the code when reading it. So far, [No happiness in Colors](https://vscodethemes.com/e/notoroszbig.theme-nohappinessincolors/no-happiness-in-colors-theme) is a great dark theme, and [Subtle Monochrome (light)](https://github.com/anotherglitchinthematrix/monochrome/) is great as a light theme. Don't forget to also [turn off rainbow bracket colors](https://github.com/chelmertz/dotfiles/commit/cf68bf6163e9f8b639eba06e56f9175b4728fd0e).

Left to fix:

- Firefox' [dark reader extension](https://github.com/darkreader/darkreader)
  - It has a "Use system color scheme" but toggling the system scheme doesn't make dark reader toggle the plugin's status
- IDEA
  - [Not supported yet on Linux](https://youtrack.jetbrains.com/issue/IJPL-54591/Implement-IDE-theme-sync-with-OS-on-Linux) and the linked plugin doesn't do it for me
    - [kant](https://github.com/abrookins/kant) is one of few monochromatic color schemes available


## Previous attempt: "signal-lamp"

I tried, and failed, to solve this problem previously, in [https://github.com/chelmertz/signal-lamp](https://github.com/chelmertz/signal-lamp).

I tried to take ownership of the current state/scheme myself, which would require hooks/APIs for every program I'd like to change.

A particularly bad example: I never found a way to change the theme of all open gnome-terminal windows, so I used [wmctrl](https://linux.die.net/man/1/wmctrl) and [xdotool](https://linux.die.net/man/1/wmctrl) in a loop 🙈 It turned out super buggy and the code wasn't even pretty to read. A do-over was necessary.

## Elsewhere: `darkman`

I just found Darkman ([intro](https://whynothugo.nl/journal/2022/02/25/introducing-darkman-v1.0.0/), [repo](https://gitlab.com/WhyNotHugo/darkman)) but it is a bit intrusive. I like the dbus parts but having a separate service for this, and the scripts for every application, is a bit like the "signal-lamp" attempt and rubs me the wrong way (configuration rabbit hole).


