package crawler

import (
	"context"
	"fmt"
	"net"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/crawler_grpc"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/repo"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler/server"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/utils"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/terr"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/grpc"
)

func Run(ctx context.Context, grpcPort int, db *mongo.Database) error {
	companyRepo := repo.NewCompanyRepo(db)
	reviewRepo := repo.NewReviewRepo(db)
	reviewGrpcServer := server.NewReviewGrpcServer(companyRepo, reviewRepo)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		return terr.Wrap(err)
	}

	llog.Msg("Start crawler grpc server").Level(llog.INFO).Data("port", grpcPort).Log(ctx)

	grpcServer := grpc.NewServer(utils.Middlewares()...)
	crawler_grpc.RegisterReviewGrpcServer(grpcServer, reviewGrpcServer)

	err = grpcServer.Serve(listener)
	if err != nil {
		return terr.Wrap(err)
	}

	return nil
}
