package loadbalancer

import (
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/balancer/base"
	"google.golang.org/grpc/metadata"
	"log"
)

const Name = "ab_testing"

func NewBuilder(groups map[string]string, defaultAddr string) balancer.Builder {
	return base.NewBalancerBuilder(Name, &pickerBuilder{
		groups:      groups,
		defaultAddr: defaultAddr,
	}, base.Config{HealthCheck: true})
}

type pickerBuilder struct {
	groups      map[string]string
	defaultAddr string
}

func (p *pickerBuilder) Build(info base.PickerBuildInfo) balancer.Picker {
	if len(info.ReadySCs) == 0 {
		return base.NewErrPicker(balancer.ErrNoSubConnAvailable)
	}

	scs := make(map[string]balancer.SubConn)
	for sc, inf := range info.ReadySCs {
		scs[inf.Address.Addr] = sc
	}

	return &picker{
		subConns:    scs,
		groups:      p.groups,
		defaultAddr: p.defaultAddr,
	}
}

type picker struct {
	subConns    map[string]balancer.SubConn
	groups      map[string]string
	defaultAddr string
}

func (p *picker) Pick(info balancer.PickInfo) (balancer.PickResult, error) {
	md, ok := metadata.FromOutgoingContext(info.Ctx)
	if !ok {
		log.Println("unable to get metadata from context, using default address")
		return p.defaultConn()
	}

	name := md.Get("user-group")
	if len(name) < 1 {
		log.Println("group not specified in metadata, using default address")
		return p.defaultConn()
	}

	addr, ok := p.groups[name[0]]
	if !ok {
		log.Println("group not in list, using default address")
		return p.defaultConn()
	}

	subConn, ok := p.subConns[addr]
	if !ok {
		log.Println("addr not in list of addrs, using default address")
		return p.defaultConn()
	}

	return balancer.PickResult{SubConn: subConn}, nil
}

func (p *picker) defaultConn() (balancer.PickResult, error) {
	conn, ok := p.subConns[p.defaultAddr]
	if !ok {
		return balancer.PickResult{}, balancer.ErrNoSubConnAvailable
	}

	return balancer.PickResult{SubConn: conn}, nil
}
