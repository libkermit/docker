package docker

import (
	"fmt"
)

const (
	// KermitLabel is the label used to add to containers
	KermitLabel = "com.github.vdemeester.libkermit"
)

var (
	// KermitHeader defines default kermit headers to pass through the API
	KermitHeader = map[string]string{
		"User-Agent": "libkermit-1.0",
	}

	// KermitLabels defines default docker labels kermit will but on
	// containers and such.
	KermitLabels = map[string]string{
		KermitLabel: "true",
	}

	// KermitLabelFilter is the filter to use to find containers managed by kermit
	KermitLabelFilter = fmt.Sprintf("label=%s", KermitLabel)

	// DefaultStopTimeout is the default timeout for the stop command
	DefaultStopTimeout = 10
)

//Remove(ContainerID string) error
//WaitInspect(containerID, expression, value string) error
