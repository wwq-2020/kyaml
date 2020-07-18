package container

import (
	"github.com/wwq-2020/kyaml/resources"
)

// Container Container
type Container struct {
	Name      string               `yaml:"name,omitempty"`
	Image     string               `yaml:"image,omitempty"`
	Resources *resources.Resources `yaml:"resources,omitempty"`
	Command   string               `yaml:"command,omitempty"`
	Args      []string             `yaml:"args,omitempty,flow"`
}

// New New
func New() *Container {
	return &Container{}
}

// WithName WithName
func (c *Container) WithName(name string) *Container {
	c.Name = name
	return c
}

// WithImage WithImage
func (c *Container) WithImage(image string) *Container {
	c.Image = image
	return c
}

// WithResources WithResources
func (c *Container) WithResources(resources *resources.Resources) *Container {
	c.Resources = resources
	return c
}

// WithCommand WithCommand
func (c *Container) WithCommand(command string) *Container {
	c.Command = command
	return c
}

// WithArgs WithArgs
func (c *Container) WithArgs(args ...string) *Container {
	c.Args = append(c.Args, args...)
	return c
}
