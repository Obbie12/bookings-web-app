package models

// TemplateData holsd data sent from handlers to Templates
type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	flash     string
	Warning   string
	Error     string
}
