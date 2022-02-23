package models

import "github.com/mehul-tandel/beachhouse_booking/internal/forms"

// TemplateData holds data sent from handlers to templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{} //for unknown type of data
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
	Form      *forms.Form
}
