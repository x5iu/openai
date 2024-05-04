//go:build !legacy

package azure

import (
	"context"
	"github.com/x5iu/openai"
)

//go:generate go run -mod=mod github.com/x5iu/defc generate --features=api/logx,api/error,api/future,api/client --func=trim_trailing_slash=openai.TrimTrailingSlash
type BaseClient[C Caller] interface {
	/*
		CreateChatCompletion POST {{ trim_trailing_slash $.BaseClient.Endpoint }}/openai\
			/deployments/{{ if $.BaseClient.DeploymentID }}{{ $.BaseClient.DeploymentID }}{{ else }}{{ $.request.Model }}{{ end }}\
			/chat/completions?\
			api-version={{ $.BaseClient.APIVersion }}
	*/
	// Content-Type: application/json
	// Api-Key: {{ $.BaseClient.APIKey }}
	CreateChatCompletion(ctx context.Context, request *openai.ChatCompletionRequest) (openai.ChatCompletion, error)

	Inner() C
	response() *openai.ResponseHandler
}
