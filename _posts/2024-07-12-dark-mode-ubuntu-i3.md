---
layout: post
title: "Toggling light/dark color scheme"
date: 2024-07-13 12:00

tags: project
comments: true
published: false
---

********
uh, take a look at https://whynothugo.nl/journal/2022/02/25/introducing-darkman-v1.0.0/ and its related links in https://github.com/doomemacs/doomemacs/issues/6424

and also:
- add a note https://github.com/orgs/community/discussions/16925 or something into signal-lamp, archive it and mention this blog-post (once it's published)
********

I didn't find an easy way to switch between light and dark mode in
[i3](https://i3wm.org) on Ubuntu using Gnome. I'm not sure if it's the age, but
I really prefer light schemes in the Swedish summer days, and a darker screen
when sitting at night. On that note, I really appreciate
[redshift](https://github.com/jonls/redshift).

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

A right click toggles between light and dark, while a left click brings up a
[zenity](https://help.gnome.org/users/zenity/stable/) menu (that's the
`"$BLOCK_BUTTON" -eq 3` part below).

Here's the [`~/bin/color-scheme` script](https://github.com/chelmertz/dotfiles/tree/master/bin/color-scheme):

```sh
#!/bin/sh

current=$(gsettings get org.gnome.desktop.interface color-scheme)
if [ $? -ne 0 ]; then echo "Could not fetch current color-scheme" >&2; exit 1; fi
echo "Current color-scheme: $current"

current_is_default="FALSE"
current_is_light="FALSE"
current_is_dark="FALSE"
case "$current" in
	"'prefer-dark'") current_is_dark="TRUE" ;;
	"'prefer-light'") current_is_light="TRUE" ;;
	"'default'") current_is_default="TRUE" ;;
esac

if [ "$BLOCK_BUTTON" -eq 3 ]; then
	# right click i3blocks icon to toggle between dark & light,
	# don't touch "default" since it has no antonym
	case "$current" in
		"'prefer-dark'") new_color_scheme="prefer-light" ;;
		"'prefer-light'") new_color_scheme="prefer-dark" ;;
	esac
else 
	new_color_scheme=$(zenity --list --title "Color scheme" --column= --column="Color scheme" --radiolist $current_is_default default $current_is_light prefer-light $current_is_dark prefer-dark)
	if [ $? -ne 0 ]; then echo "Could not get new mode from zenity" >&2; exit 1; fi
fi

echo "New color-scheme: $new_color_scheme"

if ! $(gsettings set org.gnome.desktop.interface color-scheme "$new_color_scheme"); then
	echo "Could not set new color-scheme" >&2
	exit 1
fi
```

In particular, this makes use of `gsettings` configuration entry called
`org.gnome.desktop.interface color-scheme`.

In the script, I tried to avoid my go-to `set -euo pipefail` by handling the
different errors manually with a more describing error message. Maybe I'm
infected by mostly coding in go.

![i3blocks button](/assets/colorscheme_i3blocks.png)

## Make software react to the system setting

Software that can be configured to the system wide settings:

- Firefox
- [gnome-console](https://gitlab.gnome.org/GNOME/console)
  - The default terminal, gnome-terminal, doesn't support listening in on the system wide color scheme.
    - The apt package is called `gnome-console` but the terminal's program is called `kgx`.
      - After installing it, do a `sudo update-alternatives --install /usr/bin/x-terminal-emulator x-terminal-emulator /usr/bin/kgx 1` and `sudo update-alternatives --set x-terminal-emulator /usr/bin/kgx`, to make it the default terminal.
      - Also: since I'm using i3 and it doesn't have kgx in its allowlist in /usr/bin/i3-sensible-terminal, I also `export TERMINAL=kgx` in my .zshrc


Left to fix:

- Firefox' [dark reader extension](https://github.com/darkreader/darkreader)
- VS Code
- Cursor
- bat
- IDEA
- Doom Emacs


## Previous attempt: "signal-lamp"

I tried, and failed, to solve this problem previously, in https://github.com/chelmertz/signal-lamp

I tried to take ownership of the current state/scheme myself, which would require hooks/APIs for every program I'd like to change.

A particularly bad example: I never found a way to change the theme of all open gnome-terminal windows, so I used [wmctrl](https://linux.die.net/man/1/wmctrl) and [xdotool](https://linux.die.net/man/1/wmctrl) in a loop 🙈 It turned out super buggy and the code wasn't even pretty to read. A do-over was necessary.
