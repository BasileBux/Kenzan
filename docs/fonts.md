# Font

This document describes how the font system works. The goal is that the user gives a font name in the config file and the editor loads it with the correct file path. 

The goal is to cache the font path and reload it every time the settings file is changed. Getting the file path from the font name, is really heavy so it should be made the least time possible. 

## Linux

We use the tool `fontconfig` and the command `fc-list`, this gives a list of `file path`: `font specs`

## macOS

We use `mdfind` but I cannot test this because I don't have a mac. 

## Window

Fonts can be in two folders: 
- C:\Windows\Fonts
- C:\Users\<Username>\AppData\Local\Microsoft\Windows\Fonts

So search both paths only. 
