package azure

import (
	"context"
	"github.com/x5iu/openai"
	"os"
)

type Client interface {
	CreateChatCompletion(ctx context.Context, request *openai.ChatCompletionRequest) (openai.ChatCompletion, error)
}

type Caller interface {
	Endpoint() string
	APIKey() string
	APIVersion() string

	// DeploymentID is an optional setting. If you set a DeploymentID,
	// the BaseClient will consistently use this DeploymentID for model
	// invocations. However, if you do not set a DeploymentID (i.e.,
	// the DeploymentID is empty), then the BaseClient will utilize the
	// Model parameter from the request as its DeploymentID.
	//
	// If you want to use the same BaseClient to call different models
	// (i.e., different deployments), you can leave the DeploymentID
	// blank and specify the DeploymentID using the Model parameter
	// in the request.
	DeploymentID() openai.NullableType[string]
}

func NewClient[C Caller](caller C) Client {
	return NewBaseClient(caller)
}

func DefaultClient() Client {
	return NewClient[BaseCaller](BaseCaller{})
}

// BaseCaller is a simple implementation of Caller, which retrieves the required Endpoint, APIKey and APIVersion from
// environment variables. It uses http.DefaultClient as the HTTP client and does not output logs. The environment
// variables for Endpoint, APIKey AND APIVersion follow the conventions of the OpenAI Python library:
//
//   - Endpoint:   AZURE_OPENAI_ENDPOINT
//   - APIKey:     AZURE_OPENAI_API_KEY
//   - APIVersion: OPENAI_API_VERSION
//
// Additionally, BaseCaller will not provide a DeploymentID, which means the Client will determine which Deployment to
// use based on the Model you provide.
//
// You can embed BaseCaller into your structure and additionally implement the openai.CustomHTTPClient and
// openai.Logger interfaces, which is a convenient solution.
type BaseCaller struct{}

func (BaseCaller) Endpoint() string                          { return os.Getenv("AZURE_OPENAI_ENDPOINT") }
func (BaseCaller) APIKey() string                            { return os.Getenv("AZURE_OPENAI_API_KEY") }
func (BaseCaller) APIVersion() string                        { return os.Getenv("OPENAI_API_VERSION") }
func (BaseCaller) DeploymentID() openai.NullableType[string] { return "" }
