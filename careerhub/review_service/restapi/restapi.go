package restapi

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi/repo"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi/restapi_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi/server"
	"github.com/jae2274/careerhub-review-service/utils"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, grpcPort int, db *mongo.Database) error {
	companyRepo := repo.NewCompanyRepo(db)
	reviewRepo := repo.NewReviewRepo(db)
	grpcService := server.NewReviewReaderGrpcServer(companyRepo, reviewRepo)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Start restapi grpc server").Level(llog.INFO).Data("port", grpcPort).Log(ctx)

	grpcServer := grpc.NewServer(utils.Middlewares()...)
	restapi_grpc.RegisterReviewReaderGrpcServer(grpcServer, grpcService)

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
