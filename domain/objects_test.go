package domain

import (
	"bufio"
	"errors"
	"fmt"
	"sort"
	"strings"
	"testing"
)

func ExampleBySizeDesc() {
	object := []GitObject{
		{"87a6asmall", 100, 500, "./"},
		{"87a6axxl", 31500, 500, "./"},
		{"87a6abig", 2100, 500, "./"},
		{"87a6amedium", 800, 500, "./"},
	}

	sort.Sort(BySizeDesc(object))
	fmt.Println(object)
	// Output:
	// [{87a6axxl 31500 500 ./} {87a6abig 2100 500 ./} {87a6amedium 800 500 ./} {87a6asmall 100 500 ./}]
}

func TestGetObjectsFailsIfPathIsNotAGitRepo(t *testing.T) {
	mockGit := &MockGitRepoManager{}

	mockGit.EnsureRepoPathFn = func(path string) (string, error) {
		return "", errors.New("error path")
	}

	gom := GitObjectManager{
		Git: mockGit,
	}

	_, err := gom.Get("./path/to/repo")

	if err == nil || err.Error() != "error path" {
		t.Fatal("Not error throw")
	}
}

func TestGetObjectsFailsIfVerifyPackFails(t *testing.T) {
	mockGit := &MockGitRepoManager{}

	mockGit.EnsureRepoPathFn = func(path string) (string, error) {
		return "/path/test", nil
	}

	mockGit.VerifyPackFn = func(path string) (*bufio.Scanner, error) {
		return nil, errors.New("error VerifyPack")
	}

	mockGit.RevListFn = func(path string) (*bufio.Scanner, error) {
		return bufio.NewScanner(strings.NewReader("")), nil
	}

	gom := GitObjectManager{
		Git: mockGit,
	}

	_, err := gom.Get("./path/to/repo")

	if err == nil || err.Error() != "error VerifyPack" {
		t.Fatal("Not error throw")
	}
}

func TestGetObjectsFailsIfRevListFails(t *testing.T) {
	mockGit := &MockGitRepoManager{}

	mockGit.EnsureRepoPathFn = func(path string) (string, error) {
		return "/path/test", nil
	}

	mockGit.VerifyPackFn = func(path string) (*bufio.Scanner, error) {
		return bufio.NewScanner(strings.NewReader("")), nil
	}

	mockGit.RevListFn = func(path string) (*bufio.Scanner, error) {
		return nil, errors.New("error RevList")
	}

	gom := GitObjectManager{
		Git: mockGit,
	}

	_, err := gom.Get("./path/to/repo")

	if err == nil || err.Error() != "error RevList" {
		t.Fatal("Not error throw")
	}
}

func TestGetObjectsFailsIfVerifyPackDoNotHaveProperSizeType(t *testing.T) {
	mockGit := &MockGitRepoManager{}

	mockGit.EnsureRepoPathFn = func(path string) (string, error) {
		return "/path/test", nil
	}

	mockGit.VerifyPackFn = func(path string) (*bufio.Scanner, error) {
		const input = `931179a7cb9827b662968f36f0732dceaae9cc8e blob   3530 3544 534211
0623de54c9501be250cd0907bd09f6b452c53fe6 blob   a 4067 537755
b9018b11dd21302f850984fe9b49e49828101a00 blob   4770 4751 541822
cb87bc696798d4636453465657e67546009890dd blob   3244 2342 234234`

		return bufio.NewScanner(strings.NewReader(input)), nil
	}

	mockGit.RevListFn = func(path string) (*bufio.Scanner, error) {
		const input = `0623de54c9501be250cd0907bd09f6b452c53fe6 app/AppKernel.php
931179a7cb9827b662968f36f0732dceaae9cc8e app/config
b9018b11dd21302f850984fe9b49e49828101a00 app/config/config.yml
ab6785ab65a75ba875b78a65baa758765c6578c5 file`

		return bufio.NewScanner(strings.NewReader(input)), nil
	}

	gom := GitObjectManager{
		Git: mockGit,
	}

	_, err := gom.Get("./path/to/repo")

	if err == nil || err.Error() != "strconv.ParseUint: parsing \"a\": invalid syntax" {
		t.Fatal("Not proper error thrown")
	}
}

func TestGetObjectsFailsIfVerifyPackDoNotHaveProperCompressedSizeType(t *testing.T) {
	mockGit := &MockGitRepoManager{}

	mockGit.EnsureRepoPathFn = func(path string) (string, error) {
		return "/path/test", nil
	}

	mockGit.VerifyPackFn = func(path string) (*bufio.Scanner, error) {
		const input = `931179a7cb9827b662968f36f0732dceaae9cc8e blob   3530 3544 534211
0623de54c9501be250cd0907bd09f6b452c53fe6 blob   5643 b 537755
b9018b11dd21302f850984fe9b49e49828101a00 blob   4770 4751 541822
cb87bc696798d4636453465657e67546009890dd blob   3244 2342 234234`

		return bufio.NewScanner(strings.NewReader(input)), nil
	}

	mockGit.RevListFn = func(path string) (*bufio.Scanner, error) {
		const input = `0623de54c9501be250cd0907bd09f6b452c53fe6 app/AppKernel.php
931179a7cb9827b662968f36f0732dceaae9cc8e app/config
b9018b11dd21302f850984fe9b49e49828101a00 app/config/config.yml
ab6785ab65a75ba875b78a65baa758765c6578c5 file`

		return bufio.NewScanner(strings.NewReader(input)), nil
	}

	gom := GitObjectManager{
		Git: mockGit,
	}

	_, err := gom.Get("./path/to/repo")

	if err == nil || err.Error() != "strconv.ParseUint: parsing \"b\": invalid syntax" {
		t.Fatal("Not proper error thrown")
	}
}

func ExampleGitObjectManager() {
	mockGit := &MockGitRepoManager{}

	mockGit.EnsureRepoPathFn = func(path string) (string, error) {
		return "/path/test", nil
	}

	mockGit.VerifyPackFn = func(path string) (*bufio.Scanner, error) {
		const input = `931179a7cb9827b662968f36f0732dceaae9cc8e blob   3530 3544 534211
0623de54c9501be250cd0907bd09f6b452c53fe6 blob   4053 4067 537755
b9018b11dd21302f850984fe9b49e49828101a00 blob   4770 4751 541822
cb87bc696798d4636453465657e67546009890dd blob   3244 2342 234234
76854567ed47567645a7906a0987ac09879c0978 blob   503032 234234 234235`

		return bufio.NewScanner(strings.NewReader(input)), nil
	}

	mockGit.RevListFn = func(path string) (*bufio.Scanner, error) {
		const input = `0623de54c9501be250cd0907bd09f6b452c53fe6 app/AppKernel.php
76854567ed47567645a7906a0987ac09879c0978 app/AppKernel.php
931179a7cb9827b662968f36f0732dceaae9cc8e app/config
b9018b11dd21302f850984fe9b49e49828101a00 app/config/config.yml
ab6785ab65a75ba875b78a65baa758765c6578c5 file`

		return bufio.NewScanner(strings.NewReader(input)), nil
	}

	gom := GitObjectManager{
		Git: mockGit,
	}

	objects, err := gom.Get("./path/to/repo")

	uniqueObjects := gom.GroupObjectsByFile(objects)

	fmt.Println(err)
	fmt.Println(objects)
	fmt.Println(uniqueObjects)
	// Output:
	// <nil>
	// [{76854567ed47567645a7906a0987ac09879c0978 503032 234234 app/AppKernel.php} {b9018b11dd21302f850984fe9b49e49828101a00 4770 4751 app/config/config.yml} {0623de54c9501be250cd0907bd09f6b452c53fe6 4053 4067 app/AppKernel.php} {931179a7cb9827b662968f36f0732dceaae9cc8e 3530 3544 app/config}]
	// [{ 507085 234234 app/AppKernel.php} { 4770 4751 app/config/config.yml} { 3530 3544 app/config}]
}
