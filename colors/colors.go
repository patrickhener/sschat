package colors

var (
	styles       = []*Style{}
	secretStyles = []*Style{}
)

// Style will represent a terminal prompt style
type Style struct {
	Name  string
	Apply func(string) string
}

// GetStyle will return a style based upon a color string
// If it fails it will return an error
func GetStyle(color string) (*Style, error) {
	for i, sty := range styles {
		if sty.Name == color {
			return styles[i], nil
		}
	}
	for i, sty := range secretStyles {
		if sty.Name == color {
			return secretStyles[i], nil
		}
	}

	return nil, nil

}
