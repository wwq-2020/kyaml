package main

import (
	"github.com/wwq-2020/kyaml/container"
	"github.com/wwq-2020/kyaml/pod"
	"github.com/wwq-2020/kyaml/podspec"
	"github.com/wwq-2020/kyaml/resources"
)

func main() {
	resource := resources.New().
		WithCPULimits("200Mi").
		WithMemoryRequests("100Mi")

	container := container.New().
		WithName("memory-demo-ctr").
		WithImage("polinux/stress").
		WithResources(resource).
		WithCommand("stress").
		WithArgs([]string{`--vm`, "1", "--vm-bytes", "150M", "--vm-hang", "1"}...)

	podSpec := podspec.New().
		WithContainers(container)

	pod := pod.New().
		WithAPIVersion("v1").
		WithMetadata("name", "memory-demo").
		WithMetadata("namespace", "mem-example").
		WithSpec(podSpec)

	if err := pod.ToYaml("pod.gen.yaml"); err != nil {
		panic(err)
	}
}
