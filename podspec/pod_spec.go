package podspec

import (
	"github.com/wwq-2020/kyaml/container"
)

// PodSpec PodSpec
type PodSpec struct {
	Containers []*container.Container `yaml:"containers,omitempty"`
}

// New New
func New() *PodSpec {
	return &PodSpec{}
}

// WithContainers WithContainers
func (ps *PodSpec) WithContainers(containers ...*container.Container) *PodSpec {
	ps.Containers = append(ps.Containers, containers...)
	return ps
}
