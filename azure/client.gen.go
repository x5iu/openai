// Code generated by defc, DO NOT EDIT.

package azure

import (
	"bufio"
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"reflect"
	"sync"
	"text/template"
	"time"

	"github.com/x5iu/openai"
)

const (
	CallerCreateChatCompletion       = "CreateChatCompletion"
	CallerCreateChatCompletionStream = "CreateChatCompletionStream"
)

func NewClient[C Caller](Client C) Client[C] {
	return &implClient[C]{__Client: Client}
}

type implClient[C Caller] struct {
	__Client C
}

var (
	addrTmplCreateChatCompletion         = template.Must(template.New("AddressCreateChatCompletion").Funcs(template.FuncMap{"trimTrailingSlash": openai.TrimTrailingSlash}).Parse("{{ trimTrailingSlash $.Client.BaseUrl }}/openai{{ if $.Client.DeploymentID }}/deployments/{{ $.Client.DeploymentID }}{{ end }}/chat/completions?api-version={{ $.Client.APIVersion }}"))
	headerTmplCreateChatCompletion       = template.Must(template.New("HeaderCreateChatCompletion").Funcs(template.FuncMap{"trimTrailingSlash": openai.TrimTrailingSlash}).Parse("Content-Type: application/json\r\nApi-Key: {{ $.Client.APIKey }}\r\n\r\n{{ $.request.ToJSON }}"))
	addrTmplCreateChatCompletionStream   = template.Must(template.New("AddressCreateChatCompletionStream").Funcs(template.FuncMap{"trimTrailingSlash": openai.TrimTrailingSlash}).Parse("{{ trimTrailingSlash $.Client.BaseUrl }}/openai{{ if $.Client.DeploymentID }}/deployments/{{ $.Client.DeploymentID }}{{ end }}/chat/completions?api-version={{ $.Client.APIVersion }}"))
	headerTmplCreateChatCompletionStream = template.Must(template.New("HeaderCreateChatCompletionStream").Funcs(template.FuncMap{"trimTrailingSlash": openai.TrimTrailingSlash}).Parse("Content-Type: application/json\r\nApi-Key: {{ $.Client.APIKey }}\r\n\r\n{{ $.request.ToJSON }}"))
)

func (__imp *implClient[C]) CreateChatCompletion(ctx context.Context, request *openai.ChatCompletionRequest) (*openai.Completion, error) {
	var innerCreateChatCompletion any = __imp.Inner()

	addrCreateChatCompletion := __ClientGetBuffer()
	defer __ClientPutBuffer(addrCreateChatCompletion)
	defer addrCreateChatCompletion.Reset()

	headerCreateChatCompletion := __ClientGetBuffer()
	defer __ClientPutBuffer(headerCreateChatCompletion)
	defer headerCreateChatCompletion.Reset()

	var (
		v0CreateChatCompletion           = new(openai.Completion)
		errCreateChatCompletion          error
		httpResponseCreateChatCompletion *http.Response
		responseCreateChatCompletion     ClientResponseInterface = __imp.response()
	)

	if errCreateChatCompletion = addrTmplCreateChatCompletion.Execute(addrCreateChatCompletion, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error building 'CreateChatCompletion' url: %w", errCreateChatCompletion)
	}

	if errCreateChatCompletion = headerTmplCreateChatCompletion.Execute(headerCreateChatCompletion, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error building 'CreateChatCompletion' header: %w", errCreateChatCompletion)
	}
	bufReaderCreateChatCompletion := bufio.NewReader(headerCreateChatCompletion)
	mimeHeaderCreateChatCompletion, errCreateChatCompletion := textproto.NewReader(bufReaderCreateChatCompletion).ReadMIMEHeader()
	if errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error reading 'CreateChatCompletion' header: %w", errCreateChatCompletion)
	}

	urlCreateChatCompletion := addrCreateChatCompletion.String()
	requestBodyCreateChatCompletion, errCreateChatCompletion := io.ReadAll(bufReaderCreateChatCompletion)
	if errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error reading 'CreateChatCompletion' request body: %w", errCreateChatCompletion)
	}
	requestCreateChatCompletion, errCreateChatCompletion := http.NewRequestWithContext(ctx, "POST", urlCreateChatCompletion, bytes.NewReader(requestBodyCreateChatCompletion))
	if errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error building 'CreateChatCompletion' request: %w", errCreateChatCompletion)
	}

	for kCreateChatCompletion, vvCreateChatCompletion := range mimeHeaderCreateChatCompletion {
		for _, vCreateChatCompletion := range vvCreateChatCompletion {
			requestCreateChatCompletion.Header.Add(kCreateChatCompletion, vCreateChatCompletion)
		}
	}

	startCreateChatCompletion := time.Now()

	if httpClientCreateChatCompletion, okCreateChatCompletion := innerCreateChatCompletion.(interface{ Client() *http.Client }); okCreateChatCompletion {
		httpResponseCreateChatCompletion, errCreateChatCompletion = httpClientCreateChatCompletion.Client().Do(requestCreateChatCompletion)
	} else {
		httpResponseCreateChatCompletion, errCreateChatCompletion = http.DefaultClient.Do(requestCreateChatCompletion)
	}

	if logCreateChatCompletion, okCreateChatCompletion := innerCreateChatCompletion.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okCreateChatCompletion {
		logCreateChatCompletion.Log(ctx, "CreateChatCompletion", requestCreateChatCompletion, httpResponseCreateChatCompletion, time.Since(startCreateChatCompletion))
	}

	if errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error sending 'CreateChatCompletion' request: %w", errCreateChatCompletion)
	}

	if httpResponseCreateChatCompletion.StatusCode < 200 || httpResponseCreateChatCompletion.StatusCode > 299 {
		return v0CreateChatCompletion, __ClientNewResponseError("CreateChatCompletion", httpResponseCreateChatCompletion)
	}

	if errCreateChatCompletion = responseCreateChatCompletion.FromResponse("CreateChatCompletion", httpResponseCreateChatCompletion); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error converting 'CreateChatCompletion' response: %w", errCreateChatCompletion)
	}

	addrCreateChatCompletion.Reset()
	headerCreateChatCompletion.Reset()

	if errCreateChatCompletion = responseCreateChatCompletion.Err(); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error returned from 'CreateChatCompletion' response: %w", errCreateChatCompletion)
	}

	if errCreateChatCompletion = responseCreateChatCompletion.ScanValues(v0CreateChatCompletion); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error scanning value from 'CreateChatCompletion' response: %w", errCreateChatCompletion)
	}

	return v0CreateChatCompletion, nil
}

func (__imp *implClient[C]) CreateChatCompletionStream(ctx context.Context, request *openai.ChatCompletionStreamRequest) (*openai.Stream, error) {
	var innerCreateChatCompletionStream any = __imp.Inner()

	addrCreateChatCompletionStream := __ClientGetBuffer()
	defer __ClientPutBuffer(addrCreateChatCompletionStream)
	defer addrCreateChatCompletionStream.Reset()

	headerCreateChatCompletionStream := __ClientGetBuffer()
	defer __ClientPutBuffer(headerCreateChatCompletionStream)
	defer headerCreateChatCompletionStream.Reset()

	var (
		v0CreateChatCompletionStream           = new(openai.Stream)
		errCreateChatCompletionStream          error
		httpResponseCreateChatCompletionStream *http.Response
		responseCreateChatCompletionStream     ClientResponseInterface = __imp.response()
	)

	if errCreateChatCompletionStream = addrTmplCreateChatCompletionStream.Execute(addrCreateChatCompletionStream, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error building 'CreateChatCompletionStream' url: %w", errCreateChatCompletionStream)
	}

	if errCreateChatCompletionStream = headerTmplCreateChatCompletionStream.Execute(headerCreateChatCompletionStream, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error building 'CreateChatCompletionStream' header: %w", errCreateChatCompletionStream)
	}
	bufReaderCreateChatCompletionStream := bufio.NewReader(headerCreateChatCompletionStream)
	mimeHeaderCreateChatCompletionStream, errCreateChatCompletionStream := textproto.NewReader(bufReaderCreateChatCompletionStream).ReadMIMEHeader()
	if errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error reading 'CreateChatCompletionStream' header: %w", errCreateChatCompletionStream)
	}

	urlCreateChatCompletionStream := addrCreateChatCompletionStream.String()
	requestBodyCreateChatCompletionStream, errCreateChatCompletionStream := io.ReadAll(bufReaderCreateChatCompletionStream)
	if errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error reading 'CreateChatCompletionStream' request body: %w", errCreateChatCompletionStream)
	}
	requestCreateChatCompletionStream, errCreateChatCompletionStream := http.NewRequestWithContext(ctx, "POST", urlCreateChatCompletionStream, bytes.NewReader(requestBodyCreateChatCompletionStream))
	if errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error building 'CreateChatCompletionStream' request: %w", errCreateChatCompletionStream)
	}

	for kCreateChatCompletionStream, vvCreateChatCompletionStream := range mimeHeaderCreateChatCompletionStream {
		for _, vCreateChatCompletionStream := range vvCreateChatCompletionStream {
			requestCreateChatCompletionStream.Header.Add(kCreateChatCompletionStream, vCreateChatCompletionStream)
		}
	}

	startCreateChatCompletionStream := time.Now()

	if httpClientCreateChatCompletionStream, okCreateChatCompletionStream := innerCreateChatCompletionStream.(interface{ Client() *http.Client }); okCreateChatCompletionStream {
		httpResponseCreateChatCompletionStream, errCreateChatCompletionStream = httpClientCreateChatCompletionStream.Client().Do(requestCreateChatCompletionStream)
	} else {
		httpResponseCreateChatCompletionStream, errCreateChatCompletionStream = http.DefaultClient.Do(requestCreateChatCompletionStream)
	}

	if logCreateChatCompletionStream, okCreateChatCompletionStream := innerCreateChatCompletionStream.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okCreateChatCompletionStream {
		logCreateChatCompletionStream.Log(ctx, "CreateChatCompletionStream", requestCreateChatCompletionStream, httpResponseCreateChatCompletionStream, time.Since(startCreateChatCompletionStream))
	}

	if errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error sending 'CreateChatCompletionStream' request: %w", errCreateChatCompletionStream)
	}

	if httpResponseCreateChatCompletionStream.StatusCode < 200 || httpResponseCreateChatCompletionStream.StatusCode > 299 {
		return v0CreateChatCompletionStream, __ClientNewResponseError("CreateChatCompletionStream", httpResponseCreateChatCompletionStream)
	}

	if errCreateChatCompletionStream = responseCreateChatCompletionStream.FromResponse("CreateChatCompletionStream", httpResponseCreateChatCompletionStream); errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error converting 'CreateChatCompletionStream' response: %w", errCreateChatCompletionStream)
	}

	addrCreateChatCompletionStream.Reset()
	headerCreateChatCompletionStream.Reset()

	if errCreateChatCompletionStream = responseCreateChatCompletionStream.Err(); errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error returned from 'CreateChatCompletionStream' response: %w", errCreateChatCompletionStream)
	}

	if errCreateChatCompletionStream = responseCreateChatCompletionStream.ScanValues(v0CreateChatCompletionStream); errCreateChatCompletionStream != nil {
		return v0CreateChatCompletionStream, fmt.Errorf("error scanning value from 'CreateChatCompletionStream' response: %w", errCreateChatCompletionStream)
	}

	return v0CreateChatCompletionStream, nil
}

func (__imp *implClient[C]) Inner() C {
	return __imp.__Client
}

func (*implClient[C]) response() *openai.ResponseHandler {
	return new(openai.ResponseHandler)
}

var __ClientBufferPool = sync.Pool{
	New: func() any {
		return new(bytes.Buffer)
	},
}

func __ClientGetBuffer() *bytes.Buffer {
	return __ClientBufferPool.Get().(*bytes.Buffer)
}

func __ClientPutBuffer(buffer *bytes.Buffer) {
	__ClientBufferPool.Put(buffer)
}

type ClientResponseInterface interface {
	Err() error
	ScanValues(...any) error
	FromResponse(string, *http.Response) error
	Break() bool
}

func __ClientNewType(typ reflect.Type) reflect.Value {
	switch typ.Kind() {
	case reflect.Slice:
		return reflect.MakeSlice(typ, 0, 0)
	case reflect.Map:
		return reflect.MakeMap(typ)
	case reflect.Chan:
		return reflect.MakeChan(typ, 0)
	case reflect.Func:
		return reflect.MakeFunc(typ, func(_ []reflect.Value) (results []reflect.Value) {
			results = make([]reflect.Value, typ.NumOut())
			for i := 0; i < typ.NumOut(); i++ {
				results[i] = __ClientNewType(typ.Out(i))
			}
			return results
		})
	case reflect.Pointer:
		return reflect.New(typ.Elem())
	default:
		return reflect.Zero(typ)
	}
}

func __ClientNew[T any]() (v T) {
	val := reflect.ValueOf(&v).Elem()
	switch val.Kind() {
	case reflect.Slice, reflect.Map, reflect.Chan, reflect.Func, reflect.Pointer:
		val.Set(__ClientNewType(val.Type()))
	}
	return v
}

// ClientResponseErrorInterface represents future Response error interface which would
// be used in next major version of defc, who may cause breaking changes.
//
// generated with --features=api/future
type ClientResponseErrorInterface interface {
	error
	Response() *http.Response
}

func __ClientNewResponseError(caller string, response *http.Response) ClientResponseErrorInterface {
	return &__ClientImplResponseError{
		caller:   caller,
		response: response,
	}
}

type __ClientImplResponseError struct {
	caller   string
	response *http.Response
}

func (e *__ClientImplResponseError) Error() string {
	return fmt.Sprintf("response status code %d for '%s'", e.response.StatusCode, e.caller)
}

func (e *__ClientImplResponseError) Response() *http.Response {
	return e.response
}
