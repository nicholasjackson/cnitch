package mocks

import (
	"io"
	"time"

	"golang.org/x/net/context"

	"github.com/moby/moby/api/types"
	"github.com/moby/moby/api/types/container"
	"github.com/moby/moby/api/types/filters"
	"github.com/moby/moby/api/types/network"
	"github.com/stretchr/testify/mock"
)

type DockerAPI struct {
	mock.Mock
}

func (m *DockerAPI) ContainerAttach(ctx context.Context, container string, options types.ContainerAttachOptions) (types.HijackedResponse, error) {
	return types.HijackedResponse{}, nil
}

func (m *DockerAPI) ContainerCommit(ctx context.Context, container string, options types.ContainerCommitOptions) (types.IDResponse, error) {
	return types.IDResponse{}, nil
}

func (cli *DockerAPI) ContainerCreate(ctx context.Context, config *container.Config, hostConfig *container.HostConfig, networkingConfig *network.NetworkingConfig, containerName string) (container.ContainerCreateCreatedBody, error) {
	return container.ContainerCreateCreatedBody{}, nil
}

func (cli *DockerAPI) ContainerDiff(ctx context.Context, containerID string) ([]container.ContainerChangeResponseItem, error) {
	return make([]container.ContainerChangeResponseItem, 0), nil
}

func (m *DockerAPI) ContainerExecAttach(ctx context.Context, execID string, config types.ExecConfig) (types.HijackedResponse, error) {
	return types.HijackedResponse{}, nil
}

func (m *DockerAPI) ContainerExecCreate(ctx context.Context, container string, config types.ExecConfig) (types.IDResponse, error) {
	return types.IDResponse{}, nil
}

func (m *DockerAPI) ContainerExecInspect(ctx context.Context, execID string) (types.ContainerExecInspect, error) {
	return types.ContainerExecInspect{}, nil
}

func (m *DockerAPI) ContainerExecResize(ctx context.Context, execID string, options types.ResizeOptions) error {
	return nil
}

func (m *DockerAPI) ContainerExecStart(ctx context.Context, execID string, config types.ExecStartCheck) error {
	return nil
}

func (m *DockerAPI) ContainerExport(ctx context.Context, container string) (io.ReadCloser, error) {
	return nil, nil
}

func (m *DockerAPI) ContainerInspect(ctx context.Context, container string) (types.ContainerJSON, error) {
	return types.ContainerJSON{}, nil
}

func (m *DockerAPI) ContainerInspectWithRaw(ctx context.Context, container string, getSize bool) (types.ContainerJSON, []byte, error) {
	return types.ContainerJSON{}, make([]byte, 0), nil
}

func (m *DockerAPI) ContainerKill(ctx context.Context, container, signal string) error {
	return nil
}

func (m *DockerAPI) ContainerList(ctx context.Context, options types.ContainerListOptions) ([]types.Container, error) {
	return make([]types.Container, 0), nil
}

func (m *DockerAPI) ContainerLogs(ctx context.Context, container string, options types.ContainerLogsOptions) (io.ReadCloser, error) {
	return nil, nil
}

func (m *DockerAPI) ContainerPause(ctx context.Context, container string) error {
	return nil
}

func (m *DockerAPI) ContainerRemove(ctx context.Context, container string, options types.ContainerRemoveOptions) error {
	return nil
}

func (m *DockerAPI) ContainerRename(ctx context.Context, container, newContainerName string) error {
	return nil
}

func (m *DockerAPI) ContainerResize(ctx context.Context, container string, options types.ResizeOptions) error {
	return nil
}

func (m *DockerAPI) ContainerRestart(ctx context.Context, container string, timeout *time.Duration) error {
	return nil
}

func (m *DockerAPI) ContainerStatPath(ctx context.Context, container, path string) (types.ContainerPathStat, error) {
	return types.ContainerPathStat{}, nil
}

func (m *DockerAPI) ContainerStats(ctx context.Context, container string, stream bool) (types.ContainerStats, error) {
	return types.ContainerStats{}, nil
}

func (m *DockerAPI) ContainerStart(ctx context.Context, container string, options types.ContainerStartOptions) error {
	return nil
}

func (m *DockerAPI) ContainerStop(ctx context.Context, container string, timeout *time.Duration) error {
	return nil
}

func (m *DockerAPI) ContainerUnpause(ctx context.Context, container string) error {
	return nil
}

func (m *DockerAPI) ContainerUpdate(ctx context.Context, containerID string, updateConfig container.UpdateConfig) (container.ContainerUpdateOKBody, error) {
	return container.ContainerUpdateOKBody{}, nil
}

func (m *DockerAPI) CopyFromContainer(ctx context.Context, container, srcPath string) (io.ReadCloser, types.ContainerPathStat, error) {
	return nil, types.ContainerPathStat{}, nil
}

func (m *DockerAPI) CopyToContainer(ctx context.Context, container, path string, content io.Reader, options types.CopyToContainerOptions) error {
	return nil
}

func (m *DockerAPI) ContainersPrune(ctx context.Context, pruneFilters filters.Args) (types.ContainersPruneReport, error) {
	return types.ContainersPruneReport{}, nil
}
