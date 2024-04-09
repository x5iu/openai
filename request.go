package openai

import (
	"errors"
	defc "github.com/x5iu/defc/runtime"
	"io"
	"mime/multipart"
	"runtime"
	"sync"
)

type ChatCompletionRequest struct {
	defc.JSONBody[ChatCompletionRequest]
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
	Temperature      NullableType[float64] `json:"temperature,omitempty"`
	TopP             NullableType[float64] `json:"top_p,omitempty"`
	Tools            Tools                 `json:"tools,omitempty"`
	ToolChoice       ToolChoice            `json:"tool_choice,omitempty"`
	User             string                `json:"user,omitempty"`
}

type CreateImageRequest struct {
	defc.JSONBody[CreateImageRequest]
	Model   string `json:"model"`
	Prompt  string `json:"prompt"`
	N       int    `json:"n,omitempty"`
	Size    string `json:"size,omitempty"`
	Quality string `json:"quality,omitempty"`
}

type UploadFileRequest struct {
	File     io.Reader
	Filename string
	Purpose  string

	once        sync.Once
	reader      *io.PipeReader
	errorChan   <-chan error
	contentType string
}

func (r *UploadFileRequest) init() {
	r.once.Do(func() {
		const (
			formFieldPurpose = "purpose"
			formFieldFile    = "file"
		)

		var (
			err                    error
			pipeReader, pipeWriter = io.Pipe()
			writer                 = multipart.NewWriter(pipeWriter)
			errorChan              = make(chan error, 1)
		)

		r.reader = pipeReader
		r.contentType = writer.FormDataContentType()
		r.errorChan = errorChan

		// To prevent goroutine leaks, utilize the SetFinalizer function to tidy up the Producer-side goroutine
		// upon the completion of the UploadFileRequest lifecycle.
		runtime.SetFinalizer(r, func(*UploadFileRequest) { pipeReader.Close() })

		go func() {
			defer func(ch chan<- error) {
				if err != nil {
					if !errors.Is(err, io.ErrClosedPipe) {
						ch <- err
					}
				}
				close(ch)
				pipeWriter.Close()
			}(errorChan)
			if err = writer.WriteField(formFieldPurpose, r.Purpose); err != nil {
				return
			}
			var filename string
			if naming, ok := r.File.(interface{ Name() string }); ok { // os.File has Name() method
				filename = naming.Name()
			}
			if r.Filename != "" {
				filename = r.Filename
			}
			var part io.Writer
			part, err = writer.CreateFormFile(formFieldFile, filename)
			if err != nil {
				return
			}
			if _, err = io.Copy(part, r.File); err != nil {
				return
			}
			if err = writer.Close(); err != nil {
				return
			}
			if err = pipeWriter.Close(); err != nil {
				return
			}
		}()
	})
}

func (r *UploadFileRequest) ContentType() string {
	r.init()
	return r.contentType
}

func (r *UploadFileRequest) Read(p []byte) (n int, err error) {
	r.init()
	for {
		var (
			alreadyEOF bool
			chanClosed bool
		)
		{
			var (
				chanActive bool
			)
			select {
			case err, chanActive = <-r.errorChan:
				chanClosed = !chanActive
				if chanActive {
					if err != nil {
						return n, err
					}
				}
			default:
			}
		}
		if alreadyEOF && chanClosed {
			return n, io.EOF
		}
		n, err = r.reader.Read(p)
		if err != nil && errors.Is(err, io.EOF) && !chanClosed {
			alreadyEOF = true
			continue
		}
		return n, err
	}
}
