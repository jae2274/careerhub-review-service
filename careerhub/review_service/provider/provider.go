package provider

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/provider_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/repo"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider/server"
	"github.com/jae2274/careerhub-review-service/utils"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, grpcPort int, db *mongo.Database) error {
	companyRepo := repo.NewCompanyRepo(db)
	reviewGrpcServer := server.NewCrawlingTaskGrpcServer(companyRepo)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Start suggester grpc server").Level(llog.INFO).Data("port", grpcPort).Log(ctx)

	grpcServer := grpc.NewServer(utils.Middlewares()...)
	provider_grpc.RegisterCrawlingTaskGrpcServer(grpcServer, reviewGrpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
