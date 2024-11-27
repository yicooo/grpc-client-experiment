// main.go
package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/url"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	v1 "simple-grpc/proto/service/v1"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cc, err := NewGRPCConnection("headless-svc.default.svc.cluster.local:50051", false)
	if err != nil {
		log.Fatalf("fuck me err, %s", err.Error())
	}

	ticker := time.NewTicker(100 * time.Millisecond)
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				out := &v1.Response{}
				cc.Invoke(ctx, "/service.v1.Service/ServiceImp", &v1.Request{Message: &[]string{time.Now().String()}[0]}, out)
				log.Printf("response: %s", out.GetMessage())
			}
		}
	}()

	<-ctx.Done()
	ticker.Stop()
}

func NewGRPCConnection(endpoint string, secure bool) (*grpc.ClientConn, error) {
	u, err := url.Parse(endpoint)
	if err != nil {
		return nil, err
	}

	creds := insecure.NewCredentials()
	if secure {
		creds = credentials.NewTLS(&tls.Config{
			ServerName: u.Host,
			MinVersion: tls.VersionTLS12,
		})
	}

	return grpc.NewClient(
		fmt.Sprintf("dns:///%s", endpoint),
		grpc.WithTransportCredentials(creds),
		grpc.WithDefaultServiceConfig(NewServiceConfig().EnableRoundRobin().ToString()),
		grpc.WithConnectParams(grpc.ConnectParams{
			Backoff: backoff.Config{
				BaseDelay:  time.Millisecond * 500,
				Multiplier: 1.1,
				Jitter:     0.23,
				MaxDelay:   time.Second,
			},
			MinConnectTimeout: time.Second,
		}),
	)
}
