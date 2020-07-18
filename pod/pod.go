package pod

import (
	"io/ioutil"

	"github.com/wwq-2020/kyaml/podspec"
	"gopkg.in/yaml.v2"
)

// Pod Pod
type Pod struct {
	APIVersion string            `yaml:"apiVersion,omitempty"`
	Metadata   map[string]string `yaml:"metadata,omitempty"`
	Spec       *podspec.PodSpec  `yaml:"spec,omitempty"`
}

// New New
func New() *Pod {
	return &Pod{
		Metadata: make(map[string]string),
	}
}

// WithAPIVersion WithAPIVersion
func (p *Pod) WithAPIVersion(apiVersion string) *Pod {
	p.APIVersion = apiVersion
	return p
}

// WithMetadata WithMetadata
func (p *Pod) WithMetadata(key, value string) *Pod {
	p.Metadata[key] = value
	return p
}

// WithSpec WithSpec
func (p *Pod) WithSpec(spec *podspec.PodSpec) *Pod {
	p.Spec = spec
	return p
}

// ToYaml ToYaml
func (p *Pod) ToYaml(filename string) error {
	data, err := yaml.Marshal(p)
	if err != nil {
		return err
	}
	return ioutil.WriteFile(filename, data, 0644)
}
