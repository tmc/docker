package main

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"testing"
)

func TestBuildSixtySteps(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestBuildSixtySteps")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "foobuildsixtysteps", ".")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("foobuildsixtysteps")

	logDone("build - build an image with sixty build steps")
}

func TestAddSingleFileToRoot(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "SingleFileToRoot")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add single file to root")
}

func TestAddSingleFileToExistDir(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "SingleFileToExistDir")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add single file to existing dir")
}

func TestAddSingleFileToNonExistDir(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "SingleFileToNonExistDir")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add single file to non-existing dir")
}

func TestAddDirContentToRoot(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "DirContentToRoot")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add directory contents to root")
}

func TestAddDirContentToExistDir(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "DirContentToExistDir")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add directory contents to existing dir")
}

func TestAddWholeDirToRoot(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestAdd")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testaddimg", "WholeDirToRoot")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testaddimg")

	logDone("build - add whole directory to root")
}

func TestDockerignoreWithExplicitAddOnPath(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestDockerignore/IgnorePathsCannotBeAdded")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testdockerignoreimg", ".")
	buildCmd.Dir = buildDirectory
	_, exitCode, err := runCommandWithOutput(buildCmd)

	if err == nil && exitCode == 0 {
		t.Fatal("this Dockerfile is invalid, adding an ignored file is not allowed")
	}

	deleteImages("testdockerignoreimg")

	logDone("build - add a path specified in .dockerignore will not build")
}

func TestDockerignoreWithMultiplePaths(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestDockerignore/IgnoreWorksForMultiplePaths")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testdockerignoreimg", ".")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testdockerignoreimg")

	logDone("build - add whole directory will filter out paths in .dockerignore")
}

func TestDockerignoreCanIgnoreDockerfile(t *testing.T) {
	buildDirectory := filepath.Join(workingDirectory, "build_tests", "TestDockerignore/IgnoreWorksForDockerfile")
	buildCmd := exec.Command(dockerBinary, "build", "-t", "testdockerignoreimg", ".")
	buildCmd.Dir = buildDirectory
	out, exitCode, err := runCommandWithOutput(buildCmd)
	errorOut(err, t, fmt.Sprintf("build failed to complete: %v %v", out, err))

	if err != nil || exitCode != 0 {
		t.Fatal("failed to build the image")
	}

	deleteImages("testdockerignoreimg")

	logDone("build - dockerignore works on Dockerfile too")
}

// TODO: TestCaching

// TODO: TestADDCacheInvalidation
