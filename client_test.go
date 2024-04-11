package openai

import (
	"bytes"
	"context"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

var (
	testKey         = "sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	testServer      *httptest.Server
	testFileMD5Hash []byte
)

func TestMain(m *testing.M) {
	hasher := md5.New()
	file, err := os.Open("client_test.go")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(hasher, file)
	testFileMD5Hash = hasher.Sum(nil)
	testServer = newTestServer()
	testServer.StartTLS()
	code := m.Run()
	testServer.Close()
	os.Exit(code)
}

func newTestServer() *httptest.Server {
	mux := http.NewServeMux()
	mux.Handle("/models", &handler{
		KeyResponse: keyResponse{
			plainTextKey: {
				StatusCode:   500,
				ContentType:  "text/plain",
				ResponseBody: "plain_text",
			},
			jsonArrayKey: {
				StatusCode:   500,
				ContentType:  "application/json",
				ResponseBody: "[]",
			},
		},
		ResponseData: &Models{
			Data: []struct {
				ID         string            `json:"id"`
				Object     string            `json:"object"`
				OwnedBy    string            `json:"owned_by"`
				Permission []json.RawMessage `json:"permission"`
			}{
				{
					ID:         "gpt-3.5-turbo",
					Object:     "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
					OwnedBy:    "openai",
					Permission: nil,
				},
			},
			Object: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		},
	})
	mux.Handle("/chat/completions", &handler{
		KeyResponse: keyResponse{
			streamStopKey: {
				StatusCode:  200,
				ContentType: "text/event-stream",
				ResponseBody: `data: {"id": "xxxx-xxxx-xxxx", "model": "gpt-3.5-turbo", "object": "xxxx", "created": 1700000000, "choices": [{"index": 0, "delta": {"role": "assistant", "content": "h", "tool_calls": [{"index": 0, "id": "xxxx", "type": "function", "function": {"name": "test", "arguments": "te"}}], "finish_reason": ""}}]}` + "\n" +
					`data: {"id": "xxxx-xxxx-xxxx", "model": "gpt-3.5-turbo", "object": "xxxx", "created": 1700000000, "choices": [{"index": 0, "delta": {"role": "assistant", "content": "i", "tool_calls": [{"index": 0, "id": "", "type": "", "function": {"name": "", "arguments": "st"}}], "finish_reason": "stop"}}]}` + "\n" +
					`data: [DONE]` + "\n",
			},
			streamDoneKey: {
				StatusCode:  200,
				ContentType: "text/event-stream",
				ResponseBody: `data: {"id": "xxxx-xxxx-xxxx", "model": "gpt-3.5-turbo", "object": "xxxx", "created": 1700000000, "choices": [{"index": 0, "delta": {"role": "assistant", "content": "h", "tool_calls": [{"index": 0, "id": "xxxx", "type": "function", "function": {"name": "test", "arguments": "te"}}], "finish_reason": ""}}]}` + "\n" +
					`data: {"id": "xxxx-xxxx-xxxx", "model": "gpt-3.5-turbo", "object": "xxxx", "created": 1700000000, "choices": [{"index": 0, "delta": {"role": "assistant", "content": "i", "tool_calls": [{"index": 0, "id": "", "type": "", "function": {"name": "", "arguments": "st"}}], "finish_reason": "length"}}]}` + "\n" +
					`data: [DONE]` + "\n",
			},
			invalidStreamTypeKey: {
				StatusCode:   200,
				ContentType:  "text/plain",
				ResponseBody: "data: [DONE]",
			},
			invalidStreamContentKey: {
				StatusCode:   200,
				ContentType:  "text/event-stream",
				ResponseBody: "data: []",
			},
		},
		ResponseData: &Completion{
			ID:                "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Model:             "gpt-3.5-turbo",
			Object:            "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Created:           1700000000,
			SystemFingerprint: "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
			Choices: []struct {
				Index        NullableType[int]    `json:"index"`
				Message      *Message             `json:"message"`
				FinishReason NullableType[string] `json:"finish_reason"`
				Logprobs     json.RawMessage      `json:"logprobs"`
			}{
				{
					Index: "0",
					Message: &Message{
						Role:      RoleAssistant,
						Content:   &Content{Text: "hi"},
						ToolCalls: nil,
					},
					FinishReason: "stop",
					Logprobs:     nil,
				},
			},
			Usage: struct {
				PromptTokens     int `json:"prompt_tokens"`
				CompletionTokens int `json:"completion_tokens"`
				TotalTokens      int `json:"total_tokens"`
			}{
				PromptTokens:     10,
				CompletionTokens: 5,
				TotalTokens:      15,
			},
		},
	})
	mux.Handle("/images/generations", &handler{
		KeyResponse: keyResponse{},
		ResponseData: &Images{
			Created: 1700000000,
			Data: []struct {
				RevisedPrompt string `json:"revised_prompt"`
				Url           string `json:"url"`
			}{
				{
					RevisedPrompt: "test",
					Url:           "oss://oss.test/image/test",
				},
			},
		},
	})
	mux.Handle("/files", http.HandlerFunc(uploadFile))
	mux.Handle("/files/test_file_xxx/content", http.HandlerFunc(retrieveFileContent))
	mux.Handle("/images/edits", http.HandlerFunc(imageHandler))
	mux.Handle("/images/variations", http.HandlerFunc(imageHandler))
	return httptest.NewUnstartedServer(mux)
}

var (
	plainTextKey            = "plain_text_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	jsonArrayKey            = "json_array_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	streamStopKey           = "stream_stop_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	streamDoneKey           = "stream_done_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	invalidStreamTypeKey    = "invalid_stream_type_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
	invalidStreamContentKey = "invalid_stream_content_xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
)

type (
	returnType struct {
		StatusCode   int
		ContentType  string
		ResponseBody string
	}
	keyResponse map[string]*returnType
)

type handler struct {
	KeyResponse  keyResponse
	ResponseData any
}

func (h *handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	var errorResponse struct {
		Error *Error `json:"error"`
	}
	if rt, ok := checkAuthorization(r.Header, h.KeyResponse); !ok {
		errorResponse.Error = &Error{
			Message: "Unauthorized",
			Type:    "unauthorized_error",
			Param:   "unauthorized_error",
			Code:    "unauthorized_error",
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&errorResponse)
		return
	} else if rt != nil {
		w.Header().Set("Content-Type", rt.ContentType)
		w.WriteHeader(rt.StatusCode)
		io.WriteString(w, rt.ResponseBody)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(h.ResponseData)
	return
}

func checkAuthorization(header http.Header, keyResponse keyResponse) (*returnType, bool) {
	authorization := header.Get("Authorization")
	authorization = strings.TrimSpace(strings.TrimPrefix(authorization, "Bearer"))
	if authorization == testKey {
		return nil, true
	}
	for k, r := range keyResponse {
		if authorization == k {
			return r, true
		}
	}
	return nil, false
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, ok := checkAuthorization(r.Header, keyResponse{}); !ok {
		var unauthorizedError = struct {
			Error *Error `json:"error"`
		}{
			Error: &Error{
				Message: "Unauthorized",
				Type:    "unauthorized_error",
				Param:   "unauthorized_error",
				Code:    "unauthorized_error",
			},
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&unauthorizedError)
		return
	}
	var internalServerError struct {
		Error *Error `json:"error"`
	}
	internalServerError.Error = &Error{
		Message: "Internal Server Error",
		Type:    "internal_server_error",
		Param:   "internal_server_error",
		Code:    "internal_server_error",
	}
	if err := r.ParseMultipartForm(1024 * 1024); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&internalServerError)
		return
	}
	file, header, err := r.FormFile("image")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&internalServerError)
		return
	}
	defer file.Close()
	if header.Filename != "image.png" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&internalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&Images{
		Created: 1700000000,
		Data: []struct {
			RevisedPrompt string `json:"revised_prompt"`
			Url           string `json:"url"`
		}{
			{
				RevisedPrompt: "test",
				Url:           "oss://oss.test/image/test",
			},
		},
	})
}

func uploadFile(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	if _, ok := checkAuthorization(r.Header, keyResponse{}); !ok {
		var unauthorizedError = struct {
			Error *Error `json:"error"`
		}{
			Error: &Error{
				Message: "Unauthorized",
				Type:    "unauthorized_error",
				Param:   "unauthorized_error",
				Code:    "unauthorized_error",
			},
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&unauthorizedError)
		return
	}
	var internalServerError struct {
		Error *Error `json:"error"`
	}
	internalServerError.Error = &Error{
		Message: "Internal Server Error",
		Type:    "internal_server_error",
		Param:   "internal_server_error",
		Code:    "internal_server_error",
	}
	if err := r.ParseMultipartForm(1024 * 1024); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&internalServerError)
		return
	}
	purpose := r.FormValue("purpose")
	if purpose == "" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&internalServerError)
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&internalServerError)
		return
	}
	defer file.Close()
	if header.Filename != "client_test.go" {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&internalServerError)
		return
	}
	hasher := md5.New()
	io.Copy(hasher, file)
	uploadFileMD5Hash := hasher.Sum(nil)
	if !bytes.Equal(testFileMD5Hash, uploadFileMD5Hash) {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(&internalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&File{
		ID:        "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Object:    "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx",
		Bytes:     1024,
		CreatedAt: 1700000000,
		Filename:  "client_test.go",
		Purpose:   purpose,
	})
}

func retrieveFileContent(w http.ResponseWriter, r *http.Request) {
	if _, ok := checkAuthorization(r.Header, keyResponse{}); !ok {
		var unauthorizedError = struct {
			Error *Error `json:"error"`
		}{
			Error: &Error{
				Message: "Unauthorized",
				Type:    "unauthorized_error",
				Param:   "unauthorized_error",
				Code:    "unauthorized_error",
			},
		}
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(&unauthorizedError)
		return
	}
	file, _ := os.Open("client_test.go")
	defer file.Close()
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusOK)
	io.Copy(w, file)
}

type testCaller struct {
	baseUrl string
	apiKey  string
	client  *http.Client
}

func (tc *testCaller) BaseUrl() string      { return tc.baseUrl }
func (tc *testCaller) APIKey() string       { return tc.apiKey }
func (tc *testCaller) Client() *http.Client { return tc.client }

func newClient(key string) Client[*testCaller] {
	return NewClient[*testCaller](&testCaller{
		baseUrl: testServer.URL + "/",
		apiKey:  key,
		client:  testServer.Client(),
	})
}

func TestClient(t *testing.T) {
	t.Run("error", testClientError)
	t.Run("models", testClientModels)
	t.Run("chat", testClientChat)
	t.Run("image", testClientImage)
	t.Run("file", testClientFile)
}

func testClientError(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		brokenClient := newClient(plainTextKey)
		_, err := brokenClient.ListModels(context.Background())
		if err == nil {
			t.Errorf("brokenClient.ListModels: expects errors, got nothing")
			return
		}
		defer CloseErrorResponseBody(err)
		parsedError := ParseError(err)
		if parsedError != nil {
			serialized, _ := json.Marshal(parsedError)
			t.Errorf("ParseError: not Error, but got => %s", string(serialized))
			return
		}
		response := err.(getResponse).Response()
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)
		if !bytes.Equal(body, []byte("plain_text")) {
			t.Errorf("http.Response.Body: %s != plain_text", string(body))
			return
		}
	})
	t.Run("2", func(t *testing.T) {
		brokenClient := newClient(jsonArrayKey)
		_, err := brokenClient.ListModels(context.Background())
		if err == nil {
			t.Errorf("brokenClient.ListModels: expects errors, got nothing")
			return
		}
		defer CloseErrorResponseBody(err)
		parsedError := ParseError(err)
		if parsedError != nil {
			serialized, _ := json.Marshal(parsedError)
			t.Errorf("ParseError: not Error, but got => %s", string(serialized))
			return
		}
		response := err.(getResponse).Response()
		defer response.Body.Close()
		body, _ := io.ReadAll(response.Body)
		if !bytes.Equal(body, []byte("[]")) {
			t.Errorf("http.Response.Body: %s != []", string(body))
			return
		}
	})
}

func testClientModels(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := newClient(testKey)
		models, err := client.ListModels(context.Background())
		if err != nil {
			t.Errorf("client.ListModels: %s", err)
			return
		}
		if len(models.Data) != 1 || models.Data[0].ID != "gpt-3.5-turbo" {
			serialized, _ := json.Marshal(models)
			t.Errorf("client.ListModels: unexpected Models object => %s", string(serialized))
			return
		}
	})
	t.Run("fail", func(t *testing.T) {
		noKeyClient := newClient("")
		_, err := noKeyClient.ListModels(context.Background())
		if err == nil {
			t.Errorf("noKeyClient.ListModels: expects errors, got nothing")
			return
		}
		defer CloseErrorResponseBody(err)
		parsedError := ParseError(err)
		if parsedError == nil {
			t.Errorf("ParseError: expects Error, got => %s", err)
			return
		}
		if parsedError.Type != "unauthorized_error" {
			t.Errorf("parsedError.Type != unauthorized_error")
			return
		}
		if parsedError.Error() != "openai: Unauthorized" {
			t.Errorf("parsedError.Error() != Unauthorized")
			return
		}
	})
}

func testClientChat(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		client := newClient(testKey)
		chatCompletion, err := client.CreateChatCompletion(context.Background(), &ChatCompletionRequest{
			Messages: Messages{
				{
					Role:    RoleUser,
					Content: &Content{Text: "hi"},
				},
			},
			Model: "gpt-3.5-turbo",
		})
		if err != nil {
			t.Errorf("client.CreateChatCompletion: %s", err)
			return
		}
		completion := chatCompletion.(*Completion)
		if completion.Model != "gpt-3.5-turbo" || len(completion.Choices) != 1 ||
			completion.GetFinishReason() != FinishReasonStop || completion.GetMessageRole() != RoleAssistant ||
			completion.GetMessageContent() != "hi" || completion.GetPromptTokens() != 10 ||
			completion.GetCompletionTokens() != 5 || completion.GetTotalTokens() != 15 ||
			completion.GetPromptTokens()+completion.GetCompletionTokens() != completion.GetTotalTokens() {
			serialized, _ := json.Marshal(completion)
			t.Errorf("client.CreateChatCompletion: unexpected Completion object => %s", string(serialized))
			return
		}
		t.Run("stream", func(t *testing.T) {
			t.Run("stop", func(t *testing.T) {
				streamClient := newClient(streamStopKey)
				chatCompletion, err := streamClient.CreateChatCompletion(context.Background(), &ChatCompletionRequest{
					Messages: Messages{
						{
							Role:    RoleUser,
							Content: &Content{Text: "hi"},
						},
					},
					Model:          "gpt-3.5-turbo",
					ToolChoice:     "test",
					ResponseFormat: ResponseFormatText,
					Stream:         true,
				})
				if err != nil {
					t.Errorf("streamClient.CreateChatCompletionStream: %s", err)
					return
				}
				stream := chatCompletion.(*Stream)
				defer stream.Close()
				message := stream.GetMessage()
				if streamErr := stream.Err(); streamErr != nil {
					t.Errorf("stream.Err(): %s", streamErr)
					return
				}
				if message.Role != RoleAssistant || message.Content.Text != "hi" ||
					message.ToolCalls[0].Type != ToolTypeFunction || message.ToolCalls[0].Function.Name != "test" ||
					message.ToolCalls[0].Function.Arguments != "test" {
					serialized, _ := json.Marshal(message)
					t.Errorf("streamClient.CreateChatCompletionStream: unexpected Stream Message => %s", string(serialized))
					return
				}
			})
			t.Run("done", func(t *testing.T) {
				streamClient := newClient(streamDoneKey)
				chatCompletion, err := streamClient.CreateChatCompletion(context.Background(), &ChatCompletionRequest{
					Messages: Messages{
						{
							Role:    RoleUser,
							Content: &Content{Text: "hi"},
						},
					},
					Model:          "gpt-3.5-turbo",
					ToolChoice:     "test",
					ResponseFormat: ResponseFormatText,
					Stream:         true,
				})
				if err != nil {
					t.Errorf("streamClient.CreateChatCompletionStream: %s", err)
					return
				}
				stream := chatCompletion.(*Stream)
				defer stream.Close()
				message := stream.GetMessage()
				if streamErr := stream.Err(); streamErr != nil {
					t.Errorf("stream.Err(): %s", streamErr)
					return
				}
				if message.Role != RoleAssistant || message.Content.Text != "hi" ||
					message.ToolCalls[0].Type != ToolTypeFunction || message.ToolCalls[0].Function.Name != "test" ||
					message.ToolCalls[0].Function.Arguments != "test" {
					serialized, _ := json.Marshal(message)
					t.Errorf("streamClient.CreateChatCompletionStream: unexpected Stream Message => %s", string(serialized))
					return
				}
			})
			t.Run("error", func(t *testing.T) {
				t.Run("type", func(t *testing.T) {
					invalidStreamClient := newClient(invalidStreamTypeKey)
					_, streamErr := invalidStreamClient.CreateChatCompletion(context.Background(), &ChatCompletionRequest{
						Messages: Messages{
							{
								Role:    RoleUser,
								Content: &Content{Text: "hi"},
							},
						},
						Model:  "gpt-3.5-turbo",
						Stream: true,
					})
					if streamErr == nil {
						t.Errorf("invalidStreamClient.CreateChatCompletionStream: expects errors, got nothing")
						return
					}
					if !strings.Contains(streamErr.Error(), "unresolved Content-Type: ") {
						t.Errorf("streamErr != UnresolvedContentType => %s", streamErr)
						return
					}
				})
				t.Run("content", func(t *testing.T) {
					invalidStreamClient := newClient(invalidStreamContentKey)
					chatCompletion, streamErr := invalidStreamClient.CreateChatCompletion(context.Background(), &ChatCompletionRequest{
						Messages: Messages{
							{
								Role:    RoleUser,
								Content: &Content{Text: "hi"},
							},
						},
						Model:  "gpt-3.5-turbo",
						Stream: true,
					})
					if streamErr != nil {
						t.Errorf("invalidStreamClient.CreateChatCompletionStream: %s", streamErr)
						return
					}
					stream := chatCompletion.(*Stream)
					defer stream.Close()
					if streamErr = stream.Err(); streamErr != nil {
						t.Errorf("stream.Err(): %s", streamErr)
						return
					}
					stream.GetMessage()
					if streamErr = stream.Err(); streamErr == nil {
						t.Errorf("stream.Err() expects errors, got nothing")
						return
					}
				})
			})
		})
	})
}

func testClientImage(t *testing.T) {
	t.Run("generations", func(t *testing.T) {
		client := newClient(testKey)
		images, err := client.CreateImage(context.Background(), &CreateImageRequest{
			Model:  "dall-e-2",
			Prompt: "test",
			N:      1,
		})
		if err != nil {
			t.Errorf("client.CreateImage: %s", err)
			return
		}
		if len(images.Data) != 1 || images.Data[0].RevisedPrompt != "test" ||
			images.Data[0].Url != "oss://oss.test/image/test" {
			serialized, _ := json.Marshal(images)
			t.Errorf("client.CreateImage: unexpected Images object => %s", string(serialized))
			return
		}
	})
	t.Run("edits", func(t *testing.T) {
		client := newClient(testKey)
		images, err := client.CreateImageEdit(context.Background(), &CreateImageEditRequest{
			Image:          &namedReader{name: "image.png", reader: strings.NewReader("test")},
			Prompt:         "test",
			Mask:           &namedReader{name: "mask.png", reader: strings.NewReader("test")},
			Model:          ModelDALLE2,
			N:              1,
			Size:           ImageSize256x256,
			ResponseFormat: ResponseFormatUrl,
			User:           "test_user",
		})
		if err != nil {
			t.Errorf("client.CreateImageEdit: %s", err)
			return
		}
		if len(images.Data) != 1 || images.Data[0].RevisedPrompt != "test" ||
			images.Data[0].Url != "oss://oss.test/image/test" {
			serialized, _ := json.Marshal(images)
			t.Errorf("client.CreateImageEdit: unexpected Images object => %s", string(serialized))
			return
		}
	})
	t.Run("variation", func(t *testing.T) {
		client := newClient(testKey)
		images, err := client.CreateImageVariation(context.Background(), &CreateImageVariationRequest{
			Image:          &namedReader{name: "image.png", reader: strings.NewReader("test")},
			Model:          ModelDALLE2,
			N:              1,
			ResponseFormat: ResponseFormatUrl,
			Size:           ImageSize256x256,
			User:           "test_user",
		})
		if err != nil {
			t.Errorf("client.CreateImageVariation: %s", err)
			return
		}
		if len(images.Data) != 1 || images.Data[0].RevisedPrompt != "test" ||
			images.Data[0].Url != "oss://oss.test/image/test" {
			serialized, _ := json.Marshal(images)
			t.Errorf("client.CreateImageVariation: unexpected Images object => %s", string(serialized))
			return
		}
	})
}

func testClientFile(t *testing.T) {
	t.Run("upload", func(t *testing.T) {
		client := newClient(testKey)
		testFile, err := os.Open("client_test.go")
		if err != nil {
			t.Errorf("os.Open: %s", err)
			return
		}
		defer testFile.Close()
		file, err := client.UploadFile(context.Background(), &UploadFileRequest{
			File:    testFile,
			Purpose: "fine-tune",
		})
		if err != nil {
			t.Errorf("client.UploadFile: %s", err)
			return
		}
		if file.Filename != "client_test.go" || file.Purpose != "fine-tune" {
			serialized, _ := json.Marshal(file)
			t.Errorf("client.UploadFile: unexpected File object => %s", string(serialized))
			return
		}
	})
	t.Run("retrieve_file_content", func(t *testing.T) {
		client := newClient(testKey)
		reader, contentType, err := client.RetrieveFileContent(context.Background(), "test_file_xxx")
		if err != nil {
			t.Errorf("client.RetrieveFileContent: %s", err)
			return
		}
		hasher := md5.New()
		io.Copy(hasher, reader)
		contentMD5Hash := hasher.Sum(nil)
		if !bytes.Equal(testFileMD5Hash, contentMD5Hash) {
			t.Errorf("retrieve_file_content: hash(%s) != hash(%s)",
				hex.EncodeToString(testFileMD5Hash),
				hex.EncodeToString(contentMD5Hash))
			return
		}
		if contentType != "text/plain" {
			t.Errorf("Content-Type: %q != %q", contentType, "text/plain")
			return
		}
	})
}
