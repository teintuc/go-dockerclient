// Copyright 2019 go-dockerclient authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build docker_integration,!windows

package docker

func integrationCreateContainerOpts(imageName string, hostConfig *HostConfig) CreateContainerOptions {
	return CreateContainerOptions{
		Config: &Config{
			Image: imageName,
			Cmd:   []string{"cat", "/home/gopher/file.txt"},
			User:  "gopher",
		},
		HostConfig: hostConfig,
	}
}
