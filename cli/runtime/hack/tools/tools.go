//go:build tools
// +build tools

// Copyright 2021 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

// Package tools imports things required by build scripts, to force `go mod` to see them as dependencies

package tools

import (
	_ "golang.org/x/tools/cmd/goimports"
	_ "sigs.k8s.io/controller-tools/cmd/controller-gen"
)
