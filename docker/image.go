package docker

import (
	"fmt"
	"io/ioutil"

	"golang.org/x/net/context"

	"github.com/docker/docker/reference"
	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
)

func ensureImageExists(cli client.APIClient, image string) error {
	distributionRef, err := reference.ParseNamed(image)
	if err != nil {
		return err
	}
	if reference.IsNameOnly(distributionRef) {
		distributionRef = reference.WithDefaultTag(distributionRef)
	}
	var tag string
	switch x := distributionRef.(type) {
	case reference.Canonical:
		tag = x.Digest().String()
	case reference.NamedTagged:
		tag = x.Tag()
	}

	// Check if image is already there
	_, _, err = cli.ImageInspectWithRaw(context.Background(), image, false)
	if err != nil && !client.IsErrImageNotFound(err) {
		return err
	}
	if err == nil {
		return nil
	}

	// And pull it
	options := types.ImagePullOptions{
		ImageID: distributionRef.Name(),
		Tag:     tag,
	}
	responseBody, err := cli.ImagePull(context.Background(), options, nil)
	if err != nil {
		fmt.Printf("%v", err)
		return err
	}
	defer responseBody.Close()

	_, err = ioutil.ReadAll(responseBody)
	return err
}
