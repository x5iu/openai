//go:build legacy

package openai

import (
	"context"
	"encoding/json"
	"io"
	"strings"
)

//go:generate go run -mod=mod github.com/x5iu/defc generate --features=api/logx,api/error,api/future,api/client --func=trim_trailing_slash=TrimTrailingSlash --func=encode_json=EncodeJSON
type BaseClient[C Caller] interface {
	// ListModels GET {{ trim_trailing_slash $.BaseClient.BaseUrl }}/models
	// Authorization: Bearer {{ $.BaseClient.APIKey }}
	ListModels(ctx context.Context) (*Models, error)

	// CreateChatCompletion POST {{ trim_trailing_slash $.BaseClient.BaseUrl }}/chat/completions
	// Content-Type: application/json
	// Authorization: Bearer {{ $.BaseClient.APIKey }}
	//
	// {{ encode_json $.request }}
	CreateChatCompletion(ctx context.Context, request *ChatCompletionRequest) (ChatCompletion, error)

	// CreateEmbeddings POST {{ trim_trailing_slash $.BaseClient.BaseUrl }}/embeddings
	// Content-Type: application/json
	// Authorization: Bearer {{ $.BaseClient.APIKey }}
	//
	// {{ encode_json $.request }}
	CreateEmbeddings(ctx context.Context, request *CreateEmbeddingsRequest) (*Embeddings, error)

	// CreateImage POST {{ trim_trailing_slash $.BaseClient.BaseUrl }}/images/generations
	// Content-Type: application/json
	// Authorization: Bearer {{ $.BaseClient.APIKey }}
	//
	// {{ encode_json $.request }}
	CreateImage(ctx context.Context, request *CreateImageRequest) (*Images, error)

	// CreateImageEdit POST {{ trim_trailing_slash $.BaseClient.BaseUrl }}/images/edits
	// Content-Type: {{ $.request.ContentType }}
	// Authorization: Bearer {{ $.BaseClient.APIKey }}
	CreateImageEdit(ctx context.Context, request *CreateImageEditRequest) (*Images, error)

	// CreateImageVariation POST {{ trim_trailing_slash $.BaseClient.BaseUrl }}/images/variations
	// Content-Type: {{ $.request.ContentType }}
	// Authorization: Bearer {{ $.BaseClient.APIKey }}
	CreateImageVariation(ctx context.Context, request *CreateImageVariationRequest) (*Images, error)

	// UploadFile POST {{ trim_trailing_slash $.BaseClient.BaseUrl }}/files
	// Content-Type: {{ $.request.ContentType }}
	// Authorization: Bearer {{ $.BaseClient.APIKey }}
	UploadFile(ctx context.Context, request *UploadFileRequest) (*File, error)

	// RetrieveFileContent GET {{ trim_trailing_slash $.BaseClient.BaseUrl }}/files/{{ $.fileID }}/content
	// Authorization: Bearer {{ $.BaseClient.APIKey }}
	RetrieveFileContent(ctx context.Context, fileID string) (io.ReadCloser, string, error)

	Inner() C
	response() *ResponseHandler
}

func TrimTrailingSlash(s string) string {
	for strings.HasSuffix(s, "/") {
		s = strings.TrimSuffix(s, "/")
	}
	return s
}

func EncodeJSON(obj any) (string, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
