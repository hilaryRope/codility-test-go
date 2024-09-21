package codility_test_go

import "fmt"

type Url struct {
	scheme      string
	host        string
	port        int
	path        string
	queryParams string
}

func (c *Url) Build() string {
	scheme := c.scheme
	if scheme == "" {
		scheme = "http" // Default to http if no scheme is set
	}

	port := ""
	if c.port != 0 {
		port = fmt.Sprintf(":%d", c.port)
	}
	qp := ""
	if c.queryParams != "" {
		qp = "?" + c.queryParams
	}

	return fmt.Sprintf("%s://%s%s%s%s",
		scheme,
		c.host,
		port,
		c.path,
		qp,
	)
}

type IUrlBuilder interface {
	Https() IUrlBuilder
	Host(string) IUrlBuilder
	Port(int) IUrlBuilder
	Path(string) IUrlBuilder
	QueryParams(map[string]string) IUrlBuilder
	GetUrl() *Url
}

type UrlBuilder struct {
	url *Url
}

func NewUrlBuilder() IUrlBuilder {
	return &UrlBuilder{url: &Url{}}
}

func (b *UrlBuilder) Https() IUrlBuilder {
	b.url.scheme = "https"
	return b
}

func (b *UrlBuilder) Host(host string) IUrlBuilder {
	b.url.host = host
	return b
}

func (b *UrlBuilder) Port(port int) IUrlBuilder {
	b.url.port = port
	return b
}

func (b *UrlBuilder) Path(path string) IUrlBuilder {
	b.url.path = path
	return b
}

func (b *UrlBuilder) QueryParams(params map[string]string) IUrlBuilder {
	queryParams := ""
	for key, value := range params {
		if queryParams != "" {
			queryParams += "&"
		}
		queryParams += fmt.Sprintf("%s=%s", key, value)
	}
	b.url.queryParams = queryParams
	return b
}

func (b *UrlBuilder) GetUrl() *Url {
	return b.url
}