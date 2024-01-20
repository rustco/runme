package cmd

import (
	"crypto/tls"
	"net"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stateful/runme/internal/document/editor/editorservice"
	parserv1 "github.com/stateful/runme/internal/gen/proto/go/runme/parser/v1"
	projectv1 "github.com/stateful/runme/internal/gen/proto/go/runme/project/v1"
	runnerv1 "github.com/stateful/runme/internal/gen/proto/go/runme/runner/v1"
	"github.com/stateful/runme/internal/project/projectservice"
	"github.com/stateful/runme/internal/runner"
	runnerv2service "github.com/stateful/runme/internal/runnerv2service"
	runmetls "github.com/stateful/runme/internal/tls"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthgrpc "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

func serverCmd() *cobra.Command {
	const (
		defaultAddr = "localhost:7863"
	)

	var (
		addr         string
		devMode      bool
		enableRunner bool
		tlsDir       string
	)

	cmd := cobra.Command{
		Hidden: true,
		Use:    "server",
		Short:  "Start a server with various services and a gRPC interface.",
		Long: `The server provides two services: kernel and parser.

The parser allows serializing and deserializing markdown content.

The kernel is used to run long running processes like shells and interacting with them.`,
		Args: cobra.NoArgs,
		PreRun: func(cmd *cobra.Command, args []string) {
			// By default, we want to log when running the server command.
			fLogEnabled = true
			// By default, we want to log to stderr when running the server command.
			if !cmd.Flags().Changed("log-file") {
				fLogFilePath = "stderr"
			}
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			logger, err := getLogger(devMode)
			if err != nil {
				return err
			}
			defer logger.Sync()

			var tlsConfig *tls.Config

			if !fInsecure {
				tlsConfig, err = runmetls.GenerateTLS(tlsDir, tlsFileMode, logger)

				if err != nil {
					return err
				}
			}

			var lis net.Listener
			protocol := "tcp"

			if strings.HasPrefix(addr, "unix://") {
				addr = strings.TrimPrefix(addr, "unix://")

				// TODO: consolidate removing address into a single place
				_ = os.Remove(addr)
				protocol = "unix"

				defer func() { _ = os.Remove(addr) }()
			}

			if tlsConfig == nil {
				lis, err = net.Listen(protocol, addr)
			} else {
				lis, err = tls.Listen(protocol, addr, tlsConfig)
			}

			if err != nil {
				return err
			}

			logger.Info("started listening", zap.String("addr", lis.Addr().String()))

			const maxMsgSize = 4 * 1024 * 1024 // 4 MiB

			server := grpc.NewServer(
				grpc.MaxRecvMsgSize(maxMsgSize),
				grpc.MaxSendMsgSize(maxMsgSize),
			)
			parserv1.RegisterParserServiceServer(server, editorservice.NewParserServiceServer(logger))
			projectv1.RegisterProjectServiceServer(server, projectservice.NewProjectServiceServer(logger))
			if enableRunner {
				runnerServicev1, err := runner.NewRunnerService(logger)
				if err != nil {
					return err
				}
				runnerv1.RegisterRunnerServiceServer(server, runnerServicev1)

				runnerServicev2, err := runnerv2service.NewRunnerService(logger)
				if err != nil {
					return err
				}
				runnerv2alpha1.RegisterRunnerServiceServer(server, runnerServicev2)
			}

			healthcheck := health.NewServer()
			healthgrpc.RegisterHealthServer(server, healthcheck)
			// Setting SERVING for the whole system.
			healthcheck.SetServingStatus("", healthgrpc.HealthCheckResponse_SERVING)

			reflection.Register(server)
			return server.Serve(lis)
		},
	}

	setDefaultFlags(&cmd)

	cmd.Flags().StringVarP(&addr, "address", "a", defaultAddr, "Address to create unix (unix:///path/to/socket) or IP socket (localhost:7890)")
	cmd.Flags().BoolVar(&devMode, "dev", false, "Enable development mode")
	cmd.Flags().BoolVar(&enableRunner, "runner", true, "Enable runner service (legacy, defaults to true)")
	cmd.Flags().StringVar(&tlsDir, "tls", defaultTLSDir, "Directory in which to generate TLS certificates & use for all incoming and outgoing messages")

	_ = cmd.Flags().MarkHidden("runner")

	return &cmd
}
