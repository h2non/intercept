package intercept

import (
	"encoding/json"
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

func TestResponseModifierReadStringError(t *testing.T) {
	req := &http.Request{}
	body := ioutil.NopCloser(&errorReader{})
	resp := &http.Response{Body: body}
	modifier := NewResponseModifier(req, resp)
	str, err := modifier.ReadString()
	st.Expect(t, err, errRead)
	st.Expect(t, str, "")
}

func TestResponseModifierReadBytes(t *testing.T) {
	req := &http.Request{}
	bodyStr := `{"name":"Rick"}`
	strReader := strings.NewReader(bodyStr)
	body := ioutil.NopCloser(strReader)
	resp := &http.Response{Body: body}
	modifier := NewResponseModifier(req, resp)
	bytes, err := modifier.ReadBytes()
	st.Expect(t, err, nil)
	st.Expect(t, string(bytes), bodyStr)
}

func TestResponseModifierReadBytesError(t *testing.T) {
	req := &http.Request{}
	body := ioutil.NopCloser(&errorReader{})
	resp := &http.Response{Body: body}
	modifier := NewResponseModifier(req, resp)
	bytes, err := modifier.ReadBytes()
	st.Expect(t, err, errRead)
	st.Expect(t, string(bytes), "")
}

func TestResponseModifierDecodeJSON(t *testing.T) {
	req := &http.Request{}
	bodyStr := `{"name":"Rick"}`
	strReader := strings.NewReader(bodyStr)
	body := ioutil.NopCloser(strReader)
	resp := &http.Response{Body: body}
	modifier := NewResponseModifier(req, resp)
	u := user{}
	err := modifier.DecodeJSON(&u)
	st.Expect(t, err, nil)
	st.Expect(t, u.Name, "Rick")
}

func TestResponseModifierDecodeJSONErrorFromReadBytes(t *testing.T) {
	req := &http.Request{}
	body := ioutil.NopCloser(&errorReader{})
	resp := &http.Response{Body: body}
	modifier := NewResponseModifier(req, resp)
	u := user{}
	err := modifier.DecodeJSON(&u)
	st.Expect(t, err, errRead)
	st.Expect(t, u.Name, "")
}

func TestResponseModifierDecodeJSONErrorFromEOF(t *testing.T) {
	req := &http.Request{}
	bodyStr := ""
	strReader := strings.NewReader(bodyStr)
	body := ioutil.NopCloser(strReader)
	resp := &http.Response{Body: body}
	modifier := NewResponseModifier(req, resp)
	u := user{}
	err := modifier.DecodeJSON(&u)
	st.Expect(t, err, nil)
	st.Expect(t, u.Name, "")
}

func TestResponseModifierDecodeJSONErrorFromDecode(t *testing.T) {
	req := &http.Request{}
	bodyStr := "/"
	strReader := strings.NewReader(bodyStr)
	body := ioutil.NopCloser(strReader)
	resp := &http.Response{Body: body}
	modifier := NewResponseModifier(req, resp)
	u := user{}
	err := modifier.DecodeJSON(&u)
	_, ok := (err).(*json.SyntaxError)
	st.Expect(t, ok, true)
	st.Expect(t, err.Error(), "invalid character '/' looking for beginning of value")
	st.Expect(t, u.Name, "")
}
