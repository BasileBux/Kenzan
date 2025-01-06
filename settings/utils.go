package settings

type IndentationType = uint8

const (
	TABS IndentationType = iota
	SPACES
)

// Could have been a map
func GetIndentationType(t string) IndentationType {
	switch t {
	case "tabs":
		return TABS
	case "spaces":
		return SPACES
	default:
		return TABS
	}
}
