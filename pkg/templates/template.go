package templates

// Template represents a customizable project template structure.
type Template struct {
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Files       map[string]string `json:"files"`
}

// NewTemplate creates a new project template with default values.
func NewTemplate(name, description string) *Template {
	return &Template{
		Name:        name,
		Description: description,
		Files:       make(map[string]string),
	}
}

// AddFile adds a file to the template with its content.
func (t *Template) AddFile(filePath, content string) {
	t.Files[filePath] = content
}
