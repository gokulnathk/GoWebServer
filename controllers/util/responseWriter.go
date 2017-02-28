package util

import (
	"compress/gzip"
	"net/http"
	"strings"
)

//ClosableResponseWriter interface is used to create a custom gzip writer
type ClosableResponseWriter interface {
	http.ResponseWriter
	Close()
}

type gzipResponseWriter struct {
	http.ResponseWriter
	*gzip.Writer
}

func (gzWriter gzipResponseWriter) Write(data []byte) (int, error) {
	return gzWriter.Writer.Write(data)
}

func (gzWriter gzipResponseWriter) Close() {
	gzWriter.Writer.Close()
}

func (gzWriter gzipResponseWriter) Header() http.Header {
	return gzWriter.ResponseWriter.Header()
}

type closableResponseWriter struct {
	http.ResponseWriter
}

func (closableWriter closableResponseWriter) Close() {

}

//GetResponseWriter function used to get the appropriate response writer for the client
func GetResponseWriter(w http.ResponseWriter, r *http.Request) ClosableResponseWriter {

	var resWriter ClosableResponseWriter

	if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
		w.Header().Set("Content-Encoding", "gzip")
		gW := gzipResponseWriter{
			ResponseWriter: w,
			Writer:         gzip.NewWriter(w),
		}
		resWriter = gW
	} else {
		resWriter = closableResponseWriter{ResponseWriter: w}
	}
	return resWriter
}
