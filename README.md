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
    project, err := docker.NewProjectFromEnv()
    if err != nil {
        t.Fatal(err)
    }
    container, err := p.Start("vdemeester/myawesomeimage")
    if err != nil {
        t.Fatal(err)
    }

    // Do your stuff
    // […]

    // Clean the containers managed by libkermit
    err = docker.Clean()
    if err != nil {
        t.Fatal(err)
    }
}
```

### Package `docker/testing`

This package map the `docker` package but takes a `*testing.T` struct
on all methods. The idea is to write even less. Let's write the same
example as above.


```go
package yours

import (
    "testing"

    docker "github.com/vdemeester/libkermit/docker/testing"
)

func TestItMyFriend(t *testing.T) {
    project := docker.NewProjectFromEnv(t)
    container := p.Start(t, "vdemeester/myawesomeimage")

    // Do your stuff
    // […]

    // Clean the containers managed by libkermit
    docker.Clean(t)
}
```

### Package `docker/check`

This package map the `docker` package but takes a `*check.C` struct
on all methods. The idea is to write even less. Let's write the same
example as above.


```go
package yours

import (
    "testing"

	"github.com/go-check/check"
    docker "github.com/vdemeester/libkermit/docker/check"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { check.TestingT(t) }

type CheckSuite struct{}

var _ = check.Suite(&CheckSuite{})

func (s *CheckSuite) TestItMyFriend(c *check.C) {
    project := docker.NewProjectFromEnv(c)
    container := p.Start(c, "vdemeester/myawesomeimage")

    // Do your stuff
    // […]

    // Clean the containers managed by libkermit
    docker.Clean(c)
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

### Package `compose/testing`

This package map the `compose` package but takes a `*testing.T` struct
on all methods. The idea is to write even less. Let's write the same
example as above.


```go
package yours

import (
    "testing"

    docker "github.com/vdemeester/libkermit/compose/testing"
)

func TestItMyFriend(t *testing.T) {
    project := compose.CreateProject(t, "simple", "./assets/simple.yml")
    project.Start(t)

    // Do your stuff

    project.Stop(t)
}
```

### Package `compose/check`

This package map the `compose` package but takes a `*check.C` struct
on all methods. The idea is to write even less. Let's write the same
example as above.


```go
package yours

import (
    "testing"

	"github.com/go-check/check"
    docker "github.com/vdemeester/libkermit/compose/check"
)

// Hook up gocheck into the "go test" runner
func Test(t *testing.T) { check.TestingT(t) }

type CheckSuite struct{}

var _ = check.Suite(&CheckSuite{})

func (s *CheckSuite) TestItMyFriend(c *check.C) {
    project := compose.CreateProject(c, "simple", "./assets/simple.yml")
    project.Start(c)

    // Do your stuff

    project.Stop(c)
}
```


## Other packages to come

- `suite` : functions and structs to setup tests suites.


