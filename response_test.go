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
