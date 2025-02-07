/*
   Copyright The containerd Authors.

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package logging

import (
	"context"
	"io"
)

// Config of the container logs
type Config struct {
	Stdout    io.Reader
	Stderr    io.Reader
	ID        string
	Namespace string
}

// LoggerFunc is implemented by custom v2 logging binaries.
//
// ready should be called when the logging binary finishes its setup and the container can be started.
//
// An example implementation of LoggerFunc: https://github.com/containerd/containerd/tree/main/runtime/v2#logging
type LoggerFunc func(ctx context.Context, cfg *Config, ready func() error) error
