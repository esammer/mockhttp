package responder

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

// An http.ResponseWriter that buffers body content.
//
// This type can be used to create http.Response instances by calling any combination of Header(), WriteHeader(int),
// and Write([]byte), followed by a final call to GetResponse(). It is illegal to call any methods once GetResponse()
// has been called.
//
// Unlike the real implementation used by the http package, this ResponseWriter uses a bytes.Buffer to capture any
// Write() output, making it available as the http.Response Body reader.
type ResponseWriter struct {
	header     http.Header
	statusCode int
	buf        *bytes.Buffer
}

// Create a new ResponseWriter.
func NewResponseWriter() *ResponseWriter {
	return &ResponseWriter{
		header: http.Header{},
		buf:    &bytes.Buffer{},
	}
}

func (r *ResponseWriter) Header() http.Header {
	return r.header
}

// Write b into the response.
//
// Like bytes.Buffer.Write(), this method will never return an error.
func (r *ResponseWriter) Write(b []byte) (int, error) {
	return r.buf.Write(b)
}

func (r *ResponseWriter) WriteHeader(statusCode int) {
	r.statusCode = statusCode
}

// Create and return an http.Response.
//
// If a status code has not be set, http.StatusOK is returned. The ContentLength of the response will be -1 if no calls
// have been made to Write(), or set to the length of the buffer. The Close member of the response is always true.
//
// Once this method is called, no other methods should be invoked on the instance.
func (r *ResponseWriter) GetResponse() *http.Response {
	contentLen := -1
	if r.buf != nil {
		contentLen = r.buf.Len()
	}

	if r.statusCode == 0 {
		r.statusCode = http.StatusOK
	}

	return &http.Response{
		StatusCode:    r.statusCode,
		Status:        http.StatusText(r.statusCode),
		Header:        r.header,
		Body:          ioutil.NopCloser(r.buf),
		Close:         true,
		ContentLength: int64(contentLen),
	}
}
