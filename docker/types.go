package docker

// ContainerConfig holds container libkermit configuration possibilities
type ContainerConfig struct {
	Cmd        []string
	Entrypoint []string
	Labels     map[string]string
}
