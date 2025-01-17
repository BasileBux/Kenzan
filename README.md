# Kenzan text editor

A small personal project to write a text editor from 0 in Go and [Raylib](https://www.raylib.com/). 
Cross-platform (only tested on Linux tho). This is a work in progress and pretext to learn a bunch of stuff.

The current features are really basic. Move with arrows, write text, delete text, open files, 
write files, ... If you are wondering what I will implement next, take a look at my [todos](#todo).

I also implemented syntax highlighting using [Tree-sitter](https://tree-sitter.github.io/tree-sitter/). 
For, now there is only `c`, but I want to implement other languages I know when I have time. 
If you want to throw a quick pull request and implement syntax highlighting for a language 
you like, you are more than welcome to do so. (I find it pretty boring actually)

> [!WARNING]
> Unicode is not supported for now. I have solved my issue so it might come in the future 
but not yet. So if you type non-ascii characters, they will simply be ignored

## Fallback font

As fallback font, I chose the [JetBrains Mono](https://www.jetbrains.com/lp/mono/). I do not own 
the font in any way. The font is licensed under the SIL Open Font License. For more information, 
you can find the font and license in the [/fonts/JetBrainsMono/](https://github.com/BasileBux/Kenzan/blob/main/fonts/JetBrainsMono/) folder.

### Supported language syntax highlighting

- c

## Usage

You can build it from source. You will just need to have [golang](https://go.dev/doc/install) 
installed and then you can simply clone this repo, go into it and execute `go run install/install.go`
```bash
git clone https://github.com/BasileBux/Kenzan.git
cd Kenzan
go run install/install.go
```

Or if you want to do it manually, you need [golang](https://go.dev/doc/install) installed 
on your machine. Clone the repo, execute `go mod tidy` in the directory and build it with go. 
You will also need to create a directory in your system's default cache directory. 

To run it, Just execute the program and provide the path to the file you want to edit. 
If you don't give a file, it will open a blank file which won't be able to be saved. 
The text editor cannot create a new file yet.

## Issues

- Weird padding on certain machines ? Env vars ? Debug build ?

> [!WARNING]
> The padding error is really cryptic. However, it seems to be a really weird error tied 
to the monitor or some shit like that. This should be a known issue but no fix is needed.
