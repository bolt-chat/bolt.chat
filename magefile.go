// +build mage

package main

import (
	"fmt"
	"path"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

/*
	CGO_ENABLED=0 GOOS=linux \
		go build -ldflags "-s" -a -installsuffix cgo \
		-o build/server \
		$(SERVER_ENTRY)
*/

const name string = "boltchat"
const buildDir string = "build"

const serverPrefix string = "server"
const serverEntry string = "cmd/server/server.go"

type Build mg.Namespace
type Docker mg.Namespace

type BuildOptions struct {
	Static    bool
	Extension string
}

func build(os string, arch string, entry string, opts BuildOptions) error {
	env := map[string]string{
		"GOOS":   os,
		"GOARCH": arch,
	}

	// Build static binary
	if opts.Static {
		env["CGO_ENABLED"] = "0"
	}

	outputName := fmt.Sprintf(
		"%s-%s-%s-%s", name, serverPrefix, os, arch,
	)

	outputPath := path.Join(
		buildDir,
		outputName,
	)

	if opts.Extension != "" {
		outputPath += fmt.Sprintf(".%s", opts.Extension)
	}

	args := []string{
		"build",
		"-o",
		outputPath,
		// "-ldflags='-s -w'", // TODO: this throws an error
		entry,
	}

	fmt.Println(args)

	return sh.RunWith(
		env, "go", args...,
	)
}

/*
Build
*/

// Builds all binaries
func (Build) All() {
	mg.Deps(
		Build.ServerDarwinAmd64,
		Build.ServerLinuxAmd64,
		Build.ServerWindowsAmd64,
	)
}

// Builds the server binary for Linux (amd64)
func (Build) ServerLinuxAmd64() error {
	return build("linux", "amd64", serverEntry, BuildOptions{})
}

// Builds the server binary for Windows (amd64)
func (Build) ServerWindowsAmd64() error {
	return build("windows", "amd64", serverEntry, BuildOptions{
		Extension: "exe",
	})
}

// Builds the server binary for Darwin/macOS (amd64)
func (Build) ServerDarwinAmd64() error {
	return build("darwin", "amd64", serverEntry, BuildOptions{})
}

// Builds the server binary for Darwin/macOS (arm64, M1)
// func (Build) ServerDarwinArm64() error {
// 	return build("darwin", "arm64", serverEntry, false)
// }

// Builds the server binary for use in a Docker container
func (Build) ServerContainer() error {
	return build("linux", "amd64", serverEntry, BuildOptions{
		Static: true,
	})
}

/*
Docker
*/

// Builds a Docker image for the server
func (Docker) Build() error {
	return sh.RunV("docker", "build", ".", "-t", name)
}

/*
Misc
*/

// Cleans up build directories
func Clean() {
	sh.Rm("build")
}