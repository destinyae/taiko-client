package helper

import (
	"context"

	"github.com/cenkalti/backoff/v4"
	"github.com/taikoxyz/taiko-client/pkg/jwt"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
	"github.com/taikoxyz/taiko-client/testutils"
)

func NewWsRpcClientConfig(s *testutils.ClientSuite) *rpc.ClientConfig {
	jwtSecret, err := jwt.ParseSecretFromFile(testutils.JwtSecretFile)
	s.NoError(err)
	return &rpc.ClientConfig{
		L1Endpoint:        s.L1.WsEndpoint(),
		L2Endpoint:        s.L2.WsEndpoint(),
		TaikoL1Address:    testutils.TaikoL1Address,
		TaikoTokenAddress: testutils.TaikoL1TokenAddress,
		TaikoL2Address:    testutils.TaikoL2Address,
		L2EngineEndpoint:  s.L2.AuthEndpoint(),
		JwtSecret:         string(jwtSecret),
		RetryInterval:     backoff.DefaultMaxInterval,
	}
}

func NewWsRpcClient(s *testutils.ClientSuite) *rpc.Client {
	cli, err := rpc.NewClient(context.Background(), NewWsRpcClientConfig(s))
	s.NoError(err)
	return cli
}
