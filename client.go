package openai

import (
	"context"
	"io"
	"net/http"
	"strings"
	"time"
)

type Caller interface {
	BaseUrl() string
	APIKey() string
}

type Logger interface {
	Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
}

type CustomHTTPClient interface {
	Client() *http.Client
}

//go:generate go run -mod=mod "github.com/x5iu/defc" --mode=api --output=client.gen.go --features=api/nort,api/logx,api/error,api/future,api/client --func trimTrailingSlash=TrimTrailingSlash
type Client[C Caller] interface {
	// ListModels GET {{ trimTrailingSlash $.Client.BaseUrl }}/models
	// Authorization: Bearer {{ $.Client.APIKey }}
	ListModels(ctx context.Context) (*Models, error)

	// CreateChatCompletion POST {{ trimTrailingSlash $.Client.BaseUrl }}/chat/completions
	// Content-Type: application/json
	// Authorization: Bearer {{ $.Client.APIKey }}
	//
	// {{ $.request.ToJSON }}
	CreateChatCompletion(ctx context.Context, request *ChatCompletionRequest) (*Completion, error)

	// CreateChatCompletionStream POST {{ trimTrailingSlash $.Client.BaseUrl }}/chat/completions
	// Content-Type: application/json
	// Authorization: Bearer {{ $.Client.APIKey }}
	//
	// {{ $.request.ToJSON }}
	CreateChatCompletionStream(ctx context.Context, request *ChatCompletionStreamRequest) (*Stream, error)

	// CreateImage POST {{ trimTrailingSlash $.Client.BaseUrl }}/images/generations
	// Content-Type: application/json
	// Authorization: Bearer {{ $.Client.APIKey }}
	//
	// {{ $.request.ToJSON }}
	CreateImage(ctx context.Context, request *CreateImageRequest) (*Image, error)

	// UploadFile POST {{ trimTrailingSlash $.Client.BaseUrl }}/files
	// Content-Type: {{ $.request.ContentType }}
	// Authorization: Bearer {{ $.Client.APIKey }}
	UploadFile(ctx context.Context, request *UploadFileRequest) (*File, error)

	// RetrieveFileContent GET {{ trimTrailingSlash $.Client.BaseUrl }}/files/{{ $.fileID }}/content
	// Authorization: Bearer {{ $.Client.APIKey }}
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
