# Libkermit [![Build Status](https://travis-ci.org/vdemeester/libkermit.svg?branch=add-readme)](https://travis-ci.org/vdemeester/libkermit)

> When green is all there is to be<br/>
> It could make you wonder why<br/>
> But why wonder why wonder<br/>
> I am green, and it'll do fine<br/>
> It's beautiful,<br/>
> and I think it's what I want to be.<br/>
> -- Kermit the Frog

When [Docker](https://github.com/docker/docker) meets with
integration/acceptance tests to make you see everything in
green. **Libkermit** is a Go(lang) library that aims to ease the
writing of integration tests (any non unit tests actually) with the
helps of Docker and it's ecosystem (mainly
[libcompose](https://github.com/docker/libcompose)).

The goals are :

- Easy docker manipulation, from managing a simple container to boot
  up a whole stack.
    - create, delete, pause, … containers
    - check for a certain state containers (inspect them)
    - support *compose files* to allow starting a whole stack
- Testing suite and functions, in a simple fashion.
- Works seamlessly with the Go(lang) `testing` framework.
- Try to not force any testing framework but also tries to integrate
  with them ([go-check](https://github.com/go-check/check),
  [testify](https://github.com/stretchr/testify), …).

**Note: This is experimental and not even implemented yet. You are on your own right now**


## Package `docker`

This package holds functions and structs to ease docker uses.

```go
package yours

import (
    "testing"

    "github.com/vdemeester/libkermit/docker"
)

func TestItMyFriend(t *testing.T) {
    container, err := docker.Start("vdemeester/myawesomeimage")
    if err != nil {
        t.Fatal(err)
    }

    // Do your stuff

    err = docker.Stop(container.ID)
    if err != nil {
        t.Fatal(err)
    }
}
```



## Package `compose`

This package holds functions and structs to ease docker uses.

```go
package yours

import (
    "testing"

    "github.com/vdemeester/libkermit/compose"
)

func TestItMyFriend(t *testing.T) {
    project, err := compose.CreateProject("simple", "./assets/simple.yml")
    if err != nil {
        t.Fatal(err)
    }
    err = project.Start()
	if err != nil {
		t.Fatal(err)
	}

    // Do your stuff

    err = project.Stop()
	if err != nil {
		t.Fatal(err)
	}
}
```


## Other packages to come

- `suite` : functions and structs to setup tests suites.


