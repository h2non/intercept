package intercept

import (
	"github.com/nbio/st"
	"io/ioutil"
	"net/http"
	"strings"
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

func TestResponseModifierReadString(t *testing.T) {
	req := &http.Request{}
	bodyStr := `{"name":"Rick"}`
	strReader := strings.NewReader(bodyStr)
	body := ioutil.NopCloser(strReader)
	resp := &http.Response{Body: body}
	modifier := NewResponseModifier(req, resp)
	str, err := modifier.ReadString()
	st.Expect(t, err, nil)
	st.Expect(t, str, bodyStr)
}
