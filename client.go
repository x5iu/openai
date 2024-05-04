package openai

import (
	"context"
	"io"
	"net/http"
	"os"
	"time"
)

type Client interface {
	ListModels(ctx context.Context) (*Models, error)
	CreateChatCompletion(ctx context.Context, request *ChatCompletionRequest) (ChatCompletion, error)
	CreateEmbeddings(ctx context.Context, request *CreateEmbeddingsRequest) (*Embeddings, error)
	CreateImage(ctx context.Context, request *CreateImageRequest) (*Images, error)
	CreateImageEdit(ctx context.Context, request *CreateImageEditRequest) (*Images, error)
	CreateImageVariation(ctx context.Context, request *CreateImageVariationRequest) (*Images, error)
	UploadFile(ctx context.Context, request *UploadFileRequest) (*File, error)
	RetrieveFileContent(ctx context.Context, fileID string) (io.ReadCloser, string, error)
}

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

func NewClient[C Caller](caller C) Client {
	return NewBaseClient(caller)
}

func DefaultClient() Client {
	return NewClient[BaseCaller](BaseCaller{})
}

// BaseCaller is a simple implementation of Caller, which retrieves the required BaseUrl and APIKey
// from environment variables. It uses http.DefaultClient as the HTTP client and does not output logs.
// The environment variables for BaseUrl and APIKey follow the conventions of the OpenAI Python library:
//
//   - BaseUrl: OPENAI_BASE_URL
//   - APIKey:  OPENAI_API_KEY
//
// You can embed BaseCaller into your structure and additionally implement the CustomHTTPClient and
// Logger interfaces, which is a convenient solution.
type BaseCaller struct{}

func (BaseCaller) BaseUrl() string { return os.Getenv("OPENAI_BASE_URL") }
func (BaseCaller) APIKey() string  { return os.Getenv("OPENAI_API_KEY") }
