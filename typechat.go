package typechat

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.StandardLogger()

// Translate is a top-level convenience function
func Translate(template string, prompt string, models ...interface{}) (string, error) {
	t := NewTranslator(template)
	return t.Generate(prompt, models...)
}

func DefaultTraslate(prompt string, models ...interface{}) (string, error) {
	t := NewDefaultTranslator()
	return t.Generate(prompt, models...)
}

// Translator is the main translator struct
type Translator struct {
	template string
}

const defaultTemplate = `%s
Respond strictly with JSON. The JSON should be compatible with the Go struct Response from the following:
` + "```" + `go
%s
` + "```"

// NewTranslator creates a new translator instance
func NewTranslator(template string) *Translator {
	if template == "" {
		template = defaultTemplate
	}
	return &Translator{
		template: template,
	}
}

func NewDefaultTranslator() *Translator {
	return NewTranslator("")
}

// filterModel filters and deduplicates models
func (t *Translator) filterModel(models []any) []interface{} {
	seen := make(map[string]bool)
	var result []interface{}

	for _, model := range models {
		typ := reflect.TypeOf(model)
		if typ.Kind() == reflect.Ptr {
			typ = typ.Elem()
		}

		// Only process struct types
		if typ.Kind() != reflect.Struct {
			continue
		}

		name := typ.Name()
		if !seen[name] {
			seen[name] = true
			result = append(result, model)
		} else {
			log.Warnf("found duplicated model, check your code: %s", name)
		}
	}

	return result
}

// RecoverStructDef recovers struct definition from an interface
func RecoverStructDef(i interface{}) string {
	if i == nil {
		return ""
	}

	typ := reflect.TypeOf(i)
	if typ.Kind() == reflect.Ptr {
		typ = typ.Elem()
	}

	if typ.Kind() != reflect.Struct {
		return ""
	}

	def := fmt.Sprintf("type %s struct {\n", typ.Name())
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		def += fmt.Sprintf("\t%s %s `json:\"%s\"`\n",
			field.Name,
			field.Type.String(),
			strings.ToLower(field.Name))
	}
	def += "}"
	return def
}

// toConstraint converts models to Go struct definition strings
func (t *Translator) toConstraint(models []interface{}) string {
	var definitions []string

	for _, model := range models {
		res := RecoverStructDef(model)
		logrus.Debugf("res: %s\n", res)
		definitions = append(definitions, res)
	}

	return strings.Join(definitions, "\n")
}

// Generate generates the final prompt text
func (t *Translator) Generate(prompt string, models ...any) (string, error) {
	if len(models) > 0 {
		models = t.filterModel(models)
	}

	constraint := t.toConstraint(models)
	return fmt.Sprintf(t.template, prompt, constraint), nil
}
