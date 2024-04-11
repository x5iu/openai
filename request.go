package openai

import (
	"bytes"
	"crypto/rand"
	"encoding/hex"
	"fmt"
	defc "github.com/x5iu/defc/runtime"
	"io"
	"net/textproto"
	"strings"
	"sync"
)

type ChatCompletionRequest struct {
	defc.JSONBody[ChatCompletionRequest]
	Messages         Messages              `json:"messages"`
	Model            string                `json:"model"`
	FrequencyPenalty NullableType[float64] `json:"frequency_penalty,omitempty"`
	LogitBias        Map[int]              `json:"logit_bias,omitempty"`
	Logprobs         bool                  `json:"logprobs,omitempty"`
	TopLogprobs      NullableType[int]     `json:"top_logprobs,omitempty"`
	MaxTokens        int                   `json:"max_tokens,omitempty"`
	N                int                   `json:"n,omitempty"`
	PresencePenalty  NullableType[float64] `json:"presence_penalty,omitempty"`
	ResponseFormat   ResponseFormat        `json:"response_format,omitempty"`
	Seed             int                   `json:"seed,omitempty"`
	Stop             Stop                  `json:"stop,omitempty"`
	Stream           bool                  `json:"stream,omitempty"`
	Temperature      NullableType[float64] `json:"temperature,omitempty"`
	TopP             NullableType[float64] `json:"top_p,omitempty"`
	Tools            Tools                 `json:"tools,omitempty"`
	ToolChoice       ToolChoice            `json:"tool_choice,omitempty"`
	User             string                `json:"user,omitempty"`
}

type CreateImageRequest struct {
	defc.JSONBody[CreateImageRequest]
	Prompt         string `json:"prompt"`
	Model          string `json:"model,omitempty"`
	N              int    `json:"n,omitempty"`
	Quality        string `json:"quality,omitempty"`
	ResponseFormat string `json:"response_format,omitempty"`
	Size           string `json:"size,omitempty"`
	Style          string `json:"style,omitempty"`
	User           string `json:"user,omitempty"`
}

type CreateImageEditRequest struct {
	Image          UploadFile
	Prompt         string
	Mask           UploadFile
	Model          string
	N              int
	Size           string
	ResponseFormat string
	User           string

	formReader formReader
}

func (r *CreateImageEditRequest) ContentType() string {
	return r.formReader.ContentType()
}

func (r *CreateImageEditRequest) getFormData() form {
	formData := form{
		formFieldImage:  r.Image,
		formFieldPrompt: r.Prompt,
	}
	if r.Mask != nil {
		formData[formFieldMask] = r.Mask
	}
	if r.Model != "" {
		formData[formFieldModel] = r.Model
	}
	if r.N != 0 {
		formData[formFieldN] = r.N
	}
	if r.Size != "" {
		formData[formFieldSize] = r.Size
	}
	if r.ResponseFormat != "" {
		formData[formFieldResponseFormat] = r.ResponseFormat
	}
	if r.User != "" {
		formData[formFieldUser] = r.User
	}
	return formData
}

func (r *CreateImageEditRequest) Read(p []byte) (n int, err error) {
	return r.formReader.ReadForm(r.getFormData, p)
}

type CreateImageVariationRequest struct {
	Image          UploadFile
	Model          string
	N              int
	ResponseFormat string
	Size           string
	User           string

	formReader formReader
}

func (r *CreateImageVariationRequest) ContentType() string {
	return r.formReader.ContentType()
}

func (r *CreateImageVariationRequest) getFormData() form {
	formData := form{
		formFieldImage: r.Image,
	}
	if r.Model != "" {
		formData[formFieldModel] = r.Model
	}
	if r.N != 0 {
		formData[formFieldN] = r.N
	}
	if r.ResponseFormat != "" {
		formData[formFieldResponseFormat] = r.ResponseFormat
	}
	if r.Size != "" {
		formData[formFieldSize] = r.Size
	}
	if r.User != "" {
		formData[formFieldUser] = r.User
	}
	return formData
}

func (r *CreateImageVariationRequest) Read(p []byte) (n int, err error) {
	return r.formReader.ReadForm(r.getFormData, p)
}

type UploadFileRequest struct {
	File    UploadFile
	Purpose string

	formReader formReader
}

func (r *UploadFileRequest) ContentType() string {
	return r.formReader.ContentType()
}

func (r *UploadFileRequest) getFormData() form {
	return form{
		formFieldFile:    r.File,
		formFieldPurpose: r.Purpose,
	}
}

func (r *UploadFileRequest) Read(p []byte) (n int, err error) {
	return r.formReader.ReadForm(r.getFormData, p)
}

type UploadFile interface {
	io.Reader
	Name() string
}

const (
	formFieldFile           = "file"
	formFieldPurpose        = "purpose"
	formFieldImage          = "image"
	formFieldPrompt         = "prompt"
	formFieldMask           = "mask"
	formFieldModel          = "model"
	formFieldN              = "n"
	formFieldSize           = "size"
	formFieldResponseFormat = "response_format"
	formFieldUser           = "User"
)

type form map[string]any

type formReader struct {
	once     sync.Once
	reader   io.Reader
	boundary string
}

func randomBoundary() string {
	var buf [64]byte
	io.ReadFull(rand.Reader, buf[:])
	return hex.EncodeToString(buf[:])
}

func (r *formReader) ContentType() string {
	if r.boundary == "" {
		r.boundary = randomBoundary()
	}
	b := r.boundary
	if strings.ContainsAny(b, `()<>@,;:\"/[]?= `) {
		b = `"` + b + `"`
	}
	return "multipart/form-data; boundary=" + b
}

var quoteEscaper = strings.NewReplacer("\\", "\\\\", `"`, "\\\"")

func escapeQuotes(s string) string {
	return quoteEscaper.Replace(s)
}

func (r *formReader) startRead(formData form) {
	if r.boundary == "" {
		r.boundary = randomBoundary()
	}
	readers := make([]io.Reader, 0, len(formData))
	i := 0
	for fieldname, v := range formData {
		if i == 0 {
			readers = append(readers, strings.NewReader("--"+r.boundary+"\r\n"))
		} else {
			readers = append(readers, strings.NewReader("\r\n--"+r.boundary+"\r\n"))
		}
		var (
			buf bytes.Buffer
			h   = make(textproto.MIMEHeader)
		)
		if file, ok := v.(UploadFile); ok {
			h.Set("Content-Disposition",
				fmt.Sprintf(`form-data; name="%s"; filename="%s"`, escapeQuotes(fieldname), escapeQuotes(file.Name())))
			h.Set("Content-Type", "application/octet-stream")
			for hk, hvv := range h {
				for _, hv := range hvv {
					buf.WriteString(hk + ": " + hv + "\r\n")
				}
			}
			buf.WriteString("\r\n")
			readers = append(readers, &buf)
			readers = append(readers, file)
		} else {
			var fieldvalue string
			if fieldvalue, ok = v.(string); !ok {
				fieldvalue = fmt.Sprintf("%v", v)
			}
			h.Set("Content-Disposition", fmt.Sprintf(`form-data; name="%s"`, escapeQuotes(fieldname)))
			for hk, hvv := range h {
				for _, hv := range hvv {
					buf.WriteString(hk + ": " + hv + "\r\n")
				}
			}
			buf.WriteString("\r\n")
			buf.WriteString(fieldvalue)
			readers = append(readers, &buf)
		}
		i++
	}
	readers = append(readers, strings.NewReader("\r\n--"+r.boundary+"--\r\n"))
	r.reader = io.MultiReader(readers...)
}

func (r *formReader) ReadForm(getFormData func() form, p []byte) (n int, err error) {
	r.once.Do(func() {
		r.startRead(getFormData())
	})
	return r.reader.Read(p)
}
