package conf42

import (
	"context"
	"fmt"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultIndexHostFilePath = "/usr/share/nginx/html/index.html"
)

// Conf42Container represents the Conf42 container type used in the module.
type Conf42Container struct {
	testcontainers.Container
	URI string
}

// RunContainer creates an instance of the Conf42Container type.
func RunContainer(ctx context.Context, opts ...testcontainers.ContainerCustomizer) (*Conf42Container, error) {
	req := testcontainers.ContainerRequest{
		Image:        "nginx",
		ExposedPorts: []string{"80/tcp"},
		WaitingFor:   wait.ForHTTP("/").WithStartupTimeout(10 * time.Second),
	}
	genericContainerReq := testcontainers.GenericContainerRequest{
		ContainerRequest: req,
		Started:          true,
	}

	for _, opt := range opts {
		opt.Customize(&genericContainerReq)
	}

	container, err := testcontainers.GenericContainer(ctx, genericContainerReq)
	if err != nil {
		return nil, err
	}

	ip, err := container.Host(ctx)
	if err != nil {
		return nil, err
	}

	mappedPort, err := container.MappedPort(ctx, "80")
	if err != nil {
		return nil, err
	}

	uri := fmt.Sprintf("http://%s:%s", ip, mappedPort.Port())

	return &Conf42Container{Container: container, URI: uri}, nil
}

// WithIndex allows to define a new index file
func WithIndex(indexHostFilePath string) testcontainers.CustomizeRequestOption {
	return func(req *testcontainers.GenericContainerRequest) {
		req.Files = append(req.Files, testcontainers.ContainerFile{
			HostFilePath:      indexHostFilePath,
			ContainerFilePath: defaultIndexHostFilePath,
			FileMode:          0o644,
		})
	}
}
