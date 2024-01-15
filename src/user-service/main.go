package main

import (
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/crypto"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/database"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/grpc/user-service/proto"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/lib/health"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/api/router"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/auth"
	grpc_server "github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/grpc"
	"github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/service"
	user_controller "github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/user/controller"
	user_repository "github.com/akatranlp/hsfl-master-ai-cloud-engineering/user-service/user/repository"
	"github.com/caarlos0/env/v10"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type ApplicationConfig struct {
	Database          database.PsqlConfig `envPrefix:"POSTGRES_"`
	Jwt               auth.JwtConfig      `envPrefix:"JWT_"`
	AuthIsActive      bool                `env:"AUTH_IS_ACTIVE" envDefault:"false"`
	Port              uint16              `env:"PORT" envDefault:"8080"`
	GrpcPort          uint16              `env:"GRPC_PORT" envDefault:"8081"`
	GrpcCommunication bool                `env:"GRPC_COMMUNICATION" envDefault:"true"`
}

func main() {
	godotenv.Load()

	config := ApplicationConfig{}
	if err := env.Parse(&config); err != nil {
		log.Fatalf("Couldn't parse environment %s", err.Error())
	}

	tokenGenerator, err := auth.NewJwtTokenGenerator(config.Jwt)
	if err != nil {
		log.Fatalf("could not create JWT token generator: %s", err.Error())
	}

	userRepository, err := user_repository.NewPsqlRepository(config.Database)
	if err != nil {
		log.Fatalf("could not create user repository: %s", err.Error())
	}

	if err := userRepository.Migrate(); err != nil {
		log.Fatalf("could not migrate: %s", err.Error())
	}

	hasher := crypto.NewBcryptHasher()

	service := service.NewDefaultService(userRepository, tokenGenerator, config.AuthIsActive)
	healthController := health.NewDefaultController()

	controller := user_controller.NewDefaultController(userRepository, service, hasher, tokenGenerator, config.AuthIsActive)

	handler := router.New(controller, healthController)

	if config.GrpcCommunication {
		grpcAddr := fmt.Sprintf("0.0.0.0:%d", config.GrpcPort)
		listener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			log.Fatalf("could not listen: %v", err)
		}

		srv := grpc.NewServer()
		reflection.Register(srv)
		gprcServer := grpc_server.NewServer(service, tokenGenerator)
		proto.RegisterUserServiceServer(srv, gprcServer)

		go func() {
			log.Printf("GRPC-Server started on Port: %d\n", config.GrpcPort)
			if err := srv.Serve(listener); err != nil {
				log.Fatalf("could not serve: %v", err)
			}
		}()
	}

	log.Printf("REST-Server started on Port: %d\n", config.Port)
	addr := fmt.Sprintf("0.0.0.0:%d", config.Port)
	if err := http.ListenAndServe(addr, handler); err != nil {
		log.Fatalf("error while listen and serve: %s", err.Error())
	}
}
