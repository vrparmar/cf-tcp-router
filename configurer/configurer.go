package configurer

import (
	"errors"

	cf_tcp_router "github.com/GESoftware-CF/cf-tcp-router"
	"github.com/GESoftware-CF/cf-tcp-router/configurer/haproxy"
	"github.com/pivotal-golang/lager"
)

const (
	HaProxyConfigurer = "HAProxy"
)

type RouterConfigurer interface {
	MapBackendHostsToAvailablePort(backendHostInfos cf_tcp_router.BackendHostInfos) (cf_tcp_router.RouterHostInfo, error)
}

func NewConfigurer(logger lager.Logger, tcpLoadBalancer string) RouterConfigurer {
	switch tcpLoadBalancer {
	case HaProxyConfigurer:
		return haproxy.NewHaProxyConfigurer(logger)
	default:
		logger.Fatal("not-supported", errors.New("unsupported tcp load balancer"), lager.Data{"tcp_load_balancer": tcpLoadBalancer})
		return nil
	}
}