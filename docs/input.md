# Input

The input system was written in a really ugly but functional way. The code was unmaintainable 
but worked. The reworked focused on making the input system not as rigid and easily updatable. 

## Principle

An input is this struct: 

```go
type Input struct {
	ModKey  int32 // set to -1 if no mod key
	Key     int32
	Execute func(text *[]string, state *t.ProgramState, style *st.WindowStyle)
}
```
Then, you just have to call the `.set(text, state, style)` method in the loop and then the input is added

## Defaults

Is the file where all default inputs are defined. This is possibly temporary, but right now idk. 

## Movement / arrows

Movement and arrows are different. MoveDirection is only the basic cursor movement in one 
direction and one direction only. Arrows on the other hand, are movement with more features. 
You can move word by word with ctrl + arrow and on the end of a line, the cursor goes to the logical next line. 

## Text

`textInput` is a function which isn't a valid action for an Input. It would have the 
correct prototype but it already handles inputs internally.  

## Actions

Are all non text related actions. 

## Erase

Doesn't suck anymore. The erase action is like the `x` vim motion. It is then used in the
general case in `delete` and `backspace`
