package authorizations

import (
	"context"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	opentracing "github.com/opentracing/opentracing-go"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"

	users_grpc "github.com/SPooNqir/users"
	lib "gitlab.com/SpoonQIR/Cloud/library/golang-common.git"
)

type AuthorizationsService struct {
	Authsconn *grpc.ClientConn
	Authssvc  AuthorizationsClient
	Authsreco chan bool

	Id *lib.Identity
}

// InitAuthorizations tst
func (s *AuthorizationsService) InitAuthorizations(authorizationHost string, tracer opentracing.Tracer, logger *logrus.Logger) chan bool {
	logentry := logrus.NewEntry(logger)
	logopts := []grpc_logrus.Option{
		grpc_logrus.WithDurationField(func(duration time.Duration) (key string, value interface{}) {
			return "grpc.time_ns", duration.Nanoseconds()
		}),
	}

	otopts := []grpc_opentracing.Option{
		grpc_opentracing.WithTracer(tracer),
	}

	var err error

	connect := make(chan bool)

	go func(lconn chan bool) {
		for {
			logrus.Info("Wait for connect")
			r := <-lconn
			logrus.WithFields(logrus.Fields{"reconn": r}).Info("conn chan receive")
			if r {
				for i := 1; i < 5; i++ {
					// s.authconn, s.Authssvc,
					s.Authsconn, err = grpc.Dial(authorizationHost,
						grpc.WithInsecure(),
						grpc.WithUnaryInterceptor(grpc_middleware.ChainUnaryClient(
							grpc_logrus.UnaryClientInterceptor(logentry, logopts...),
							grpc_opentracing.UnaryClientInterceptor(otopts...),
							grpc_prometheus.UnaryClientInterceptor,
						)),
						grpc.WithStreamInterceptor(grpc_middleware.ChainStreamClient(
							grpc_logrus.StreamClientInterceptor(logentry, logopts...),
							grpc_opentracing.StreamClientInterceptor(otopts...),
							grpc_prometheus.StreamClientInterceptor,
						)),
					)
					if err != nil {
						logger.Fatalf("did not connect: %v, try : %d - sleep 5s", err, i)
						time.Sleep(2 * time.Second)
					} else {
						s.Authssvc = NewAuthorizationsClient(s.Authsconn)
						break
					}
				}
			} else {
				logrus.Info("end of goroutine - reconnect")
				return
			}
		}
	}(connect)

	logger.WithFields(logrus.Fields{"host": authorizationHost}).Info("Connexion au service gRPC 'Authorizations'")
	connect <- true
	return connect
}

func (s *AuthorizationsService) GetAuthorizations(ctx context.Context, usr *users_grpc.User) (*Authorizations, error) {
	for i := 1; i <= 5; i++ {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			logrus.Error("cannot get outgoing data")
		}
		logrus.WithFields(logrus.Fields{"usr": usr, "jwt": md.Get("authorization")}).Debug("get authorizations")
		grp, err := s.Authssvc.Get(metadata.NewOutgoingContext(ctx, md), usr)
		logrus.WithFields(logrus.Fields{"ctx.err": ctx.Err(), "err": err}).Trace("error ctx get object")
		if err != nil {
			logrus.WithFields(logrus.Fields{"err": err}).Error("error get object")
			errStatus, _ := status.FromError(err)
			if errStatus.Code() == codes.Unavailable {
				s.Authsreco <- true
			} else if errStatus.Code() == codes.Canceled {
				s.Authsreco <- true
			} else if errStatus.Code() == codes.DeadlineExceeded {
				s.Authsreco <- true
			} else if errStatus.Code() == codes.Aborted {
				s.Authsreco <- true
			} else if errStatus.Code() == codes.Unauthenticated {
				logrus.WithFields(logrus.Fields{"jwt": md.Get("authorization")}).Info("ws-identity not identified")
				return nil, status.Error(codes.Unauthenticated, "unauthenticated")
			} else if errStatus.Code() == codes.InvalidArgument {
				return nil, status.Errorf(codes.InvalidArgument, "argument invalid %v", err)
			} else if errStatus.Code() == codes.NotFound {
				return nil, nil
			}
			// errStatus.Code() == codes.Internal = retry
		} else if ctx.Err() != nil {
			s.Authsreco <- true
		} else {
			return grp, nil
		}
	}
	return nil, status.Errorf(codes.NotFound, "authorization not found")
}
