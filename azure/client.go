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

//go:generate go run -mod=mod github.com/x5iu/defc generate --features=api/nort,api/logx,api/error,api/future,api/client --func=trimTrailingSlash=openai.TrimTrailingSlash
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
	// {{ $.request.ToJSON }}
	CreateChatCompletion(ctx context.Context, request *openai.ChatCompletionRequest) (*openai.Completion, error)

	/*
		CreateChatCompletionStream POST {{ trimTrailingSlash $.Client.BaseUrl }}/openai\
			{{ if $.Client.DeploymentID }}/deployments/{{ $.Client.DeploymentID }}{{ end }}\
			/chat/completions?\
			api-version={{ $.Client.APIVersion }}
	*/
	// Content-Type: application/json
	// Api-Key: {{ $.Client.APIKey }}
	//
	// {{ $.request.ToJSON }}
	CreateChatCompletionStream(ctx context.Context, request *openai.ChatCompletionStreamRequest) (*openai.Stream, error)

	Inner() C
	response() *openai.ResponseHandler
}
