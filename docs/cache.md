# Cache docs

This file describes how cache files are formatted and how they are used in the text editor. The cache is a json file stored in system default cache path: 
- Linux: `~/.cache/kenzan/cache.json`
- MacOS: `~/Library/Caches/kenzan/cache.json`
- Windows: `C:\Users\<Username>\AppData\Local\kenzan\cache.json`

## Cached data

### Font path

Why: Retrieving the path to a font from its name is pretty heavy.\
Invalidation: When the user config font has been changed
    - [font name](#font-name)
Name: `font_path`\
Type: `string`\

### Font name

Why: Invalidation of [font path](#font-path)
Invalidation: When this doesn't correspond to the one in the user config file\
    - [self](#font-name)
Name: `font_name`\
Type: `string`\

