package openai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

type ChatCompletion interface {
	GetMessage() *Message
}

type Error struct {
	Message string `json:"message"`
	Type    string `json:"type"`
	Param   string `json:"param"`
	Code    any    `json:"code"`
}

func (e *Error) Error() string {
	return "openai: " + e.Message
}

type Models struct {
	Data []struct {
		ID         string            `json:"id"`
		Object     string            `json:"object"`
		OwnedBy    string            `json:"owned_by"`
		Permission []json.RawMessage `json:"permission"`
	} `json:"data"`
	Object string `json:"object"`
}

type Message struct {
	Role       string    `json:"role"`
	Content    *Content  `json:"content,omitempty"`
	Name       string    `json:"name,omitempty"`
	ToolCallID string    `json:"tool_call_id,omitempty"`
	ToolCalls  ToolCalls `json:"tool_calls,omitempty"`
}

type Messages []*Message

type Content struct {
	Text  string
	Parts Parts
}

func (c *Content) MarshalJSON() ([]byte, error) {
	if c == nil || (c.Text == "" && c.Parts == nil) {
		return json.Marshal(nil)
	}
	if c.Text != "" {
		return json.Marshal(c.Text)
	}
	return json.Marshal(c.Parts)
}

func (c *Content) UnmarshalJSON(data []byte) error {
	decoder := newDecoder(data)
	tok, _ := decoder.Token()
	if tok == nil {
		return nil
	}
	switch tokVal := tok.(type) {
	case string:
		return json.Unmarshal(data, &c.Text)
	case json.Delim:
		if tokVal == '[' {
			return json.Unmarshal(data, &c.Parts)
		}
	}
	return fmt.Errorf("cannot unmarshal %s into Go value of type Content", tokenType(tok))
}

type Part struct {
	Type     string    `json:"type"`
	Text     string    `json:"text,omitempty"`
	ImageUrl *ImageUrl `json:"image_url,omitempty"`
}

type Parts []*Part

type ImageUrl struct {
	Url    string `json:"url"`
	Detail string `json:"detail,omitempty"`
}

type Completion struct {
	ID                string `json:"id"`
	Model             string `json:"model"`
	Object            string `json:"object"`
	Created           int    `json:"created"`
	SystemFingerprint string `json:"system_fingerprint"`
	Choices           []struct {
		Index        NullableType[int]    `json:"index"`
		Message      *Message             `json:"message"`
		FinishReason NullableType[string] `json:"finish_reason"`
		LogProbs     json.RawMessage      `json:"logprobs"`
	} `json:"choices"`
	Usage struct {
		PromptTokens     int `json:"prompt_tokens"`
		CompletionTokens int `json:"completion_tokens"`
		TotalTokens      int `json:"total_tokens"`
	} `json:"usage"`
}

func (c *Completion) GetMessageByIndex(idx int) (message *Message) {
	if idx > len(c.Choices)-1 {
		return new(Message)
	}
	return c.Choices[idx].Message
}

func (c *Completion) GetMessageContentByIndex(idx int) string {
	message := c.GetMessageByIndex(idx)
	if message.Content == nil {
		return ""
	}
	return strings.TrimSpace(message.Content.Text)
}

func (c *Completion) GetMessageRoleByIndex(idx int) string {
	message := c.GetMessageByIndex(idx)
	return strings.TrimSpace(message.Role)
}

func (c *Completion) GetFinishReasonByIndex(idx int) string {
	if idx > len(c.Choices)-1 {
		return ""
	}
	if finishReason := c.Choices[idx].FinishReason; finishReason.IsNull() {
		return ""
	} else {
		return strings.TrimSpace(finishReason.Value())
	}
}

func (c *Completion) GetToolCallsByIndex(idx int) ToolCalls {
	message := c.GetMessageByIndex(idx)
	return message.ToolCalls
}

func (c *Completion) GetMessage() *Message      { return c.GetMessageByIndex(0) }
func (c *Completion) GetMessageContent() string { return c.GetMessageContentByIndex(0) }
func (c *Completion) GetMessageRole() string    { return c.GetMessageRoleByIndex(0) }
func (c *Completion) GetFinishReason() string   { return c.GetFinishReasonByIndex(0) }
func (c *Completion) GetToolCalls() ToolCalls   { return c.GetToolCallsByIndex(0) }
func (c *Completion) GetPromptTokens() int      { return c.Usage.PromptTokens }
func (c *Completion) GetCompletionTokens() int  { return c.Usage.CompletionTokens }
func (c *Completion) GetTotalTokens() int       { return c.Usage.TotalTokens }

type Chunk struct {
	ID                string `json:"id"`
	Model             string `json:"model"`
	Object            string `json:"object"`
	Created           int    `json:"created"`
	SystemFingerprint string `json:"system_fingerprint"`
	Choices           []struct {
		Index        NullableType[int]    `json:"index"`
		Delta        *Message             `json:"delta"`
		FinishReason NullableType[string] `json:"finish_reason"`
		LogProbs     json.RawMessage      `json:"logprobs"`
	} `json:"choices"`
}

func (c *Chunk) GetDeltaByIndex(idx int) (delta *Message) {
	if idx > len(c.Choices)-1 {
		return new(Message)
	}
	return c.Choices[idx].Delta
}

func (c *Chunk) GetDeltaContentByIndex(idx int) string {
	delta := c.GetDeltaByIndex(idx)
	if delta.Content == nil {
		return ""
	}
	return strings.TrimSpace(delta.Content.Text)
}

func (c *Chunk) GetDeltaRoleByIndex(idx int) string {
	delta := c.GetDeltaByIndex(idx)
	return strings.TrimSpace(delta.Role)
}

func (c *Chunk) GetFinishReasonByIndex(idx int) string {
	if idx > len(c.Choices)-1 {
		return ""
	}
	if finishReason := c.Choices[idx].FinishReason; finishReason.IsNull() {
		return ""
	} else {
		return strings.TrimSpace(finishReason.Value())
	}
}

func (c *Chunk) GetToolCallsByIndex(idx int) ToolCalls {
	delta := c.GetDeltaByIndex(idx)
	return delta.ToolCalls
}

func (c *Chunk) GetDelta() *Message      { return c.GetDeltaByIndex(0) }
func (c *Chunk) GetDeltaContent() string { return c.GetDeltaContentByIndex(0) }
func (c *Chunk) GetDeltaRole() string    { return c.GetDeltaRoleByIndex(0) }
func (c *Chunk) GetFinishReason() string { return c.GetFinishReasonByIndex(0) }
func (c *Chunk) GetToolCalls() ToolCalls { return c.GetToolCallsByIndex(0) }

type Stream struct {
	C     <-chan *Chunk
	error <-chan error
}

func (s *Stream) Err() error {
	select {
	case streamErr := <-s.error:
		if streamErr != nil {
			return fmt.Errorf("stream: %w", streamErr)
		}
	default:
	}
	return nil
}

func (s *Stream) Close() error {
	for range s.C {
	}
	return s.Err()
}

func (s *Stream) GetMessage() (message *Message) {
	var (
		messageContentBuilder     strings.Builder
		toolCallArgumentsBuilders []strings.Builder
	)
	message = new(Message)
	for chunk := range s.C {
		if role := chunk.GetDeltaRole(); role != "" {
			message.Role = role
		}
		messageContentBuilder.WriteString(chunk.GetDeltaContent())
		if toolCalls := chunk.GetToolCalls(); toolCalls != nil {
			if message.ToolCalls == nil {
				message.ToolCalls = make(ToolCalls, 0, len(toolCalls))
			}
			for _, toolCall := range toolCalls {
				if !toolCall.Index.IsNull() {
					toolCallIndex := toolCall.Index.Value()
					if len(message.ToolCalls) <= toolCallIndex {
						for i := 0; i < toolCallIndex-len(message.ToolCalls)+1; i++ {
							message.ToolCalls = append(message.ToolCalls, new(ToolCall))
							toolCallArgumentsBuilders = append(toolCallArgumentsBuilders, strings.Builder{})
						}
					}
					messageToolCall := message.ToolCalls[toolCallIndex]
					if messageToolCall.ID == "" {
						message.ToolCalls[toolCallIndex].ID = toolCall.ID
					}
					if messageToolCall.Type == "" {
						message.ToolCalls[toolCallIndex].Type = toolCall.Type
					}
					if functionCall := toolCall.Function; functionCall != nil {
						if messageToolCall.Function == nil {
							messageToolCall.Function = new(FunctionCall)
						}
						if messageToolCall.Function.Name == "" {
							messageToolCall.Function.Name = functionCall.Name
						}
						toolCallArgumentsBuilders[toolCallIndex].WriteString(functionCall.Arguments)
					}
				}
			}
		}
	}
	message.Content = &Content{Text: messageContentBuilder.String()}
	for i := 0; i < len(message.ToolCalls); i++ {
		message.ToolCalls[i].Function.Arguments = toolCallArgumentsBuilders[i].String()
	}
	return message
}

type FunctionCall struct {
	Name      string `json:"name"`
	Arguments string `json:"arguments"`
}

type Tool struct {
	Type     string   `json:"type"`
	Function Function `json:"function"`
}

type Tools []*Tool

type ToolCall struct {
	Index    NullableType[int] `json:"index,omitempty"`
	ID       string            `json:"id"`
	Type     string            `json:"type"`
	Function *FunctionCall     `json:"function"`
}

type ToolCalls []*ToolCall

type Image struct {
	Created int `json:"created"`
	Data    []struct {
		RevisedPrompt string `json:"revised_prompt"`
		Url           string `json:"url"`
	} `json:"data"`
}

type File struct {
	ID        string `json:"id"`
	Object    string `json:"object"`
	Bytes     int    `json:"bytes"`
	CreatedAt int    `json:"created_at"`
	Filename  string `json:"filename"`
	Purpose   string `json:"purpose"`
}

type ToolChoice string

func (t ToolChoice) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type": "function",
		"function": map[string]any{
			"name": string(t),
		},
	})
}

type ResponseFormat string

func (r ResponseFormat) MarshalJSON() ([]byte, error) {
	return json.Marshal(map[string]any{
		"type": string(r),
	})
}

type Stop interface {
	sequences() []string
}

type (
	StopSequence  string
	StopSequences []string
)

func (s StopSequence) sequences() []string  { return []string{string(s)} }
func (s StopSequences) sequences() []string { return s }

type Map[T any] map[string]T

type NullableType[T interface {
	string | int | float64 | bool
}] string

func (nullableVal NullableType[T]) IsNull() bool {
	var ty T
	switch any(ty).(type) {
	case string:
		return false
	default:
		return nullableVal == ""
	}
}

func (nullableVal NullableType[T]) Value() (val T) {
	if !nullableVal.IsNull() {
		if valPtr, isString := any(&val).(*string); isString {
			*valPtr = string(nullableVal)
			return val
		}
		if unmarshalErr := json.Unmarshal([]byte(nullableVal), &val); unmarshalErr != nil {
			panic(fmt.Errorf("unexpected json.Unmarshal error: %w", unmarshalErr))
		}
	}
	return val
}

func (nullableVal NullableType[T]) MarshalJSON() ([]byte, error) {
	var ty T
	switch any(ty).(type) {
	case string:
		return json.Marshal(string(nullableVal))
	case int:
		if nullableVal == "" {
			return json.Marshal(nil)
		}
		intVal, parseIntErr := strconv.Atoi(string(nullableVal))
		if parseIntErr != nil {
			return nil, fmt.Errorf("cannot convert %q to integer number: %w", nullableVal, parseIntErr)
		}
		return json.Marshal(intVal)
	case float64:
		if nullableVal == "" {
			return json.Marshal(nil)
		}
		floVal, parseFloErr := strconv.ParseFloat(string(nullableVal), 64)
		if parseFloErr != nil {
			return nil, fmt.Errorf("cannot convert %q to float number: %w", nullableVal, parseFloErr)
		}
		return json.Marshal(floVal)
	case bool:
		if nullableVal == "" {
			return json.Marshal(nil)
		}
		boolVal, parseBoolErr := strconv.ParseBool(string(nullableVal))
		if parseBoolErr != nil {
			return nil, fmt.Errorf("cannot convert %q to bool: %w", nullableVal, parseBoolErr)
		}
		return json.Marshal(boolVal)
	}
	return []byte(nullableVal), nil
}

func (nullableVal *NullableType[T]) UnmarshalJSON(data []byte) error {
	decoder := newDecoder(data)
	tok, _ := decoder.Token()
	if tok == nil {
		return nil
	}
	var (
		ty         T
		assignable bool
	)
	switch tokVal := tok.(type) {
	case string:
		_, assignable = any(ty).(string)
		if assignable {
			var stringVal string
			if unmarshalErr := json.Unmarshal(data, &stringVal); unmarshalErr != nil {
				return unmarshalErr
			}
			*nullableVal = NullableType[T](stringVal)
			return nil
		}
	case json.Number:
		switch any(ty).(type) {
		case int:
			_, parseIntErr := tokVal.Int64()
			assignable = parseIntErr == nil
		case float64:
			_, parseFloErr := tokVal.Float64()
			assignable = parseFloErr == nil
		default:
			assignable = false
		}
	case bool:
		_, assignable = any(ty).(bool)
	}
	if assignable {
		*nullableVal = NullableType[T](data)
		return nil
	} else {
		return fmt.Errorf("cannot unmarshal %s into Go value of type NullableType[%T]", tokenType(tok), ty)
	}
}

func newDecoder(data []byte) *json.Decoder {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	return decoder
}

func tokenType(tok json.Token) string {
	if tok == nil {
		return "null"
	}
	switch tokVal := tok.(type) {
	case json.Delim:
		switch tokVal {
		case '[':
			return "array"
		case '{':
			return "object"
		}
	case bool:
		return "bool"
	case json.Number:
		return "number"
	case string:
		return "string"
	}
	return "unknown"
}

// Function should be one of these types:
//
//  1. json.Marshaler, whose value must be a valid JSON object
//  2. struct, who can be correctly serialized as a JSON object
//  3. json.RawMessage, who holds a value of a valid JSON object
//  4. map[string]any
//
// For convenience, it is recommended to directly use the DefFunction type.
type Function any

type DefFunction struct {
	Name        string       `json:"name"`
	Description string       `json:"description"`
	Parameters  *DefProperty `json:"parameters"`
}

type DefProperty struct {
	Type        string
	Description string
	Properties  DefProperties
	Required    DefRequired
	Items       DefItems

	// ExtraDefs
	// Other fields that comply with the JSON Schema definition can be
	// added to ExtraDefs and submitted to the model together.
	ExtraDefs Map[any]
}

func (dp DefProperty) MarshalJSON() ([]byte, error) {
	props := make(map[string]any, 5+3)
	props["type"] = dp.Type
	if dp.Description != "" {
		props["description"] = dp.Description
	}
	if dp.Properties != nil {
		props["properties"] = dp.Properties
	}
	if dp.Required != nil {
		props["required"] = dp.Required
	}
	if dp.Items != nil {
		props["items"] = dp.Items
	}
	for k, v := range dp.ExtraDefs {
		props[k] = v
	}
	return json.Marshal(props)
}

type (
	DefProperties map[string]*DefProperty
	DefRequired   []string
	DefItems      []*DefProperty
)
