package multipart

import (
	"bytes"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"strconv"
	"strings"

	c "gopkg.in/h2non/gentleman.v2/context"
	p "gopkg.in/h2non/gentleman.v2/plugin"
)

// Values represents multiple multipart from values.
type Values []string

// DataFields represents a map of text based fields.
type DataFields map[string]Values

// FormFile represents the file form field data.
type FormFile struct {
	Name   string
	Reader io.Reader
}

// FormData represents the supported form fields by file and string data.
type FormData struct {
	Data  DataFields
	Files []FormFile
}

// File creates a new multipart form based on a unique file field
// from the given io.ReadCloser stream.
func File(name string, reader io.Reader) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		file := FormFile{name, reader}
		data := FormData{Files: []FormFile{file}}
		handle(ctx, h, data)
	})
}

// Files creates a multipart form based on files fields.
func Files(files []FormFile) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		data := FormData{Files: files}
		handle(ctx, h, data)
	})
}

// Fields creates a new multipart form based on string based fields.
func Fields(fields DataFields) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		data := FormData{Data: fields}
		handle(ctx, h, data)
	})
}

// Data creates custom form based on the given form data
// who can have files and string based fields.
func Data(data FormData) p.Plugin {
	return p.NewRequestPlugin(func(ctx *c.Context, h c.Handler) {
		handle(ctx, h, data)
	})
}

func handle(ctx *c.Context, h c.Handler, data FormData) {
	if err := createForm(data, ctx); err != nil {
		h.Error(ctx, err)
		return
	}
	h.Next(ctx)
}

func createForm(data FormData, ctx *c.Context) error {
	body := &bytes.Buffer{}
	multipartWriter := multipart.NewWriter(body)

	for index, file := range data.Files {
		if err := writeFile(multipartWriter, data, file, index); err != nil {
			return err
		}
	}

	// Populate the other parts of the form (if there are any)
	for key, values := range data.Data {
		for _, value := range values {
			multipartWriter.WriteField(key, value)
		}
	}
	if err := multipartWriter.Close(); err != nil {
		return err
	}

	ctx.Request.Method = setMethod(ctx)
	ctx.Request.Body = ioutil.NopCloser(body)
	ctx.Request.Header.Add("Content-Type", multipartWriter.FormDataContentType())

	return nil
}

func writeFile(multipartWriter *multipart.Writer, data FormData, file FormFile, index int) error {
	if file.Reader == nil {
		return errors.New("gentleman: file reader cannot be nil")
	}

	rc, ok := file.Reader.(io.ReadCloser)
	if !ok && file.Reader != nil {
		rc = ioutil.NopCloser(file.Reader)
	}

	fileName := "file"
	if len(data.Files) > 1 {
		fileName = strings.Join([]string{fileName, strconv.Itoa(index + 1)}, "")
	}
	if file.Name != "" {
		fileName = file.Name
	}

	writer, err := multipartWriter.CreateFormFile(fileName, file.Name)
	if err != nil {
		return err
	}
	if _, err = io.Copy(writer, rc); err != nil && err != io.EOF {
		return err
	}
	rc.Close()

	return nil
}

func setMethod(ctx *c.Context) string {
	method := ctx.Request.Method
	if method == "GET" || method == "" {
		return "POST"
	}
	return method
}
