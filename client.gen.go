// Code generated by defc, DO NOT EDIT.

package openai

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/textproto"
	"text/template"
	"time"

	__rt "github.com/x5iu/defc/runtime"
)

const (
	CallerListModels           = "ListModels"
	CallerCreateChatCompletion = "CreateChatCompletion"
	CallerCreateEmbeddings     = "CreateEmbeddings"
	CallerCreateImage          = "CreateImage"
	CallerCreateImageEdit      = "CreateImageEdit"
	CallerCreateImageVariation = "CreateImageVariation"
	CallerUploadFile           = "UploadFile"
	CallerRetrieveFileContent  = "RetrieveFileContent"
)

func NewClient[C Caller](Client C) Client[C] {
	return &implClient[C]{__Client: Client}
}

type implClient[C Caller] struct {
	__Client C
}

var (
	addrTmplListModels             = template.Must(template.New("AddressListModels").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("{{ trimTrailingSlash $.Client.BaseUrl }}/models"))
	headerTmplListModels           = template.Must(template.New("HeaderListModels").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("Authorization: Bearer {{ $.Client.APIKey }}\r\n\r\n"))
	addrTmplCreateChatCompletion   = template.Must(template.New("AddressCreateChatCompletion").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("{{ trimTrailingSlash $.Client.BaseUrl }}/chat/completions"))
	headerTmplCreateChatCompletion = template.Must(template.New("HeaderCreateChatCompletion").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("Content-Type: application/json\r\nAuthorization: Bearer {{ $.Client.APIKey }}\r\n\r\n"))
	addrTmplCreateEmbeddings       = template.Must(template.New("AddressCreateEmbeddings").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("{{ trimTrailingSlash $.Client.BaseUrl }}/embeddings"))
	headerTmplCreateEmbeddings     = template.Must(template.New("HeaderCreateEmbeddings").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("Content-Type: application/json\r\nAuthorization: Bearer {{ $.Client.APIKey }}\r\n\r\n"))
	addrTmplCreateImage            = template.Must(template.New("AddressCreateImage").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("{{ trimTrailingSlash $.Client.BaseUrl }}/images/generations"))
	headerTmplCreateImage          = template.Must(template.New("HeaderCreateImage").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("Content-Type: application/json\r\nAuthorization: Bearer {{ $.Client.APIKey }}\r\n\r\n"))
	addrTmplCreateImageEdit        = template.Must(template.New("AddressCreateImageEdit").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("{{ trimTrailingSlash $.Client.BaseUrl }}/images/edits"))
	headerTmplCreateImageEdit      = template.Must(template.New("HeaderCreateImageEdit").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("Content-Type: {{ $.request.ContentType }}\r\nAuthorization: Bearer {{ $.Client.APIKey }}\r\n\r\n"))
	addrTmplCreateImageVariation   = template.Must(template.New("AddressCreateImageVariation").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("{{ trimTrailingSlash $.Client.BaseUrl }}/images/variations"))
	headerTmplCreateImageVariation = template.Must(template.New("HeaderCreateImageVariation").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("Content-Type: {{ $.request.ContentType }}\r\nAuthorization: Bearer {{ $.Client.APIKey }}\r\n\r\n"))
	addrTmplUploadFile             = template.Must(template.New("AddressUploadFile").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("{{ trimTrailingSlash $.Client.BaseUrl }}/files"))
	headerTmplUploadFile           = template.Must(template.New("HeaderUploadFile").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("Content-Type: {{ $.request.ContentType }}\r\nAuthorization: Bearer {{ $.Client.APIKey }}\r\n\r\n"))
	addrTmplRetrieveFileContent    = template.Must(template.New("AddressRetrieveFileContent").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("{{ trimTrailingSlash $.Client.BaseUrl }}/files/{{ $.fileID }}/content"))
	headerTmplRetrieveFileContent  = template.Must(template.New("HeaderRetrieveFileContent").Funcs(template.FuncMap{"trimTrailingSlash": TrimTrailingSlash}).Parse("Authorization: Bearer {{ $.Client.APIKey }}\r\n\r\n"))
)

func (__imp *implClient[C]) ListModels(ctx context.Context) (*Models, error) {
	var innerListModels any = __imp.Inner()

	addrListModels := __rt.GetBuffer()
	defer __rt.PutBuffer(addrListModels)
	defer addrListModels.Reset()

	headerListModels := __rt.GetBuffer()
	defer __rt.PutBuffer(headerListModels)
	defer headerListModels.Reset()

	var (
		v0ListModels           = new(Models)
		errListModels          error
		httpResponseListModels *http.Response
		responseListModels     __rt.FutureResponse = __imp.response()
	)

	if errListModels = addrTmplListModels.Execute(addrListModels, map[string]any{
		"Client": __imp.Inner(),
		"ctx":    ctx,
	}); errListModels != nil {
		return v0ListModels, fmt.Errorf("error building 'ListModels' url: %w", errListModels)
	}

	if errListModels = headerTmplListModels.Execute(headerListModels, map[string]any{
		"Client": __imp.Inner(),
		"ctx":    ctx,
	}); errListModels != nil {
		return v0ListModels, fmt.Errorf("error building 'ListModels' header: %w", errListModels)
	}
	bufReaderListModels := bufio.NewReader(headerListModels)
	mimeHeaderListModels, errListModels := textproto.NewReader(bufReaderListModels).ReadMIMEHeader()
	if errListModels != nil {
		return v0ListModels, fmt.Errorf("error reading 'ListModels' header: %w", errListModels)
	}

	urlListModels := addrListModels.String()
	requestListModels, errListModels := http.NewRequestWithContext(ctx, "GET", urlListModels, http.NoBody)
	if errListModels != nil {
		return v0ListModels, fmt.Errorf("error building 'ListModels' request: %w", errListModels)
	}

	for kListModels, vvListModels := range mimeHeaderListModels {
		for _, vListModels := range vvListModels {
			requestListModels.Header.Add(kListModels, vListModels)
		}
	}

	startListModels := time.Now()

	if httpClientListModels, okListModels := innerListModels.(interface{ Client() *http.Client }); okListModels {
		httpResponseListModels, errListModels = httpClientListModels.Client().Do(requestListModels)
	} else {
		httpResponseListModels, errListModels = http.DefaultClient.Do(requestListModels)
	}

	if logListModels, okListModels := innerListModels.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okListModels {
		logListModels.Log(ctx, "ListModels", requestListModels, httpResponseListModels, time.Since(startListModels))
	}

	if errListModels != nil {
		return v0ListModels, fmt.Errorf("error sending 'ListModels' request: %w", errListModels)
	}

	if httpResponseListModels.StatusCode < 200 || httpResponseListModels.StatusCode > 299 {
		return v0ListModels, __rt.NewFutureResponseError("ListModels", httpResponseListModels)
	}

	if errListModels = responseListModels.FromResponse("ListModels", httpResponseListModels); errListModels != nil {
		return v0ListModels, fmt.Errorf("error converting 'ListModels' response: %w", errListModels)
	}

	addrListModels.Reset()
	headerListModels.Reset()

	if errListModels = responseListModels.Err(); errListModels != nil {
		return v0ListModels, fmt.Errorf("error returned from 'ListModels' response: %w", errListModels)
	}

	if errListModels = responseListModels.ScanValues(v0ListModels); errListModels != nil {
		return v0ListModels, fmt.Errorf("error scanning value from 'ListModels' response: %w", errListModels)
	}

	return v0ListModels, nil
}

func (__imp *implClient[C]) CreateChatCompletion(ctx context.Context, request *ChatCompletionRequest) (ChatCompletion, error) {
	var innerCreateChatCompletion any = __imp.Inner()

	addrCreateChatCompletion := __rt.GetBuffer()
	defer __rt.PutBuffer(addrCreateChatCompletion)
	defer addrCreateChatCompletion.Reset()

	headerCreateChatCompletion := __rt.GetBuffer()
	defer __rt.PutBuffer(headerCreateChatCompletion)
	defer headerCreateChatCompletion.Reset()

	var (
		v0CreateChatCompletion           = __rt.New[ChatCompletion]()
		errCreateChatCompletion          error
		httpResponseCreateChatCompletion *http.Response
		responseCreateChatCompletion     __rt.FutureResponse = __imp.response()
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
	requestCreateChatCompletion, errCreateChatCompletion := http.NewRequestWithContext(ctx, "POST", urlCreateChatCompletion, request)
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
		return v0CreateChatCompletion, __rt.NewFutureResponseError("CreateChatCompletion", httpResponseCreateChatCompletion)
	}

	if errCreateChatCompletion = responseCreateChatCompletion.FromResponse("CreateChatCompletion", httpResponseCreateChatCompletion); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error converting 'CreateChatCompletion' response: %w", errCreateChatCompletion)
	}

	addrCreateChatCompletion.Reset()
	headerCreateChatCompletion.Reset()

	if errCreateChatCompletion = responseCreateChatCompletion.Err(); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error returned from 'CreateChatCompletion' response: %w", errCreateChatCompletion)
	}

	if errCreateChatCompletion = responseCreateChatCompletion.ScanValues(&v0CreateChatCompletion); errCreateChatCompletion != nil {
		return v0CreateChatCompletion, fmt.Errorf("error scanning value from 'CreateChatCompletion' response: %w", errCreateChatCompletion)
	}

	return v0CreateChatCompletion, nil
}

func (__imp *implClient[C]) CreateEmbeddings(ctx context.Context, request *CreateEmbeddingsRequest) (*Embeddings, error) {
	var innerCreateEmbeddings any = __imp.Inner()

	addrCreateEmbeddings := __rt.GetBuffer()
	defer __rt.PutBuffer(addrCreateEmbeddings)
	defer addrCreateEmbeddings.Reset()

	headerCreateEmbeddings := __rt.GetBuffer()
	defer __rt.PutBuffer(headerCreateEmbeddings)
	defer headerCreateEmbeddings.Reset()

	var (
		v0CreateEmbeddings           = new(Embeddings)
		errCreateEmbeddings          error
		httpResponseCreateEmbeddings *http.Response
		responseCreateEmbeddings     __rt.FutureResponse = __imp.response()
	)

	if errCreateEmbeddings = addrTmplCreateEmbeddings.Execute(addrCreateEmbeddings, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateEmbeddings != nil {
		return v0CreateEmbeddings, fmt.Errorf("error building 'CreateEmbeddings' url: %w", errCreateEmbeddings)
	}

	if errCreateEmbeddings = headerTmplCreateEmbeddings.Execute(headerCreateEmbeddings, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateEmbeddings != nil {
		return v0CreateEmbeddings, fmt.Errorf("error building 'CreateEmbeddings' header: %w", errCreateEmbeddings)
	}
	bufReaderCreateEmbeddings := bufio.NewReader(headerCreateEmbeddings)
	mimeHeaderCreateEmbeddings, errCreateEmbeddings := textproto.NewReader(bufReaderCreateEmbeddings).ReadMIMEHeader()
	if errCreateEmbeddings != nil {
		return v0CreateEmbeddings, fmt.Errorf("error reading 'CreateEmbeddings' header: %w", errCreateEmbeddings)
	}

	urlCreateEmbeddings := addrCreateEmbeddings.String()
	requestCreateEmbeddings, errCreateEmbeddings := http.NewRequestWithContext(ctx, "POST", urlCreateEmbeddings, request)
	if errCreateEmbeddings != nil {
		return v0CreateEmbeddings, fmt.Errorf("error building 'CreateEmbeddings' request: %w", errCreateEmbeddings)
	}

	for kCreateEmbeddings, vvCreateEmbeddings := range mimeHeaderCreateEmbeddings {
		for _, vCreateEmbeddings := range vvCreateEmbeddings {
			requestCreateEmbeddings.Header.Add(kCreateEmbeddings, vCreateEmbeddings)
		}
	}

	startCreateEmbeddings := time.Now()

	if httpClientCreateEmbeddings, okCreateEmbeddings := innerCreateEmbeddings.(interface{ Client() *http.Client }); okCreateEmbeddings {
		httpResponseCreateEmbeddings, errCreateEmbeddings = httpClientCreateEmbeddings.Client().Do(requestCreateEmbeddings)
	} else {
		httpResponseCreateEmbeddings, errCreateEmbeddings = http.DefaultClient.Do(requestCreateEmbeddings)
	}

	if logCreateEmbeddings, okCreateEmbeddings := innerCreateEmbeddings.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okCreateEmbeddings {
		logCreateEmbeddings.Log(ctx, "CreateEmbeddings", requestCreateEmbeddings, httpResponseCreateEmbeddings, time.Since(startCreateEmbeddings))
	}

	if errCreateEmbeddings != nil {
		return v0CreateEmbeddings, fmt.Errorf("error sending 'CreateEmbeddings' request: %w", errCreateEmbeddings)
	}

	if httpResponseCreateEmbeddings.StatusCode < 200 || httpResponseCreateEmbeddings.StatusCode > 299 {
		return v0CreateEmbeddings, __rt.NewFutureResponseError("CreateEmbeddings", httpResponseCreateEmbeddings)
	}

	if errCreateEmbeddings = responseCreateEmbeddings.FromResponse("CreateEmbeddings", httpResponseCreateEmbeddings); errCreateEmbeddings != nil {
		return v0CreateEmbeddings, fmt.Errorf("error converting 'CreateEmbeddings' response: %w", errCreateEmbeddings)
	}

	addrCreateEmbeddings.Reset()
	headerCreateEmbeddings.Reset()

	if errCreateEmbeddings = responseCreateEmbeddings.Err(); errCreateEmbeddings != nil {
		return v0CreateEmbeddings, fmt.Errorf("error returned from 'CreateEmbeddings' response: %w", errCreateEmbeddings)
	}

	if errCreateEmbeddings = responseCreateEmbeddings.ScanValues(v0CreateEmbeddings); errCreateEmbeddings != nil {
		return v0CreateEmbeddings, fmt.Errorf("error scanning value from 'CreateEmbeddings' response: %w", errCreateEmbeddings)
	}

	return v0CreateEmbeddings, nil
}

func (__imp *implClient[C]) CreateImage(ctx context.Context, request *CreateImageRequest) (*Images, error) {
	var innerCreateImage any = __imp.Inner()

	addrCreateImage := __rt.GetBuffer()
	defer __rt.PutBuffer(addrCreateImage)
	defer addrCreateImage.Reset()

	headerCreateImage := __rt.GetBuffer()
	defer __rt.PutBuffer(headerCreateImage)
	defer headerCreateImage.Reset()

	var (
		v0CreateImage           = new(Images)
		errCreateImage          error
		httpResponseCreateImage *http.Response
		responseCreateImage     __rt.FutureResponse = __imp.response()
	)

	if errCreateImage = addrTmplCreateImage.Execute(addrCreateImage, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateImage != nil {
		return v0CreateImage, fmt.Errorf("error building 'CreateImage' url: %w", errCreateImage)
	}

	if errCreateImage = headerTmplCreateImage.Execute(headerCreateImage, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateImage != nil {
		return v0CreateImage, fmt.Errorf("error building 'CreateImage' header: %w", errCreateImage)
	}
	bufReaderCreateImage := bufio.NewReader(headerCreateImage)
	mimeHeaderCreateImage, errCreateImage := textproto.NewReader(bufReaderCreateImage).ReadMIMEHeader()
	if errCreateImage != nil {
		return v0CreateImage, fmt.Errorf("error reading 'CreateImage' header: %w", errCreateImage)
	}

	urlCreateImage := addrCreateImage.String()
	requestCreateImage, errCreateImage := http.NewRequestWithContext(ctx, "POST", urlCreateImage, request)
	if errCreateImage != nil {
		return v0CreateImage, fmt.Errorf("error building 'CreateImage' request: %w", errCreateImage)
	}

	for kCreateImage, vvCreateImage := range mimeHeaderCreateImage {
		for _, vCreateImage := range vvCreateImage {
			requestCreateImage.Header.Add(kCreateImage, vCreateImage)
		}
	}

	startCreateImage := time.Now()

	if httpClientCreateImage, okCreateImage := innerCreateImage.(interface{ Client() *http.Client }); okCreateImage {
		httpResponseCreateImage, errCreateImage = httpClientCreateImage.Client().Do(requestCreateImage)
	} else {
		httpResponseCreateImage, errCreateImage = http.DefaultClient.Do(requestCreateImage)
	}

	if logCreateImage, okCreateImage := innerCreateImage.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okCreateImage {
		logCreateImage.Log(ctx, "CreateImage", requestCreateImage, httpResponseCreateImage, time.Since(startCreateImage))
	}

	if errCreateImage != nil {
		return v0CreateImage, fmt.Errorf("error sending 'CreateImage' request: %w", errCreateImage)
	}

	if httpResponseCreateImage.StatusCode < 200 || httpResponseCreateImage.StatusCode > 299 {
		return v0CreateImage, __rt.NewFutureResponseError("CreateImage", httpResponseCreateImage)
	}

	if errCreateImage = responseCreateImage.FromResponse("CreateImage", httpResponseCreateImage); errCreateImage != nil {
		return v0CreateImage, fmt.Errorf("error converting 'CreateImage' response: %w", errCreateImage)
	}

	addrCreateImage.Reset()
	headerCreateImage.Reset()

	if errCreateImage = responseCreateImage.Err(); errCreateImage != nil {
		return v0CreateImage, fmt.Errorf("error returned from 'CreateImage' response: %w", errCreateImage)
	}

	if errCreateImage = responseCreateImage.ScanValues(v0CreateImage); errCreateImage != nil {
		return v0CreateImage, fmt.Errorf("error scanning value from 'CreateImage' response: %w", errCreateImage)
	}

	return v0CreateImage, nil
}

func (__imp *implClient[C]) CreateImageEdit(ctx context.Context, request *CreateImageEditRequest) (*Images, error) {
	var innerCreateImageEdit any = __imp.Inner()

	addrCreateImageEdit := __rt.GetBuffer()
	defer __rt.PutBuffer(addrCreateImageEdit)
	defer addrCreateImageEdit.Reset()

	headerCreateImageEdit := __rt.GetBuffer()
	defer __rt.PutBuffer(headerCreateImageEdit)
	defer headerCreateImageEdit.Reset()

	var (
		v0CreateImageEdit           = new(Images)
		errCreateImageEdit          error
		httpResponseCreateImageEdit *http.Response
		responseCreateImageEdit     __rt.FutureResponse = __imp.response()
	)

	if errCreateImageEdit = addrTmplCreateImageEdit.Execute(addrCreateImageEdit, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateImageEdit != nil {
		return v0CreateImageEdit, fmt.Errorf("error building 'CreateImageEdit' url: %w", errCreateImageEdit)
	}

	if errCreateImageEdit = headerTmplCreateImageEdit.Execute(headerCreateImageEdit, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateImageEdit != nil {
		return v0CreateImageEdit, fmt.Errorf("error building 'CreateImageEdit' header: %w", errCreateImageEdit)
	}
	bufReaderCreateImageEdit := bufio.NewReader(headerCreateImageEdit)
	mimeHeaderCreateImageEdit, errCreateImageEdit := textproto.NewReader(bufReaderCreateImageEdit).ReadMIMEHeader()
	if errCreateImageEdit != nil {
		return v0CreateImageEdit, fmt.Errorf("error reading 'CreateImageEdit' header: %w", errCreateImageEdit)
	}

	urlCreateImageEdit := addrCreateImageEdit.String()
	requestCreateImageEdit, errCreateImageEdit := http.NewRequestWithContext(ctx, "POST", urlCreateImageEdit, request)
	if errCreateImageEdit != nil {
		return v0CreateImageEdit, fmt.Errorf("error building 'CreateImageEdit' request: %w", errCreateImageEdit)
	}

	for kCreateImageEdit, vvCreateImageEdit := range mimeHeaderCreateImageEdit {
		for _, vCreateImageEdit := range vvCreateImageEdit {
			requestCreateImageEdit.Header.Add(kCreateImageEdit, vCreateImageEdit)
		}
	}

	startCreateImageEdit := time.Now()

	if httpClientCreateImageEdit, okCreateImageEdit := innerCreateImageEdit.(interface{ Client() *http.Client }); okCreateImageEdit {
		httpResponseCreateImageEdit, errCreateImageEdit = httpClientCreateImageEdit.Client().Do(requestCreateImageEdit)
	} else {
		httpResponseCreateImageEdit, errCreateImageEdit = http.DefaultClient.Do(requestCreateImageEdit)
	}

	if logCreateImageEdit, okCreateImageEdit := innerCreateImageEdit.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okCreateImageEdit {
		logCreateImageEdit.Log(ctx, "CreateImageEdit", requestCreateImageEdit, httpResponseCreateImageEdit, time.Since(startCreateImageEdit))
	}

	if errCreateImageEdit != nil {
		return v0CreateImageEdit, fmt.Errorf("error sending 'CreateImageEdit' request: %w", errCreateImageEdit)
	}

	if httpResponseCreateImageEdit.StatusCode < 200 || httpResponseCreateImageEdit.StatusCode > 299 {
		return v0CreateImageEdit, __rt.NewFutureResponseError("CreateImageEdit", httpResponseCreateImageEdit)
	}

	if errCreateImageEdit = responseCreateImageEdit.FromResponse("CreateImageEdit", httpResponseCreateImageEdit); errCreateImageEdit != nil {
		return v0CreateImageEdit, fmt.Errorf("error converting 'CreateImageEdit' response: %w", errCreateImageEdit)
	}

	addrCreateImageEdit.Reset()
	headerCreateImageEdit.Reset()

	if errCreateImageEdit = responseCreateImageEdit.Err(); errCreateImageEdit != nil {
		return v0CreateImageEdit, fmt.Errorf("error returned from 'CreateImageEdit' response: %w", errCreateImageEdit)
	}

	if errCreateImageEdit = responseCreateImageEdit.ScanValues(v0CreateImageEdit); errCreateImageEdit != nil {
		return v0CreateImageEdit, fmt.Errorf("error scanning value from 'CreateImageEdit' response: %w", errCreateImageEdit)
	}

	return v0CreateImageEdit, nil
}

func (__imp *implClient[C]) CreateImageVariation(ctx context.Context, request *CreateImageVariationRequest) (*Images, error) {
	var innerCreateImageVariation any = __imp.Inner()

	addrCreateImageVariation := __rt.GetBuffer()
	defer __rt.PutBuffer(addrCreateImageVariation)
	defer addrCreateImageVariation.Reset()

	headerCreateImageVariation := __rt.GetBuffer()
	defer __rt.PutBuffer(headerCreateImageVariation)
	defer headerCreateImageVariation.Reset()

	var (
		v0CreateImageVariation           = new(Images)
		errCreateImageVariation          error
		httpResponseCreateImageVariation *http.Response
		responseCreateImageVariation     __rt.FutureResponse = __imp.response()
	)

	if errCreateImageVariation = addrTmplCreateImageVariation.Execute(addrCreateImageVariation, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateImageVariation != nil {
		return v0CreateImageVariation, fmt.Errorf("error building 'CreateImageVariation' url: %w", errCreateImageVariation)
	}

	if errCreateImageVariation = headerTmplCreateImageVariation.Execute(headerCreateImageVariation, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errCreateImageVariation != nil {
		return v0CreateImageVariation, fmt.Errorf("error building 'CreateImageVariation' header: %w", errCreateImageVariation)
	}
	bufReaderCreateImageVariation := bufio.NewReader(headerCreateImageVariation)
	mimeHeaderCreateImageVariation, errCreateImageVariation := textproto.NewReader(bufReaderCreateImageVariation).ReadMIMEHeader()
	if errCreateImageVariation != nil {
		return v0CreateImageVariation, fmt.Errorf("error reading 'CreateImageVariation' header: %w", errCreateImageVariation)
	}

	urlCreateImageVariation := addrCreateImageVariation.String()
	requestCreateImageVariation, errCreateImageVariation := http.NewRequestWithContext(ctx, "POST", urlCreateImageVariation, request)
	if errCreateImageVariation != nil {
		return v0CreateImageVariation, fmt.Errorf("error building 'CreateImageVariation' request: %w", errCreateImageVariation)
	}

	for kCreateImageVariation, vvCreateImageVariation := range mimeHeaderCreateImageVariation {
		for _, vCreateImageVariation := range vvCreateImageVariation {
			requestCreateImageVariation.Header.Add(kCreateImageVariation, vCreateImageVariation)
		}
	}

	startCreateImageVariation := time.Now()

	if httpClientCreateImageVariation, okCreateImageVariation := innerCreateImageVariation.(interface{ Client() *http.Client }); okCreateImageVariation {
		httpResponseCreateImageVariation, errCreateImageVariation = httpClientCreateImageVariation.Client().Do(requestCreateImageVariation)
	} else {
		httpResponseCreateImageVariation, errCreateImageVariation = http.DefaultClient.Do(requestCreateImageVariation)
	}

	if logCreateImageVariation, okCreateImageVariation := innerCreateImageVariation.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okCreateImageVariation {
		logCreateImageVariation.Log(ctx, "CreateImageVariation", requestCreateImageVariation, httpResponseCreateImageVariation, time.Since(startCreateImageVariation))
	}

	if errCreateImageVariation != nil {
		return v0CreateImageVariation, fmt.Errorf("error sending 'CreateImageVariation' request: %w", errCreateImageVariation)
	}

	if httpResponseCreateImageVariation.StatusCode < 200 || httpResponseCreateImageVariation.StatusCode > 299 {
		return v0CreateImageVariation, __rt.NewFutureResponseError("CreateImageVariation", httpResponseCreateImageVariation)
	}

	if errCreateImageVariation = responseCreateImageVariation.FromResponse("CreateImageVariation", httpResponseCreateImageVariation); errCreateImageVariation != nil {
		return v0CreateImageVariation, fmt.Errorf("error converting 'CreateImageVariation' response: %w", errCreateImageVariation)
	}

	addrCreateImageVariation.Reset()
	headerCreateImageVariation.Reset()

	if errCreateImageVariation = responseCreateImageVariation.Err(); errCreateImageVariation != nil {
		return v0CreateImageVariation, fmt.Errorf("error returned from 'CreateImageVariation' response: %w", errCreateImageVariation)
	}

	if errCreateImageVariation = responseCreateImageVariation.ScanValues(v0CreateImageVariation); errCreateImageVariation != nil {
		return v0CreateImageVariation, fmt.Errorf("error scanning value from 'CreateImageVariation' response: %w", errCreateImageVariation)
	}

	return v0CreateImageVariation, nil
}

func (__imp *implClient[C]) UploadFile(ctx context.Context, request *UploadFileRequest) (*File, error) {
	var innerUploadFile any = __imp.Inner()

	addrUploadFile := __rt.GetBuffer()
	defer __rt.PutBuffer(addrUploadFile)
	defer addrUploadFile.Reset()

	headerUploadFile := __rt.GetBuffer()
	defer __rt.PutBuffer(headerUploadFile)
	defer headerUploadFile.Reset()

	var (
		v0UploadFile           = new(File)
		errUploadFile          error
		httpResponseUploadFile *http.Response
		responseUploadFile     __rt.FutureResponse = __imp.response()
	)

	if errUploadFile = addrTmplUploadFile.Execute(addrUploadFile, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error building 'UploadFile' url: %w", errUploadFile)
	}

	if errUploadFile = headerTmplUploadFile.Execute(headerUploadFile, map[string]any{
		"Client":  __imp.Inner(),
		"ctx":     ctx,
		"request": request,
	}); errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error building 'UploadFile' header: %w", errUploadFile)
	}
	bufReaderUploadFile := bufio.NewReader(headerUploadFile)
	mimeHeaderUploadFile, errUploadFile := textproto.NewReader(bufReaderUploadFile).ReadMIMEHeader()
	if errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error reading 'UploadFile' header: %w", errUploadFile)
	}

	urlUploadFile := addrUploadFile.String()
	requestUploadFile, errUploadFile := http.NewRequestWithContext(ctx, "POST", urlUploadFile, request)
	if errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error building 'UploadFile' request: %w", errUploadFile)
	}

	for kUploadFile, vvUploadFile := range mimeHeaderUploadFile {
		for _, vUploadFile := range vvUploadFile {
			requestUploadFile.Header.Add(kUploadFile, vUploadFile)
		}
	}

	startUploadFile := time.Now()

	if httpClientUploadFile, okUploadFile := innerUploadFile.(interface{ Client() *http.Client }); okUploadFile {
		httpResponseUploadFile, errUploadFile = httpClientUploadFile.Client().Do(requestUploadFile)
	} else {
		httpResponseUploadFile, errUploadFile = http.DefaultClient.Do(requestUploadFile)
	}

	if logUploadFile, okUploadFile := innerUploadFile.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okUploadFile {
		logUploadFile.Log(ctx, "UploadFile", requestUploadFile, httpResponseUploadFile, time.Since(startUploadFile))
	}

	if errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error sending 'UploadFile' request: %w", errUploadFile)
	}

	if httpResponseUploadFile.StatusCode < 200 || httpResponseUploadFile.StatusCode > 299 {
		return v0UploadFile, __rt.NewFutureResponseError("UploadFile", httpResponseUploadFile)
	}

	if errUploadFile = responseUploadFile.FromResponse("UploadFile", httpResponseUploadFile); errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error converting 'UploadFile' response: %w", errUploadFile)
	}

	addrUploadFile.Reset()
	headerUploadFile.Reset()

	if errUploadFile = responseUploadFile.Err(); errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error returned from 'UploadFile' response: %w", errUploadFile)
	}

	if errUploadFile = responseUploadFile.ScanValues(v0UploadFile); errUploadFile != nil {
		return v0UploadFile, fmt.Errorf("error scanning value from 'UploadFile' response: %w", errUploadFile)
	}

	return v0UploadFile, nil
}

func (__imp *implClient[C]) RetrieveFileContent(ctx context.Context, fileID string) (io.ReadCloser, string, error) {
	var innerRetrieveFileContent any = __imp.Inner()

	addrRetrieveFileContent := __rt.GetBuffer()
	defer __rt.PutBuffer(addrRetrieveFileContent)
	defer addrRetrieveFileContent.Reset()

	headerRetrieveFileContent := __rt.GetBuffer()
	defer __rt.PutBuffer(headerRetrieveFileContent)
	defer headerRetrieveFileContent.Reset()

	var (
		v0RetrieveFileContent           = __rt.New[io.ReadCloser]()
		v1RetrieveFileContent           = __rt.New[string]()
		errRetrieveFileContent          error
		httpResponseRetrieveFileContent *http.Response
		responseRetrieveFileContent     __rt.FutureResponse = __imp.response()
	)

	if errRetrieveFileContent = addrTmplRetrieveFileContent.Execute(addrRetrieveFileContent, map[string]any{
		"Client": __imp.Inner(),
		"ctx":    ctx,
		"fileID": fileID,
	}); errRetrieveFileContent != nil {
		return v0RetrieveFileContent, v1RetrieveFileContent, fmt.Errorf("error building 'RetrieveFileContent' url: %w", errRetrieveFileContent)
	}

	if errRetrieveFileContent = headerTmplRetrieveFileContent.Execute(headerRetrieveFileContent, map[string]any{
		"Client": __imp.Inner(),
		"ctx":    ctx,
		"fileID": fileID,
	}); errRetrieveFileContent != nil {
		return v0RetrieveFileContent, v1RetrieveFileContent, fmt.Errorf("error building 'RetrieveFileContent' header: %w", errRetrieveFileContent)
	}
	bufReaderRetrieveFileContent := bufio.NewReader(headerRetrieveFileContent)
	mimeHeaderRetrieveFileContent, errRetrieveFileContent := textproto.NewReader(bufReaderRetrieveFileContent).ReadMIMEHeader()
	if errRetrieveFileContent != nil {
		return v0RetrieveFileContent, v1RetrieveFileContent, fmt.Errorf("error reading 'RetrieveFileContent' header: %w", errRetrieveFileContent)
	}

	urlRetrieveFileContent := addrRetrieveFileContent.String()
	requestRetrieveFileContent, errRetrieveFileContent := http.NewRequestWithContext(ctx, "GET", urlRetrieveFileContent, http.NoBody)
	if errRetrieveFileContent != nil {
		return v0RetrieveFileContent, v1RetrieveFileContent, fmt.Errorf("error building 'RetrieveFileContent' request: %w", errRetrieveFileContent)
	}

	for kRetrieveFileContent, vvRetrieveFileContent := range mimeHeaderRetrieveFileContent {
		for _, vRetrieveFileContent := range vvRetrieveFileContent {
			requestRetrieveFileContent.Header.Add(kRetrieveFileContent, vRetrieveFileContent)
		}
	}

	startRetrieveFileContent := time.Now()

	if httpClientRetrieveFileContent, okRetrieveFileContent := innerRetrieveFileContent.(interface{ Client() *http.Client }); okRetrieveFileContent {
		httpResponseRetrieveFileContent, errRetrieveFileContent = httpClientRetrieveFileContent.Client().Do(requestRetrieveFileContent)
	} else {
		httpResponseRetrieveFileContent, errRetrieveFileContent = http.DefaultClient.Do(requestRetrieveFileContent)
	}

	if logRetrieveFileContent, okRetrieveFileContent := innerRetrieveFileContent.(interface {
		Log(ctx context.Context, caller string, request *http.Request, response *http.Response, elapse time.Duration)
	}); okRetrieveFileContent {
		logRetrieveFileContent.Log(ctx, "RetrieveFileContent", requestRetrieveFileContent, httpResponseRetrieveFileContent, time.Since(startRetrieveFileContent))
	}

	if errRetrieveFileContent != nil {
		return v0RetrieveFileContent, v1RetrieveFileContent, fmt.Errorf("error sending 'RetrieveFileContent' request: %w", errRetrieveFileContent)
	}

	if httpResponseRetrieveFileContent.StatusCode < 200 || httpResponseRetrieveFileContent.StatusCode > 299 {
		return v0RetrieveFileContent, v1RetrieveFileContent, __rt.NewFutureResponseError("RetrieveFileContent", httpResponseRetrieveFileContent)
	}

	if errRetrieveFileContent = responseRetrieveFileContent.FromResponse("RetrieveFileContent", httpResponseRetrieveFileContent); errRetrieveFileContent != nil {
		return v0RetrieveFileContent, v1RetrieveFileContent, fmt.Errorf("error converting 'RetrieveFileContent' response: %w", errRetrieveFileContent)
	}

	addrRetrieveFileContent.Reset()
	headerRetrieveFileContent.Reset()

	if errRetrieveFileContent = responseRetrieveFileContent.Err(); errRetrieveFileContent != nil {
		return v0RetrieveFileContent, v1RetrieveFileContent, fmt.Errorf("error returned from 'RetrieveFileContent' response: %w", errRetrieveFileContent)
	}

	if errRetrieveFileContent = responseRetrieveFileContent.ScanValues(&v0RetrieveFileContent, &v1RetrieveFileContent); errRetrieveFileContent != nil {
		return v0RetrieveFileContent, v1RetrieveFileContent, fmt.Errorf("error scanning value from 'RetrieveFileContent' response: %w", errRetrieveFileContent)
	}

	return v0RetrieveFileContent, v1RetrieveFileContent, nil
}

func (__imp *implClient[C]) Inner() C {
	return __imp.__Client
}

func (*implClient[C]) response() *ResponseHandler {
	return new(ResponseHandler)
}
