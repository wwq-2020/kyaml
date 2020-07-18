package resources

import "sync"

// Resources Resources
type Resources struct {
	Limits       *Limits   `yaml:"limits,omitempty"`
	Requests     *Requests `yaml:"requests,omitempty"`
	limitsOnce   sync.Once
	requestsOnce sync.Once
}

// Limits Limits
type Limits struct {
	Memory string `yaml:"memory,omitempty"`
	CPU    string `yaml:"cpu,omitempty"`
}

// Requests Requests
type Requests struct {
	Memory string `yaml:"memory,omitempty"`
	CPU    string `yaml:"cpu,omitempty"`
}

// New New
func New() *Resources {
	return &Resources{}
}

// WithMemoryLimits WithMemoryLimits
func (r *Resources) WithMemoryLimits(value string) *Resources {
	r.limitsOnce.Do(func() {
		r.Limits = &Limits{}
	})
	r.Limits.Memory = value
	return r
}

// WithCPULimits WithCPULimits
func (r *Resources) WithCPULimits(value string) *Resources {
	r.limitsOnce.Do(func() {
		r.Limits = &Limits{}
	})
	r.Limits.CPU = value
	return r
}

// WithMemoryRequests WithMemoryRequests
func (r *Resources) WithMemoryRequests(value string) *Resources {
	r.requestsOnce.Do(func() {
		r.Requests = &Requests{}
	})
	r.Requests.Memory = value
	return r
}

// WithCPURequests WithCPURequests
func (r *Resources) WithCPURequests(value string) *Resources {
	r.requestsOnce.Do(func() {
		r.Requests = &Requests{}
	})
	r.Requests.CPU = value
	return r
}
