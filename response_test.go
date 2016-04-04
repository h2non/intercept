package intercept

import (
	"github.com/nbio/st"
	"net/http"
	"testing"
)

func TestNewResponseModifier(t *testing.T) {
	header := http.Header{}
	req := &http.Request{}
	resp := &http.Response{Header: header}
	modifier := NewResponseModifier(req, resp)
	st.Expect(t, modifier.Request, req)
	st.Expect(t, modifier.Response, resp)
	st.Expect(t, modifier.Header, header)
}

func TestStatus(t *testing.T) {
	req := &http.Request{}
	resp := &http.Response{}
	modifier := NewResponseModifier(req, resp)
	modifier.Status(404)
	st.Expect(t, resp.StatusCode, 404)
	st.Expect(t, resp.Status, "404 Not Found")
}
