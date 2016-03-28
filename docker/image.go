package docker

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/net/context"

	// FIXME(vdemeester) replace this with docker/distribution reference package
	"github.com/docker/distribution/reference"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
)

// Pull pulls the given reference (image)
func (p *Project) Pull(ref string) error {
	return p.ensureImageExists(ref, true)
}

func (p *Project) ensureImageExists(ref string, force bool) error {
	distributionRef, err := reference.ParseNamed(ref)
	if err != nil {
		return err
	}

	var tag string
	switch x := distributionRef.(type) {
	case reference.Canonical:
		tag = x.Digest().String()
	case reference.NamedTagged:
		tag = x.Tag()
	default:
		// FIXME(vdemeester) this will go at some point :D
		tag = "latest"
	}

	if !force {
		// Check if ref is already there
		_, _, err = p.Client.ImageInspectWithRaw(context.Background(), ref, false)
		if err != nil && !client.IsErrImageNotFound(err) {
			return err
		}
		if err == nil {
			return nil
		}
	}

	// And pull it
	options := types.ImagePullOptions{
		ImageID: distributionRef.Name(),
		Tag:     tag,
	}
	responseBody, err := p.Client.ImagePull(context.Background(), options, nil)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	defer responseBody.Close()

	_, err = ioutil.ReadAll(responseBody)
	return err
}
