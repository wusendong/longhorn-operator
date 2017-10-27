package monitoring

import (
	"test/k8s/client/monitoring/v1"

	"k8s.io/client-go/rest"
	"k8s.io/client-go/util/flowcontrol"
)

var _ Interface = &Clientset{}

type Interface interface {
	MonitoringV1() v1.MonitoringV1Interface
}

type Clientset struct {
	*v1.MonitoringV1Client
}

func (c *Clientset) MonitoringV1() v1.MonitoringV1Interface {
	if c == nil {
		return nil
	}
	return c.MonitoringV1Client
}

func NewForConfig(apiGroup string, c *rest.Config) (*Clientset, error) {
	configShallowCopy := *c
	if configShallowCopy.RateLimiter == nil && configShallowCopy.QPS > 0 {
		configShallowCopy.RateLimiter = flowcontrol.NewTokenBucketRateLimiter(configShallowCopy.QPS, configShallowCopy.Burst)
	}
	var cs Clientset
	var err error

	cs.MonitoringV1Client, err = v1.NewForConfig(apiGroup, &configShallowCopy)
	if err != nil {
		return nil, err
	}

	return &cs, nil
}
