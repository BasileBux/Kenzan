# Font

This document describes how the font system works. The goal is that the user gives a font name in the config file and the editor loads it with the correct file path. 

The goal is to cache the font path and reload it every time the settings file is changed. Getting the file path from the font name, is really heavy so it should be made the least time possible. 

The only font types supported are `.ttf` and `.otf`

The fallback font is located in `~/.config/kenzan/fonts/`. I chose the [GeistMono font](https://vercel.com/font). I do not own the font, I just really enjoy it. You can find the license to that font under `./fonts/GeistMono/OFL.txt` in the repo. 

## Linux

We use the tool `fontconfig` and the command `fc-list`, this gives a list of `file path`: `font specs`

## macOS

We won't use `mdfind` as I don't find it really useful. On macOS, fonts can be located: 
- `/System/Library/Fonts/` 
- `/System/Library/Fonts/Supplemental/`
- `/Library/Fonts/`
- `~/Library/Fonts/`
This won't allow to access fonts installed by 3rd party software by name. But path is always available.\
I use `mdls` to get info on the fonts. This gives XML which I have to parse and do stuff with. This is a really long process. 

## Window

Fonts can be in two folders: 
- C:\Windows\Fonts
- C:\Users\<Username>\AppData\Local\Microsoft\Windows\Fonts

So search both paths only. 

## Choosing default style

Default style should be 400. But this isn't present in every font metadata. We are going to make some sort of cascade search. So first, we try if the style of the font is "Regular". If it has regular and something else, we search "Regular", "Medium". If we don't find anything in the font metadata, we search a filename with "regular", then "medium". If none is found, we take the first one in the list so basically random. 
