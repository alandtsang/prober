package prober

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/labstack/gommon/log"

	"github.com/alandtsang/prober/config"
)

type URIScheme string

const (
	// URISchemeHTTP means that the scheme used will be http://
	URISchemeHTTP URIScheme = "http"
	// URISchemeHTTPS means that the scheme used will be https://
	URISchemeHTTPS URIScheme = "https"
)

// HTTPHeader describes a custom header to be used in HTTP probes
type HTTPHeader struct {
	// The header field name
	Name string
	// The header field value
	Value string
}

type HTTPProbe struct {
	Path        string
	Port        int32
	Host        string
	Scheme      URIScheme
	HTTPHeaders []HTTPHeader
	Conf        *config.Config
}

func ProbeHTTP(target string) bool {
	if !strings.HasPrefix(target, "http://") && !strings.HasPrefix(target, "https://") {
		target = "http://" + target
	}

	targetURL, err := url.Parse(target)
	if err != nil {
		log.Error("Could not parse target URL, err", err)
		return false
	}

	targetHost := targetURL.Hostname()
	targetPort := targetURL.Port()

	fmt.Printf("host:%s, port:%s\n", targetHost, targetPort)

	client := newClient()

	req, err := http.NewRequest("GET", targetURL.String(), nil)
	if err != nil {
		log.Errorf("ProbeHTTP new request failed, %v", err)
		return false
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Errorf("ProbeHTTP client do failed, %v", err)
		return false
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Errorf("ProbeHTTP read body failed, %v", err)
		return false
	}

	fmt.Println("resp body:", string(body))
	return true
}

// newClient returns a http.Client using the specified http.RoundTripper.
func newClient() *http.Client {
	rt := newTransport()
	return &http.Client{Transport: rt}
}

func newTransport() *http.Transport {
	return &http.Transport{}
}
