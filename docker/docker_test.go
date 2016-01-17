package docker

import (
	"fmt"
	"io"

	"github.com/docker/engine-api/client"
	"github.com/docker/engine-api/types"
	"github.com/docker/engine-api/types/container"
	"github.com/docker/engine-api/types/filters"
	"github.com/docker/engine-api/types/network"
	"github.com/docker/engine-api/types/registry"
)

type fakeClient struct {
	image         string
	containerID   string
	containerName string
	running       bool
}

func (c *fakeClient) ClientVersion() string {
	return ""
}

func (c *fakeClient) ContainerAttach(options types.ContainerAttachOptions) (types.HijackedResponse, error) {
	return types.HijackedResponse{}, nil
}

func (c *fakeClient) ContainerCommit(options types.ContainerCommitOptions) (types.ContainerCommitResponse, error) {
	return types.ContainerCommitResponse{}, nil
}

func (c *fakeClient) ContainerCreate(config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (types.ContainerCreateResponse, error) {
	if c.image != config.Image {
		return types.ContainerCreateResponse{}, fmt.Errorf("image not found")
	}
	c.containerName = containerName
	return types.ContainerCreateResponse{
		ID: c.containerID,
	}, nil
}

func (c *fakeClient) ContainerInspect(containerID string) (types.ContainerJSON, error) {
	if c.containerID != "" && c.containerID == containerID {
		return types.ContainerJSON{
			ContainerJSONBase: &types.ContainerJSONBase{
				ID:    c.containerID,
				Name:  c.containerName,
				Image: c.image,
				State: &types.ContainerState{
					Running: c.running,
				},
			},
		}, nil
	}
	return types.ContainerJSON{}, fmt.Errorf("container not found: %s vs %s", c.containerID, containerID)
}

func (c *fakeClient) ContainerDiff(containerID string) ([]types.ContainerChange, error) {
	return []types.ContainerChange{}, nil
}

func (c *fakeClient) ContainerExecAttach(execID string, config types.ExecConfig) (types.HijackedResponse, error) {
	return types.HijackedResponse{}, nil
}

func (c *fakeClient) ContainerExecCreate(config types.ExecConfig) (types.ContainerExecCreateResponse, error) {
	return types.ContainerExecCreateResponse{}, nil
}

func (c *fakeClient) ContainerExecInspect(execID string) (types.ContainerExecInspect, error) {
	return types.ContainerExecInspect{}, nil
}

func (c *fakeClient) ContainerExecResize(options types.ResizeOptions) error {
	return nil
}

func (c *fakeClient) ContainerExecStart(execID string, config types.ExecStartCheck) error {
	return nil
}

func (c *fakeClient) ContainerExport(containerID string) (io.ReadCloser, error) {
	return nil, nil
}

func (c *fakeClient) ContainerInspectWithRaw(containerID string, getSize bool) (types.ContainerJSON, []byte, error) {
	return types.ContainerJSON{}, []byte{}, nil
}

func (c *fakeClient) ContainerKill(containerID, signal string) error {
	return nil
}

func (c *fakeClient) ContainerList(options types.ContainerListOptions) ([]types.Container, error) {
	return []types.Container{}, nil
}

func (c *fakeClient) ContainerLogs(options types.ContainerLogsOptions) (io.ReadCloser, error) {
	return nil, nil
}

func (c *fakeClient) ContainerPause(containerID string) error {
	return nil
}

func (c *fakeClient) ContainerRemove(options types.ContainerRemoveOptions) error {
	return nil
}

func (c *fakeClient) ContainerRename(containerID, newContainerName string) error {
	return nil
}

func (c *fakeClient) ContainerResize(options types.ResizeOptions) error {
	return nil
}

func (c *fakeClient) ContainerRestart(containerID string, timeout int) error {
	return nil
}

func (c *fakeClient) ContainerStatPath(containerID, path string) (types.ContainerPathStat, error) {
	return types.ContainerPathStat{}, nil
}

func (c *fakeClient) ContainerStats(containerID string, stream bool) (io.ReadCloser, error) {
	return nil, nil
}

func (c *fakeClient) ContainerStart(containerID string) error {
	if !c.running {
		return fmt.Errorf("could not start the container")
	}
	return nil
}

func (c *fakeClient) ContainerStop(containerID string, timeout int) error {
	if c.containerID != "" && c.containerID == containerID {
		return nil
	}
	return fmt.Errorf("Could not stop the container")
}

func (c *fakeClient) ContainerTop(containerID string, arguments []string) (types.ContainerProcessList, error) {
	return types.ContainerProcessList{}, nil
}

func (c *fakeClient) ContainerUnpause(containerID string) error {
	return nil
}

func (c *fakeClient) ContainerUpdate(containerID string, updateConfig container.UpdateConfig) error {
	return nil
}

func (c *fakeClient) ContainerWait(containerID string) (int, error) {
	return 0, nil
}

func (c *fakeClient) CopyFromContainer(containerID, srcPath string) (io.ReadCloser, types.ContainerPathStat, error) {
	return nil, types.ContainerPathStat{}, nil
}

func (c *fakeClient) CopyToContainer(options types.CopyToContainerOptions) error {
	return nil
}

func (c *fakeClient) Events(options types.EventsOptions) (io.ReadCloser, error) {
	return nil, nil
}

func (c *fakeClient) ImageBuild(options types.ImageBuildOptions) (types.ImageBuildResponse, error) {
	return types.ImageBuildResponse{}, nil
}

func (c *fakeClient) ImageCreate(options types.ImageCreateOptions) (io.ReadCloser, error) {
	return nil, nil
}

func (c *fakeClient) ImageHistory(imageID string) ([]types.ImageHistory, error) {
	return []types.ImageHistory{}, nil
}

func (c *fakeClient) ImageImport(options types.ImageImportOptions) (io.ReadCloser, error) {
	return nil, nil
}

func (c *fakeClient) ImageInspectWithRaw(imageID string, getSize bool) (types.ImageInspect, []byte, error) {
	return types.ImageInspect{}, []byte{}, nil
}

func (c *fakeClient) ImageList(options types.ImageListOptions) ([]types.Image, error) {
	return []types.Image{}, nil
}

func (c *fakeClient) ImageLoad(input io.Reader) (types.ImageLoadResponse, error) {
	return types.ImageLoadResponse{}, nil
}

func (c *fakeClient) ImagePull(options types.ImagePullOptions, privilegeFunc client.RequestPrivilegeFunc) (io.ReadCloser, error) {
	return nil, nil
}

func (c *fakeClient) ImagePush(options types.ImagePushOptions, privilegeFunc client.RequestPrivilegeFunc) (io.ReadCloser, error) {
	return nil, nil
}

func (c *fakeClient) ImageRemove(options types.ImageRemoveOptions) ([]types.ImageDelete, error) {
	return []types.ImageDelete{}, nil
}

func (c *fakeClient) ImageSearch(options types.ImageSearchOptions, privilegeFunc client.RequestPrivilegeFunc) ([]registry.SearchResult, error) {
	return []registry.SearchResult{}, nil
}

func (c *fakeClient) ImageSave(imageIDs []string) (io.ReadCloser, error) {
	return nil, nil
}

func (c *fakeClient) ImageTag(options types.ImageTagOptions) error {
	return nil
}

func (c *fakeClient) Info() (types.Info, error) {
	return types.Info{}, nil
}

func (c *fakeClient) NetworkConnect(networkID, containerID string, config *network.EndpointSettings) error {
	return nil
}

func (c *fakeClient) NetworkCreate(options types.NetworkCreate) (types.NetworkCreateResponse, error) {
	return types.NetworkCreateResponse{}, nil
}

func (c *fakeClient) NetworkDisconnect(networkID, containerID string, force bool) error {
	return nil
}

func (c *fakeClient) NetworkInspect(networkID string) (types.NetworkResource, error) {
	return types.NetworkResource{}, nil
}

func (c *fakeClient) NetworkList(options types.NetworkListOptions) ([]types.NetworkResource, error) {
	return []types.NetworkResource{}, nil
}

func (c *fakeClient) NetworkRemove(networkID string) error {
	return nil
}

func (c *fakeClient) RegistryLogin(auth types.AuthConfig) (types.AuthResponse, error) {
	return types.AuthResponse{}, nil
}

func (c *fakeClient) ServerVersion() (types.Version, error) {
	return types.Version{}, nil
}

func (c *fakeClient) VolumeCreate(options types.VolumeCreateRequest) (types.Volume, error) {
	return types.Volume{}, nil
}

func (c *fakeClient) VolumeInspect(volumeID string) (types.Volume, error) {
	return types.Volume{}, nil
}

func (c *fakeClient) VolumeList(filter filters.Args) (types.VolumesListResponse, error) {
	return types.VolumesListResponse{}, nil
}

func (c *fakeClient) VolumeRemove(volumeID string) error {
	return nil
}

var _ client.APIClient = &fakeClient{}
