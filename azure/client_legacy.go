//go:build legacy

package azure

import (
	"context"
	"github.com/x5iu/openai"
)

type Caller interface {
	openai.Caller
	APIVersion() string
	DeploymentID() string
}

//go:generate go run -mod=mod github.com/x5iu/defc generate --features=api/logx,api/error,api/future,api/client --func=trimTrailingSlash=openai.TrimTrailingSlash --func=encodejson=openai.EncodeJSON --import "github.com/x5iu/openai"
type Client[C Caller] interface {
	/*
		CreateChatCompletion POST {{ trimTrailingSlash $.Client.BaseUrl }}/openai\
			{{ if $.Client.DeploymentID }}/deployments/{{ $.Client.DeploymentID }}{{ end }}\
			/chat/completions?\
			api-version={{ $.Client.APIVersion }}
	*/
	// Content-Type: application/json
	// Api-Key: {{ $.Client.APIKey }}
	//
	// {{ encodejson $.request }}
	CreateChatCompletion(ctx context.Context, request *openai.ChatCompletionRequest) (openai.ChatCompletion, error)

	Inner() C
	response() *openai.ResponseHandler
}
