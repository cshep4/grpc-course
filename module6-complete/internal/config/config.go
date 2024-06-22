package config

type RetryPolicy struct {
	MaxAttempts          int      `json:"maxAttempts"`
	InitialBackoff       string   `json:"initialBackoff"`
	MaxBackoff           string   `json:"maxBackoff"`
	BackoffMultiplier    float64  `json:"backoffMultiplier"`
	RetryableStatusCodes []string `json:"retryableStatusCodes"`
}

type NameConfig struct {
	Service string `json:"service,omitempty"`
	Method  string `json:"method,omitempty"`
}

type MethodConfig struct {
	Name        []NameConfig `json:"name,omitempty"`
	RetryPolicy *RetryPolicy `json:"retryPolicy,omitempty"`
	Timeout     string       `json:"timeout,omitempty"`
}

type Config struct {
	LoadBalancingPolicy string         `json:"loadBalancingPolicy,omitempty"`
	MethodConfig        []MethodConfig `json:"methodConfig,omitempty"`
}
