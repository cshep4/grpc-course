//go:build tools
// +build tools

// This file ensures that the mockgen tool is vendored so we always generate mocks with a consistent version.
package tools

import _ "github.com/golang/mock/mockgen"
