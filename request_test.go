package openai

import (
	"errors"
	"io"
	"os"
	"strings"
	"testing"
)

type closedReader struct{}

func (*closedReader) Name() string             { return "closed_reader" }
func (*closedReader) Read([]byte) (int, error) { return 0, os.ErrClosed }

type namedReader struct {
	reader io.Reader
}

func (nr *namedReader) Name() string                     { return "named_reader" }
func (nr *namedReader) Read(p []byte) (n int, err error) { return nr.reader.Read(p) }

func TestUploadFileRequest(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := &UploadFileRequest{
			File:    &namedReader{strings.NewReader("test")},
			Purpose: "fine-tune",
		}
		if _, err := io.ReadAll(r); err != nil {
			t.Errorf("UploadFileRequest: %s", err)
			return
		}
	})
	t.Run("error", func(t *testing.T) {
		r := &UploadFileRequest{
			File:    &closedReader{},
			Purpose: "fine-tune",
		}
		if _, err := io.ReadAll(r); err == nil {
			t.Errorf("UploadFileRequest: expects errors, got nil")
			return
		} else if !errors.Is(err, os.ErrClosed) {
			t.Errorf("UploadFileRequest: expects os.ErrClosed, got => %s", err)
			return
		}
	})
}
