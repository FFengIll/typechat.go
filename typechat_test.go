package typechat

import (
	"strings"
	"testing"
)

// 用于测试的示例结构体
type TestResponse struct {
	Message string
	Code    int
}

func TestCustomTranslate(t *testing.T) {
	tests := []struct {
		name     string
		template string
		prompt   string
		model    interface{}
		want     string
		wantErr  bool
	}{
		{
			name:     "basic translation",
			template: "",
			prompt:   "Hello",
			model:    TestResponse{},
			want:     "Hello\nRespond strictly with JSON",
			wantErr:  false,
		},
		{
			name:     "custom template",
			template: "Custom: %s\nStruct: %s",
			prompt:   "Test",
			model:    TestResponse{},
			want:     "Custom: Test\nStruct:",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CustomTranslate(tt.template, tt.prompt, tt.model)
			if (err != nil) != tt.wantErr {
				t.Errorf("Translate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !strings.Contains(got, tt.want) {
				t.Errorf("Translate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTranslate(t *testing.T) {
	got, err := Traslate("Test prompt", TestResponse{})
	if err != nil {
		t.Errorf("DefaultTraslate() error = %v", err)
		return
	}
	if !strings.Contains(got, "Test prompt") {
		t.Errorf("DefaultTraslate() = %v, want to contain 'Test prompt'", got)
	}
}

func TestRecoverStructDef(t *testing.T) {
	tests := []struct {
		name  string
		input interface{}
		want  string
	}{
		{
			name:  "nil input",
			input: nil,
			want:  "",
		},
		{
			name:  "non-struct input",
			input: "string",
			want:  "",
		},
		{
			name:  "valid struct",
			input: TestResponse{},
			want:  "type TestResponse struct",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RecoverStructDef(tt.input)
			if tt.want == "" && got != "" {
				t.Errorf("RecoverStructDef() = %v, want empty string", got)
			}
			if tt.want != "" && !strings.Contains(got, tt.want) {
				t.Errorf("RecoverStructDef() = %v, want to contain %v", got, tt.want)
			}
		})
	}
}

func TestTranslator_filterModel(t *testing.T) {
	translator := NewTranslatorWithTemplate("")
	models := []interface{}{
		TestResponse{},
		TestResponse{}, // 重复的结构体
		"not a struct", // 非结构体
	}

	filtered := translator.filterModel(models)
	if len(filtered) != 1 {
		t.Errorf("filterModel() returned %d models, want 1", len(filtered))
	}
}

func TestTranslator_Generate(t *testing.T) {
	translator := NewTranslatorWithTemplate("")
	prompt := "Test prompt"
	model := TestResponse{}

	result, err := translator.Generate(prompt, model)
	if err != nil {
		t.Errorf("Generate() error = %v", err)
		return
	}

	expectedContents := []string{
		prompt,
		"TestResponse",
		"Message",
		"Code",
	}

	for _, expected := range expectedContents {
		if !strings.Contains(result, expected) {
			t.Errorf("Generate() result doesn't contain expected content: %s", expected)
		}
	}
}
