package mongo

import (
	"context"
	"sync"

	"github.com/phuhao00/broker"
	mongobrocker "github.com/phuhao00/broker/mongo"
)

var (
	Client        *mongobrocker.Client
	onceInitMongo sync.Once
)

func init() {
	ctx := context.Background()
	onceInitMongo.Do(func() {
		tc := &mongobrocker.Client{
			BaseComponent: broker.NewBaseComponent(),
			RealCli: mongobrocker.NewClient(ctx, &mongobrocker.Config{
				URI:         "mongodb://huaweiyun:27017",
				MinPoolSize: 3,
				MaxPoolSize: 3000,
			}),
		}
		tc.Launch()
		defer tc.Stop()
	})
}
