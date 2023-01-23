package informer

import (
	"context"
	"time"

	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

const informerResync = time.Hour * 24

// NewFactory constructs a new instance of sharedInformerFactory that starts with a manager
func NewFactory(m ctrl.Manager, client kubernetes.Interface) (informers.SharedInformerFactory, error) {
	factory := informers.NewSharedInformerFactory(client, informerResync)

	if err := m.Add(manager.RunnableFunc(func(ctx context.Context) error {
		m.GetLogger().WithName("informerFactory").Info("starting")
		factory.Start(ctx.Done())
		return nil
	})); err != nil {
		return nil, err
	}

	return factory, nil
}