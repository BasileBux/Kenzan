package renderer

import (
	st "github.com/basilebux/kenzan/settings"
	t "github.com/basilebux/kenzan/types"
	rl "github.com/gen2brain/raylib-go/raylib"
	tree_sitter "github.com/tree-sitter/go-tree-sitter"
	tree_sitter_c "github.com/tree-sitter/tree-sitter-c/bindings/go"
)

func highlight(text *string, color *rl.Color, cursor *t.TextRenderCursor, state *t.ProgramState, style *st.WindowStyle) {
	tmp := calculateOffset(cursor, text, state, style)
	state.Cache.Syntax = append(state.Cache.Syntax, t.SyntaxCache{Text: *text, Color: color, Cursor: tmp})
}

func RenderText(lang t.Language, text *string, state *t.ProgramState, style *st.WindowStyle) {
	if lang != t.NONE {

		if state.Update.SyntaxHighlight {

			code := ([]byte)(*text)
			parser := tree_sitter.NewParser()
			defer parser.Close()

			switch lang {
			case t.C:
				parser.SetLanguage(tree_sitter.NewLanguage(tree_sitter_c.Language()))
			}

			tree := parser.Parse(code, nil)
			defer tree.Close()

			root := tree.RootNode()

			switch lang {
			case t.C:
				state.HighlightErr = syntaxHighlightingC(root, code, state, style)
			}
		}
		if state.HighlightErr != nil {
			noSyntaxHighlight(text, &state.Nav.ScrollOffset, style)
		} else {
			renderHighlight(state, style)
		}
	} else {
		noSyntaxHighlight(text, &state.Nav.ScrollOffset, style)
	}
}
