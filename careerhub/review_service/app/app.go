package app

import (
	"context"
	"os"

	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/mongocfg"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/common/vars"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/crawler"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/provider"
	"github.com/jae2274/careerhub-review-service/careerhub/review_service/restapi"
	"github.com/jae2274/goutils/llog"
	"github.com/jae2274/goutils/mw"
)

const (
	appName     = "review-service"
	serviceName = "careerhub"

	ctxKeyTraceID = string(mw.CtxKeyTraceID)
)

func initLogger(ctx context.Context) error {
	llog.SetMetadata("service", serviceName)
	llog.SetMetadata("app", appName)
	llog.SetDefaultContextData(ctxKeyTraceID)

	hostname, err := os.Hostname()
	if err != nil {
		return err
	}

	llog.SetMetadata("hostname", hostname)

	return nil
}

func Run(ctx context.Context) {
	err := initLogger(ctx)
	checkErr(ctx, err)
	llog.Info(ctx, "Start Application")

	envVars, err := vars.Variables()
	checkErr(ctx, err)
	db, err := mongocfg.NewDatabase(envVars.MongoUri, envVars.DbName, envVars.DBUser)
	checkErr(ctx, err)

	//TODO: implement collection struct
	// err := mongocfg.InitCollections(db, &matchjob.MatchJob{}, &scrapjob.ScrapJob{})
	// checkErr(ctx, err)

	runErr := make(chan error)

	go func() {
		err := provider.Run(ctx, envVars.ProviderGrpcPort, db)
		runErr <- err
	}()

	go func() {
		err := crawler.Run(ctx, envVars.CrawlerGrpcPort, db)
		runErr <- err
	}()

	go func() {
		err := restapi.Run(ctx, envVars.RestapiGrpcPort, db)
		runErr <- err
	}()

	select {
	case <-ctx.Done():
		llog.Info(ctx, "Finish Application")
	case err := <-runErr:
		checkErr(ctx, err)
	}
}

func checkErr(ctx context.Context, err error) {
	if err != nil {
		llog.LogErr(ctx, err)
		os.Exit(1)
	}
}
