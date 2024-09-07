package resolve

import (
	"google.golang.org/grpc/resolver"
	"log"
)

const scheme = "chris"

var serverAddresses = []string{"localhost:50051", "localhost:50052", "localhost:50053"}

// Builder implements resolver.Builder.
type Builder struct {
}

// Build creates a new resolver for the given target.
func (b *Builder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &simpleResolver{
		cc: cc,
	}
	r.start()
	return r, nil
}

func (*Builder) Scheme() string { return scheme }

// simpleResolver implements resolver.Resolver.
type simpleResolver struct {
	cc resolver.ClientConn
}

// start starts the resolver.
func (r *simpleResolver) start() {
	// In a real application, you would watch for changes in the backend servers here and call cc.UpdateState as needed.
	addrs := make([]resolver.Address, len(serverAddresses))
	for i, s := range serverAddresses {
		addrs[i] = resolver.Address{Addr: s}
	}

	if err := r.cc.UpdateState(resolver.State{Addresses: addrs}); err != nil {
		log.Fatal(err)
	}
}

func (*simpleResolver) ResolveNow(resolver.ResolveNowOptions) {}
func (*simpleResolver) Close()                                {}
