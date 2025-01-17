package cli

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"strings"

	"github.com/tomcanham/iam/config"
	"golang.org/x/oauth2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/oauth"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

const (
	APP_NAME = "ias_clone"
)

var (
	errMissingMetadata = status.Errorf(codes.InvalidArgument, "missing metadata")
	errInvalidToken    = status.Errorf(codes.Unauthenticated, "invalid token")
)

type CLI struct {
	config.Config
	server *grpc.Server
}

func New() *CLI {
	config := config.New()

	cli := &CLI{
		Config: config,
	}

	return cli
}

func (c *CLI) GetEndpoint() string {
	return fmt.Sprintf("%s:%d", c.Host, c.Port)
}

func (c *CLI) NewClientConn() (*grpc.ClientConn, error) {
	addr := fmt.Sprintf("%s:%d", c.Host, c.Port)

	// Set up the credentials for the connection.
	perRPC := oauth.TokenSource{TokenSource: oauth2.StaticTokenSource(fetchToken())}
	creds, err := credentials.NewClientTLSFromFile(c.CertPath, c.Host)
	if err != nil {
		log.Fatalf("failed to load credentials: %v", err)
	}

	opts := []grpc.DialOption{
		// In addition to the following grpc.DialOption, callers may also use
		// the grpc.CallOption grpc.PerRPCCredentials with the RPC invocation
		// itself.
		// See: https://godoc.org/google.golang.org/grpc#PerRPCCredentials
		grpc.WithPerRPCCredentials(perRPC),
		// oauth.TokenSource requires the configuration of transport
		// credentials.
		// grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithTransportCredentials(creds),
	}

	conn, err := grpc.NewClient(addr, opts...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (c *CLI) Server() (*grpc.Server, error) {
	if c.server == nil {
		cert, err := tls.LoadX509KeyPair(c.CertPath, c.KeyPath)
		if err != nil {
			log.Fatalf("failed to load key pair: %s", err)
		}

		// // Parse the certificate
		// x509Cert, err := x509.ParseCertificate(cert.Certificate[0])
		// if err != nil {
		// 	log.Fatal(err)
		// }

		// // Pretty-print the certificate information
		// fmt.Println("Subject:", x509Cert.Subject)
		// fmt.Println("Issuer:", x509Cert.Issuer)
		// fmt.Println("Not Before:", x509Cert.NotBefore)
		// fmt.Println("Not After:", x509Cert.NotAfter)
		// fmt.Println("Serial Number:", x509Cert.SerialNumber)
		// fmt.Println("Signature Algorithm:", x509Cert.SignatureAlgorithm)

		opts := []grpc.ServerOption{
			// The following grpc.ServerOption adds an interceptor for all unary
			// RPCs. To configure an interceptor for streaming RPCs, see:
			// https://godoc.org/google.golang.org/grpc#StreamInterceptor
			grpc.UnaryInterceptor(ensureValidToken),
			// Enable TLS for all incoming connections.
			grpc.Creds(credentials.NewServerTLSFromCert(&cert)),
		}

		c.server = grpc.NewServer(opts...)
	}

	return c.server, nil
}

// fetchToken simulates a token lookup and omits the details of proper token
// acquisition. For examples of how to acquire an OAuth2 token, see:
// https://godoc.org/golang.org/x/oauth2
func fetchToken() *oauth2.Token {
	return &oauth2.Token{
		AccessToken: "some-secret-token",
	}
}

// ensureValidToken ensures a valid token exists within a request's metadata. If
// the token is missing or invalid, the interceptor blocks execution of the
// handler and returns an error. Otherwise, the interceptor invokes the unary
// handler.
func ensureValidToken(ctx context.Context, req any, _ *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
	md, ok := metadata.FromIncomingContext(ctx)

	if !ok {
		return nil, errMissingMetadata
	}

	log.Printf("in interceptor, md: %v", md)

	// The keys within metadata.MD are normalized to lowercase.
	// See: https://godoc.org/google.golang.org/grpc/metadata#New
	if !valid(md["authorization"]) {
		log.Printf("invalid token")
		return nil, errInvalidToken
	}
	// Continue execution of handler after ensuring a valid token.
	return handler(ctx, req)
}

// valid validates the authorization.
func valid(authorization []string) bool {
	if len(authorization) < 1 {
		return false
	}
	token := strings.TrimPrefix(authorization[0], "Bearer ")
	// Perform the token validation here. For the sake of this example, the code
	// here forgoes any of the usual OAuth2 token validation and instead checks
	// for a token matching an arbitrary string.
	return token == "some-secret-token"
}
