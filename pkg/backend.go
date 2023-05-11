package pkg

import (
	"context"
	"strings"
	"sync"

	"github.com/hashicorp/vault/sdk/framework"
	"github.com/hashicorp/vault/sdk/logical"
	"github.com/surrealdb/surrealdb.go"
)

const backendHelp = ``

type hashiCupsClient struct {
}

type SurrealDBBackend struct {
	*framework.Backend
	lock   sync.RWMutex
	client *hashiCupsClient
}

func (b *SurrealDBBackend) reset() {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.client = nil
}

func Factory(ctx context.Context, conf *logical.BackendConfig) (logical.Backend, error) {
	b := backend()
	if err := b.Setup(ctx, conf); err != nil {
		return nil, err
	}
	return b, nil
}

func backend() *SurrealDBBackend {
	var b = SurrealDBBackend{}

	b.Backend = &framework.Backend{
		Help: strings.TrimSpace(backendHelp),
		PathsSpecial: &logical.Paths{
			LocalStorage: []string{},
			SealWrapStorage: []string{
				"config",
				"role/*",
			},
		},
		Paths:       framework.PathAppend(),
		Secrets:     []*framework.Secret{},
		BackendType: logical.TypeLogical,
		Invalidate:  b.invalidate,
	}
	return &b
}

func (b *SurrealDBBackend) invalidate(ctx context.Context, key string) {
	if key == "config" {
		b.reset()
	}
}

func something() {
	_, err := surrealdb.New("ws://localhost:8000/rpc")
	if err != nil {
		panic(err)
	}
}
