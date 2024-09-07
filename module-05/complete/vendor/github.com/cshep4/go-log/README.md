# golog
[![CircleCI](https://circleci.com/bb/curvecard/golog.svg?style=svg&circle-token=647ca9102ee21dc4ad4e2415cd04159b202f1e23)](https://circleci.com/bb/curvecard/golog)

## Some issues golog tries to solve
* Inconsistencies in how our services use their loggers. Many services would use
Go's standard 'log' package, logrus and slogger. This lead to strangeness and 
made it rudimentarily difficult to insert contextual data into logs, or even to
understand what was even going on.
* Effectively injecting contextual data without making things overly
complicated.
* Lack of flexibility due to slogger being a custom type. The only methods
exposed were `Info` and `Error`, taking away all of logrus's functionality.

## Examples
Fortunately golog has a very simple, light-weight and yet flexible API. It's
designed to be as fool-proof as possible production-ready out of the box.

### Basic logging
golog is opinionated about how many levels there should be and therefore only
exposes error and info levels.

```go
golog.Info(ctx, "hello, world!", zap.Int("pid", os.Getpid()))

golog.Error(ctx, "uhm", zap.Error("error", err))
```

### Advanced usage
Sometimes, it may not be enough to simply call the default `Error` and `Print`
methods. golog was designed to be flexible and provides `From` and `With`
which allow much greater control.

```go
l := golog.From(ctx)

// we now have a *zap.Logger, giving us full control. For instance if we wanted
// to add a hook to our logger, we can do so simply such:
l = l.WithOptions(zap.Hooks(func(e zap.Entry) error {
	// do something with our entry.
}))

// we can then put it back into the context to propagate these changes down the
// chain.
ctx = golog.With(ctx, l)
```

## Contextual log data injection

### gRPC Server
```go
import (
	"github.com/cshep4/go-log/logdefault/logdefaultgrpc"
	grpcmiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
	datadoggrpc "gopkg.in/DataDog/dd-trace-go.v1/contrib/google.golang.org/grpc"
)

var service string

ddopts := []datadoggrpc.Option{
	datadoggrpc.WithServiceName(service),
	datadoggrpc.WithAnalytics(true),
}

s := grpc.NewServer(
	grpc.StreamInterceptor(grpcmiddleware.ChainStreamServer(
		datadoggrpc.StreamServerInterceptor(ddopts...),
		logdefaultgrpc.StreamServerInterceptor,
	)),
	grpc.UnaryInterceptor(grpcmiddleware.ChainUnaryServer(
		datadoggrpc.UnaryServerInterceptor(ddopts...),
		logdefaultgrpc.UnaryServerInterceptor,
	)),
)
```

### HTTP Server
```go
import (
	"net/http"
	"github.com/cshep4/go-log"
	"github.com/cshep4/go-log/logdefault/logdefaulthttp"
	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/contrib/gorilla/mux"
)

var service string

r := mux.NewRouter(mux.WithServiceName(service))

r.Use(logdefaulthttp.Middleware)

r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
	golog.Print(r.Context(), "Hello!", zap.String("path", r.URL.Path))
})
```

## Testing
A simple example of using gologtest would be as follows:

```go
import (
	"context"
	
	"github.com/cshep4/go-log"
	"github.com/cshep4/go-log/gologtest"
	"github.com/stretchr/testify/assert"
)

const someMessage = "some-message"

o, ctx := gologtest.WithObserver(context.Background())

golog.Print(ctx, someMessage)

assert.NotZero(t, o.FilterMessage(someMessage).Len())
``` 
