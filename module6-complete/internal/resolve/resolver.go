package resolve

import (
	"errors"
	"google.golang.org/grpc/resolver"
	"log"
)

const scheme = "chris"

// builder implements resolver.Builder.
type builder struct {
	addresses []string
}

func NewBuilder(addresses []string) (*builder, error) {
	if len(addresses) == 0 {
		return nil, errors.New("addresses cannot be empty")
	}

	return &builder{addresses: addresses}, nil
}

// Build creates a new resolver for the given target.
func (b *builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &simpleResolver{
		cc:        cc,
		addresses: b.addresses,
	}
	r.start()
	return r, nil
}

func (*builder) Scheme() string { return scheme }

// simpleResolver implements resolver.Resolver.
type simpleResolver struct {
	cc        resolver.ClientConn
	addresses []string
}

// start starts the resolver.
func (r *simpleResolver) start() {
	// In a real application, you would watch for changes in the backend servers here and call cc.UpdateState as needed.
	addrs := make([]resolver.Address, len(r.addresses))
	for i, s := range r.addresses {
		addrs[i] = resolver.Address{Addr: s}
	}

	if err := r.cc.UpdateState(resolver.State{Addresses: addrs}); err != nil {
		log.Fatal(err)
	}
}

func (*simpleResolver) ResolveNow(resolver.ResolveNowOptions) {}
func (*simpleResolver) Close()                                {}
