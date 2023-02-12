package hlog

const (
	DefaultHeader = "X-Request-Id"
)

var global = &Config{RequestIdHeader: DefaultHeader}

type Config struct {
	RequestIdHeader string
}
