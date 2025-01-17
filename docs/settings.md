# Settings docs

The location of the config files is the following: 
- Linux: `~/.config/kenzan/`
- MacOS: `~/Library/Application Support/kenzan/`
- Windows: `C:\Users\<username>\Appdata\local\kenzan\`

Your config file should be named one of these: 
- settings.json
- user.json
- kenzan.json

The settings are in json as it is really common amongst other editors and I don't mind it.
You can access the defaults in `config_directory/default.json` to look up the global structure. This defaults file is not the one used by the app. If you modify it, it won't have any effect. Its only purpose is to give an example of how the configuration should look. 

This file structure is copied from the [zed docs](https://zed.dev/docs/configuring-zed)


## Padding

- Description: Space between text and edge of the window.
- Setting: `padding`
- Default: 1, 1, 1, 0

#### Options

`integer` values\
Sub-settings: `top`, `right`, `bottom`, `left`

```toml
"padding": {
    "top": 1,
    "right": 1,
    "bottom": 1,
    "left": 0
},
```

## Font Family

- Description: Font family to use in the editor
- Setting: `font_family`
- Default: "" (for info on fallback font, refer to /docs/fonts.md)

#### Options

`string` values. Can be font name or path to the font file. Leave empty for default. 

## Font size

- Description: Font size to use in the editor
- Setting: `font_size`
- Default: 30

#### Options

`integers` values

## Font spacing

- Description: Space between characters
- Setting: `font_spacing`
- Default: 1

#### Options

`integer` values

## Scroll padding

- Description: Number of characters and lines to keep visible around the cursor when scrolling, maintaining a buffer in all directions.
- Setting: `scroll_padding`
- Default: 5

#### Options

positive `integer` values (can be 0)

## Cursor ratio

- Description: Ratio of cursor height to text height.
- Setting: `cursor_ratio`
- Default: 1

#### Options

`integers` values between 0 and 1 with 0 not included = ]0;1]

## Theme

- Description: Editor color theme name
- Setting: `theme`
- Default: "ayu-light"

#### Options

`string` values. Theme name which corresponds to the name of the theme file in ~/.config/kenzan/themes/

## Line numbers

- Description: Section which handles line numbers
- Setting: `line_numbers`

### Show

- Description: Show the line numbers or not
- Setting: `show`
- Default: true

#### Options

`boolean`

### Relative

- Description: Set line numbers as relative or absolute
- Setting: `relative`
- Default: false

#### Options

`boolean`

### Padding left

- Description: Space from left window border to line numbers
- Setting: `padding_left`
- Default: 24

#### Options

`integer` values. If the value is too small or too big, it will just look ugly

### Padding right

- Description: Space from line numbers to text
- Setting: `padding_right`
- Default: 10

#### Options

`integer` values. If the value is too small or too big, it will just look ugly

### Line width

- Description: Width of the line separating line numbers from text
- Setting: `line_width`
- Default: 0

#### Options

Positive `integer` values. Set to 0 to remove line

### Offset current

- Description: Aligns the current line number to the right to 3 digits
- Settings: `offset_current`
- Default: true

#### Options

`boolean`

## Line highlight

- Description: Toggle highlight on current line
- Setting: `line_highlight`
- Default: true

#### Options

`boolean`

## High dpi

- Description: Enable high dpi mode
- Setting: `high_dpi`
- Default: true

#### Options

`boolean` values. True is activated. 

## FPS

- Description: FPS to run the program at
- Settings: `fps`
- Default: 60

#### Options

Non-null positive `integers`. 

## Indentation

- Description: Section which handles indentation options
- Setting: `indentation`

### Type

- Description: Set indentation as spaces or tabs
- Setting: `type`
- Default: "tabs"

#### Options

A `string` of these: 
- "tabs"
- "spaces"

### Size

- Description: Size of tab symbol
- Setting: `size`
- Default: 4

#### Options

Positive non-null `integers`

