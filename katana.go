package katana

import (
	"context"
	"fmt"
	"github.com/yoshiakiley/katana/common"
	"github.com/yoshiakiley/katana/core"
	"github.com/yoshiakiley/katana/store"
	"github.com/yoshiakiley/katana/store/mongo"
)

var serviceMap = map[string]any{}

func InitStore(ctx context.Context, initStore common.StoreType, addr string) error {
	switch initStore {
	case common.StoreTypeMongo:
		return mongo.InitMongoCli(ctx, addr)
	}
	return nil
}

func NewService[R core.IObject](schema, collection string) *Service[R] {
	service := &Service[R]{
		Schema:     schema,
		Collection: collection,
	}
	service.Set(store.NewStore[R](mongo.NewMongoCli[R](), nil))

	return service
}

func GetService[R core.IObject](schema, collection, resource string) *Service[R] {
	serviceName := fmt.Sprintf("%s-%s-%s", schema, collection, resource)
	if service, exists := serviceMap[serviceName]; exists {
		return service.(*Service[R])

	}

	service := &Service[R]{
		Schema:     schema,
		Collection: collection,
	}
	service.Set(store.NewStore[R](mongo.NewMongoCli[R](), nil))
	serviceMap[serviceName] = service
	return service
}
