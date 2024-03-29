package provider

import (
	"context"
	"os"

	"go.jetpack.io/envsec"
)

type EnvSec interface {
	// string here refers to "selected provider"
	Get(context.Context, string) (EnvSecConfig, error)
	NewEnvId(ctx context.Context, projectId string, env string) (*envsec.EnvID, error)
}

type EnvSecConfig interface {
	GetRegion() string
	GetAccessKeyId() string
	GetSecretAccessKey() string
	GetSessionToken() string
	GetKmsKeyId() string
}

type defaultEnvSec struct {
}

func DefaultEnvSecProvider() EnvSec {
	return &defaultEnvSec{}
}

// Both "jetpack" and "" are valid as input for selectedProvider.
func (p *defaultEnvSec) Get(ctx context.Context, selectedProvider string) (EnvSecConfig, error) {
	return nil, nil
}

func (p *defaultEnvSec) NewEnvId(_ctx context.Context, projectId string, env string) (*envsec.EnvID, error) {
	orgId := os.Getenv("LAUNCHPAD_ORG_ID")
	envId, err := envsec.NewEnvID(projectId, orgId, env)
	return &envId, err
}
