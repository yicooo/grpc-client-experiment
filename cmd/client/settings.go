package main

import "encoding/json"

const (
	roundRobin = "round_robin"
)

type serviceConfig struct {
	LoadBalancingConfig []map[string]interface{} `json:"loadBalancingConfig"`
}

func NewServiceConfig() *serviceConfig {
	return &serviceConfig{}
}

// ToString returns state of serviceConfig as string.
// WithDefaultServiceConfig method accepts json string.
func (s *serviceConfig) ToString() string {
	ss, err := json.Marshal(s)
	if err != nil {
		return ""
	}
	return string(ss)
}

// EnableRoundRobin adds required key to serviceConfig to enable
// round-robin load balancing option for GRPC client.
func (s *serviceConfig) EnableRoundRobin() *serviceConfig {
	s.LoadBalancingConfig = append(s.LoadBalancingConfig,
		map[string]interface{}{
			// Interesting design choice by Google developers.
			roundRobin: struct{}{},
		},
	)

	return s
}
