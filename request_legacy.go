//go:build legacy

package openai

import (
	"io"

	defc "github.com/x5iu/defc/runtime"
)

type UploadFile interface {
	io.Reader
	Name() string
}

type ChatCompletionRequest struct {
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
	StreamOptions    StreamOptions         `json:"stream_options,omitempty"`
	Temperature      NullableType[float64] `json:"temperature,omitempty"`
	TopP             NullableType[float64] `json:"top_p,omitempty"`
	Tools            Tools                 `json:"tools,omitempty"`
	ToolChoice       ToolChoice            `json:"tool_choice,omitempty"`
	User             string                `json:"user,omitempty"`
}

type CreateEmbeddingsRequest struct {
	Input          Input  `json:"input"`
	Model          string `json:"model"`
	EncodingFormat string `json:"encoding_format,omitempty"`
	Dimensions     int    `json:"dimensions,omitempty"`
	User           string `json:"user,omitempty"`
}

type CreateImageRequest struct {
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
	defc.MultipartBody[CreateImageEditRequest]

	Image          UploadFile `form:"image"`
	Prompt         string     `form:"prompt"`
	Mask           UploadFile `form:"mask,omitempty"`
	Model          string     `form:"model,omitempty"`
	N              int        `form:"n,omitempty"`
	Size           string     `form:"size,omitempty"`
	ResponseFormat string     `form:"response_format,omitempty"`
	User           string     `form:"user,omitempty"`
}

type CreateImageVariationRequest struct {
	defc.MultipartBody[CreateImageVariationRequest]

	Image          UploadFile `form:"image"`
	Model          string     `form:"model,omitempty"`
	N              int        `form:"n,omitempty"`
	ResponseFormat string     `form:"response_format,omitempty"`
	Size           string     `form:"size,omitempty"`
	User           string     `form:"user,omitempty"`
}

type UploadFileRequest struct {
	defc.MultipartBody[UploadFileRequest]

	File    UploadFile `form:"file"`
	Purpose string     `form:"purpose"`
}
