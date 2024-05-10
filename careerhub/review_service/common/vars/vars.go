package vars

import (
	"fmt"
	"os"
	"strconv"
)

type DBUser struct {
	Username string
	Password string
}

type Vars struct {
	MongoUri         string
	DbName           string
	DBUser           *DBUser
	ProviderGrpcPort int
	CrawlerGrpcPort  int
	RestapiGrpcPort  int
}

type ErrNotExistedVar struct {
	VarName string
}

func NotExistedVar(varName string) *ErrNotExistedVar {
	return &ErrNotExistedVar{VarName: varName}
}

func (e *ErrNotExistedVar) Error() string {
	return fmt.Sprintf("%s is not existed", e.VarName)
}

func Variables() (*Vars, error) {
	mongoUri, err := getFromEnv("MONGO_URI")
	if err != nil {
		return nil, err
	}

	dbUsername := getFromEnvPtr("DB_USERNAME")
	dbPassword := getFromEnvPtr("DB_PASSWORD")

	var dbUser *DBUser
	if dbUsername != nil && dbPassword != nil {
		dbUser = &DBUser{
			Username: *dbUsername,
			Password: *dbPassword,
		}
	}

	dbName, err := getFromEnv("DB_NAME")
	if err != nil {
		return nil, err
	}

	providerGrpcPort, err := getFromEnv("PROVIDER_GRPC_PORT")
	if err != nil {
		return nil, err
	}

	providerGrpcPortInt, err := strconv.ParseInt(providerGrpcPort, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("PROVIDER_GRPC_PORT is not integer.\tPROVIDER_GRPC_PORT: %s", providerGrpcPort)
	}

	crawlerGrpcPort, err := getFromEnv("CRAWLER_GRPC_PORT")
	if err != nil {
		return nil, err
	}
	crawlerGrpcPortInt, err := strconv.ParseInt(crawlerGrpcPort, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("CRAWLER_GRPC_PORT is not integer.\tCRAWLER_GRPC_PORT: %s", crawlerGrpcPort)
	}

	restapiGrpcPort, err := getFromEnv("RESTAPI_GRPC_PORT")
	if err != nil {
		return nil, err
	}

	restapiGrpcPortInt, err := strconv.ParseInt(restapiGrpcPort, 10, 32)
	if err != nil {
		return nil, fmt.Errorf("RESTAPI_GRPC_PORT is not integer.\tRESTAPI_GRPC_PORT: %s", restapiGrpcPort)
	}

	return &Vars{
		MongoUri:         mongoUri,
		DBUser:           dbUser,
		DbName:           dbName,
		ProviderGrpcPort: int(providerGrpcPortInt),
		CrawlerGrpcPort:  int(crawlerGrpcPortInt),
		RestapiGrpcPort:  int(restapiGrpcPortInt),
	}, nil
}

func getFromEnv(envVar string) (string, error) {
	ev := os.Getenv(envVar)

	if ev == "" {
		return "", fmt.Errorf("%s is not existed", envVar)
	}

	return ev, nil
}

func getFromEnvPtr(envVar string) *string {
	ev := os.Getenv(envVar)

	if ev == "" {
		return nil
	}

	return &ev
}
