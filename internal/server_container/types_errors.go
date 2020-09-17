package server_container

import (
	"fmt"
	"strings"
)

//UnreachableContainerError is the error sent back if the container was started
//But could not be reached.
type UnreachableContainerError struct {
	name string
}

func (uc *UnreachableContainerError) Error() string {
	var builder strings.Builder
	builder.WriteString(
		fmt.Sprintf(
			"%s was created but never successfully responsed to a POST request. ",
			uc.name,
		),
	)
	builder.WriteString("The server IP address was sent back anyway.")
	return builder.String()
}
