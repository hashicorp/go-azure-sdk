package v2021_06_30

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/commands"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/commons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/connecttosourcemysqltasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/connecttosourcepostgresqlsynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/connecttosourcesqlservertasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/connecttotargetazuredbformysqltasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/connecttotargetazuredbforpostgresqlsynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/connecttotargetsqldbtasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/connecttotargetsqlmisynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/connecttotargetsqlmitasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/connecttotargetsqlsqldbsynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/customoperation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/delete"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/fieresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/fileresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/get"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/gettdecertificatessqltasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/getusertablesmysqltasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/getusertablessqlsynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/getusertablessqltasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/migratemysqlazuredbformysqlofflinetasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/migratemysqlazuredbformysqlsynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/migratepostgresqlazuredbforpostgresqlsynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/migrateschemasqlserversqldbtasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/migratesqlserversqldbsynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/migratesqlserversqldbtasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/migratesqlserversqlmisynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/migratesqlserversqlmitasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/migratessistasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/mongodbtasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/oracleazuredbpostgresqlsynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/patch"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/post"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/projectresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/put"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/serviceresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/servicetaskresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/servicetasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/standardoperation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/taskresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/tasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/taskscommons"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/validatemigrationinputsqlserversqlmisynctasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/validatemigrationinputsqlserversqlmitasks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/datamigration/2021-06-30/validatesyncmigrationinputsqlservertasks"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Commands                                       *commands.CommandsClient
	Commons                                        *commons.CommonsClient
	ConnectToSourceMySqlTasks                      *connecttosourcemysqltasks.ConnectToSourceMySqlTasksClient
	ConnectToSourcePostgreSqlSyncTasks             *connecttosourcepostgresqlsynctasks.ConnectToSourcePostgreSqlSyncTasksClient
	ConnectToSourceSqlServerTasks                  *connecttosourcesqlservertasks.ConnectToSourceSqlServerTasksClient
	ConnectToTargetAzureDbForMySqlTasks            *connecttotargetazuredbformysqltasks.ConnectToTargetAzureDbForMySqlTasksClient
	ConnectToTargetAzureDbForPostgreSqlSyncTasks   *connecttotargetazuredbforpostgresqlsynctasks.ConnectToTargetAzureDbForPostgreSqlSyncTasksClient
	ConnectToTargetSqlDbTasks                      *connecttotargetsqldbtasks.ConnectToTargetSqlDbTasksClient
	ConnectToTargetSqlMITasks                      *connecttotargetsqlmitasks.ConnectToTargetSqlMITasksClient
	ConnectToTargetSqlMiSyncTasks                  *connecttotargetsqlmisynctasks.ConnectToTargetSqlMiSyncTasksClient
	ConnectToTargetSqlSqlDbSyncTasks               *connecttotargetsqlsqldbsynctasks.ConnectToTargetSqlSqlDbSyncTasksClient
	CustomOperation                                *customoperation.CustomOperationClient
	DELETE                                         *delete.DELETEClient
	FieResource                                    *fieresource.FieResourceClient
	FileResource                                   *fileresource.FileResourceClient
	GET                                            *get.GETClient
	GetTdeCertificatesSqlTasks                     *gettdecertificatessqltasks.GetTdeCertificatesSqlTasksClient
	GetUserTablesMySqlTasks                        *getusertablesmysqltasks.GetUserTablesMySqlTasksClient
	GetUserTablesSqlSyncTasks                      *getusertablessqlsynctasks.GetUserTablesSqlSyncTasksClient
	GetUserTablesSqlTasks                          *getusertablessqltasks.GetUserTablesSqlTasksClient
	MigrateMySqlAzureDbForMySqlOfflineTasks        *migratemysqlazuredbformysqlofflinetasks.MigrateMySqlAzureDbForMySqlOfflineTasksClient
	MigrateMySqlAzureDbForMySqlSyncTasks           *migratemysqlazuredbformysqlsynctasks.MigrateMySqlAzureDbForMySqlSyncTasksClient
	MigratePostgreSqlAzureDbForPostgreSqlSyncTasks *migratepostgresqlazuredbforpostgresqlsynctasks.MigratePostgreSqlAzureDbForPostgreSqlSyncTasksClient
	MigrateSchemaSqlServerSqlDbTasks               *migrateschemasqlserversqldbtasks.MigrateSchemaSqlServerSqlDbTasksClient
	MigrateSqlServerSqlDbSyncTasks                 *migratesqlserversqldbsynctasks.MigrateSqlServerSqlDbSyncTasksClient
	MigrateSqlServerSqlDbTasks                     *migratesqlserversqldbtasks.MigrateSqlServerSqlDbTasksClient
	MigrateSqlServerSqlMITasks                     *migratesqlserversqlmitasks.MigrateSqlServerSqlMITasksClient
	MigrateSqlServerSqlMiSyncTasks                 *migratesqlserversqlmisynctasks.MigrateSqlServerSqlMiSyncTasksClient
	MigrateSsisTasks                               *migratessistasks.MigrateSsisTasksClient
	MongoDbTasks                                   *mongodbtasks.MongoDbTasksClient
	OracleAzureDbPostgreSqlSyncTasks               *oracleazuredbpostgresqlsynctasks.OracleAzureDbPostgreSqlSyncTasksClient
	PATCH                                          *patch.PATCHClient
	POST                                           *post.POSTClient
	PUT                                            *put.PUTClient
	ProjectResource                                *projectresource.ProjectResourceClient
	ServiceResource                                *serviceresource.ServiceResourceClient
	ServiceTaskResource                            *servicetaskresource.ServiceTaskResourceClient
	ServiceTasks                                   *servicetasks.ServiceTasksClient
	StandardOperation                              *standardoperation.StandardOperationClient
	TaskResource                                   *taskresource.TaskResourceClient
	Tasks                                          *tasks.TasksClient
	TasksCommons                                   *taskscommons.TasksCommonsClient
	ValidateMigrationInputSqlServerSqlMITasks      *validatemigrationinputsqlserversqlmitasks.ValidateMigrationInputSqlServerSqlMITasksClient
	ValidateMigrationInputSqlServerSqlMiSyncTasks  *validatemigrationinputsqlserversqlmisynctasks.ValidateMigrationInputSqlServerSqlMiSyncTasksClient
	ValidateSyncMigrationInputSqlServerTasks       *validatesyncmigrationinputsqlservertasks.ValidateSyncMigrationInputSqlServerTasksClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	commandsClient, err := commands.NewCommandsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Commands client: %+v", err)
	}
	configureFunc(commandsClient.Client)

	commonsClient, err := commons.NewCommonsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Commons client: %+v", err)
	}
	configureFunc(commonsClient.Client)

	connectToSourceMySqlTasksClient, err := connecttosourcemysqltasks.NewConnectToSourceMySqlTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectToSourceMySqlTasks client: %+v", err)
	}
	configureFunc(connectToSourceMySqlTasksClient.Client)

	connectToSourcePostgreSqlSyncTasksClient, err := connecttosourcepostgresqlsynctasks.NewConnectToSourcePostgreSqlSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectToSourcePostgreSqlSyncTasks client: %+v", err)
	}
	configureFunc(connectToSourcePostgreSqlSyncTasksClient.Client)

	connectToSourceSqlServerTasksClient, err := connecttosourcesqlservertasks.NewConnectToSourceSqlServerTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectToSourceSqlServerTasks client: %+v", err)
	}
	configureFunc(connectToSourceSqlServerTasksClient.Client)

	connectToTargetAzureDbForMySqlTasksClient, err := connecttotargetazuredbformysqltasks.NewConnectToTargetAzureDbForMySqlTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectToTargetAzureDbForMySqlTasks client: %+v", err)
	}
	configureFunc(connectToTargetAzureDbForMySqlTasksClient.Client)

	connectToTargetAzureDbForPostgreSqlSyncTasksClient, err := connecttotargetazuredbforpostgresqlsynctasks.NewConnectToTargetAzureDbForPostgreSqlSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectToTargetAzureDbForPostgreSqlSyncTasks client: %+v", err)
	}
	configureFunc(connectToTargetAzureDbForPostgreSqlSyncTasksClient.Client)

	connectToTargetSqlDbTasksClient, err := connecttotargetsqldbtasks.NewConnectToTargetSqlDbTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectToTargetSqlDbTasks client: %+v", err)
	}
	configureFunc(connectToTargetSqlDbTasksClient.Client)

	connectToTargetSqlMITasksClient, err := connecttotargetsqlmitasks.NewConnectToTargetSqlMITasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectToTargetSqlMITasks client: %+v", err)
	}
	configureFunc(connectToTargetSqlMITasksClient.Client)

	connectToTargetSqlMiSyncTasksClient, err := connecttotargetsqlmisynctasks.NewConnectToTargetSqlMiSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectToTargetSqlMiSyncTasks client: %+v", err)
	}
	configureFunc(connectToTargetSqlMiSyncTasksClient.Client)

	connectToTargetSqlSqlDbSyncTasksClient, err := connecttotargetsqlsqldbsynctasks.NewConnectToTargetSqlSqlDbSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ConnectToTargetSqlSqlDbSyncTasks client: %+v", err)
	}
	configureFunc(connectToTargetSqlSqlDbSyncTasksClient.Client)

	customOperationClient, err := customoperation.NewCustomOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CustomOperation client: %+v", err)
	}
	configureFunc(customOperationClient.Client)

	dELETEClient, err := delete.NewDELETEClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DELETE client: %+v", err)
	}
	configureFunc(dELETEClient.Client)

	fieResourceClient, err := fieresource.NewFieResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FieResource client: %+v", err)
	}
	configureFunc(fieResourceClient.Client)

	fileResourceClient, err := fileresource.NewFileResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building FileResource client: %+v", err)
	}
	configureFunc(fileResourceClient.Client)

	gETClient, err := get.NewGETClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GET client: %+v", err)
	}
	configureFunc(gETClient.Client)

	getTdeCertificatesSqlTasksClient, err := gettdecertificatessqltasks.NewGetTdeCertificatesSqlTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GetTdeCertificatesSqlTasks client: %+v", err)
	}
	configureFunc(getTdeCertificatesSqlTasksClient.Client)

	getUserTablesMySqlTasksClient, err := getusertablesmysqltasks.NewGetUserTablesMySqlTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GetUserTablesMySqlTasks client: %+v", err)
	}
	configureFunc(getUserTablesMySqlTasksClient.Client)

	getUserTablesSqlSyncTasksClient, err := getusertablessqlsynctasks.NewGetUserTablesSqlSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GetUserTablesSqlSyncTasks client: %+v", err)
	}
	configureFunc(getUserTablesSqlSyncTasksClient.Client)

	getUserTablesSqlTasksClient, err := getusertablessqltasks.NewGetUserTablesSqlTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GetUserTablesSqlTasks client: %+v", err)
	}
	configureFunc(getUserTablesSqlTasksClient.Client)

	migrateMySqlAzureDbForMySqlOfflineTasksClient, err := migratemysqlazuredbformysqlofflinetasks.NewMigrateMySqlAzureDbForMySqlOfflineTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrateMySqlAzureDbForMySqlOfflineTasks client: %+v", err)
	}
	configureFunc(migrateMySqlAzureDbForMySqlOfflineTasksClient.Client)

	migrateMySqlAzureDbForMySqlSyncTasksClient, err := migratemysqlazuredbformysqlsynctasks.NewMigrateMySqlAzureDbForMySqlSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrateMySqlAzureDbForMySqlSyncTasks client: %+v", err)
	}
	configureFunc(migrateMySqlAzureDbForMySqlSyncTasksClient.Client)

	migratePostgreSqlAzureDbForPostgreSqlSyncTasksClient, err := migratepostgresqlazuredbforpostgresqlsynctasks.NewMigratePostgreSqlAzureDbForPostgreSqlSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigratePostgreSqlAzureDbForPostgreSqlSyncTasks client: %+v", err)
	}
	configureFunc(migratePostgreSqlAzureDbForPostgreSqlSyncTasksClient.Client)

	migrateSchemaSqlServerSqlDbTasksClient, err := migrateschemasqlserversqldbtasks.NewMigrateSchemaSqlServerSqlDbTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrateSchemaSqlServerSqlDbTasks client: %+v", err)
	}
	configureFunc(migrateSchemaSqlServerSqlDbTasksClient.Client)

	migrateSqlServerSqlDbSyncTasksClient, err := migratesqlserversqldbsynctasks.NewMigrateSqlServerSqlDbSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrateSqlServerSqlDbSyncTasks client: %+v", err)
	}
	configureFunc(migrateSqlServerSqlDbSyncTasksClient.Client)

	migrateSqlServerSqlDbTasksClient, err := migratesqlserversqldbtasks.NewMigrateSqlServerSqlDbTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrateSqlServerSqlDbTasks client: %+v", err)
	}
	configureFunc(migrateSqlServerSqlDbTasksClient.Client)

	migrateSqlServerSqlMITasksClient, err := migratesqlserversqlmitasks.NewMigrateSqlServerSqlMITasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrateSqlServerSqlMITasks client: %+v", err)
	}
	configureFunc(migrateSqlServerSqlMITasksClient.Client)

	migrateSqlServerSqlMiSyncTasksClient, err := migratesqlserversqlmisynctasks.NewMigrateSqlServerSqlMiSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrateSqlServerSqlMiSyncTasks client: %+v", err)
	}
	configureFunc(migrateSqlServerSqlMiSyncTasksClient.Client)

	migrateSsisTasksClient, err := migratessistasks.NewMigrateSsisTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MigrateSsisTasks client: %+v", err)
	}
	configureFunc(migrateSsisTasksClient.Client)

	mongoDbTasksClient, err := mongodbtasks.NewMongoDbTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MongoDbTasks client: %+v", err)
	}
	configureFunc(mongoDbTasksClient.Client)

	oracleAzureDbPostgreSqlSyncTasksClient, err := oracleazuredbpostgresqlsynctasks.NewOracleAzureDbPostgreSqlSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building OracleAzureDbPostgreSqlSyncTasks client: %+v", err)
	}
	configureFunc(oracleAzureDbPostgreSqlSyncTasksClient.Client)

	pATCHClient, err := patch.NewPATCHClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PATCH client: %+v", err)
	}
	configureFunc(pATCHClient.Client)

	pOSTClient, err := post.NewPOSTClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building POST client: %+v", err)
	}
	configureFunc(pOSTClient.Client)

	pUTClient, err := put.NewPUTClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PUT client: %+v", err)
	}
	configureFunc(pUTClient.Client)

	projectResourceClient, err := projectresource.NewProjectResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProjectResource client: %+v", err)
	}
	configureFunc(projectResourceClient.Client)

	serviceResourceClient, err := serviceresource.NewServiceResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceResource client: %+v", err)
	}
	configureFunc(serviceResourceClient.Client)

	serviceTaskResourceClient, err := servicetaskresource.NewServiceTaskResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceTaskResource client: %+v", err)
	}
	configureFunc(serviceTaskResourceClient.Client)

	serviceTasksClient, err := servicetasks.NewServiceTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ServiceTasks client: %+v", err)
	}
	configureFunc(serviceTasksClient.Client)

	standardOperationClient, err := standardoperation.NewStandardOperationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StandardOperation client: %+v", err)
	}
	configureFunc(standardOperationClient.Client)

	taskResourceClient, err := taskresource.NewTaskResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TaskResource client: %+v", err)
	}
	configureFunc(taskResourceClient.Client)

	tasksClient, err := tasks.NewTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Tasks client: %+v", err)
	}
	configureFunc(tasksClient.Client)

	tasksCommonsClient, err := taskscommons.NewTasksCommonsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building TasksCommons client: %+v", err)
	}
	configureFunc(tasksCommonsClient.Client)

	validateMigrationInputSqlServerSqlMITasksClient, err := validatemigrationinputsqlserversqlmitasks.NewValidateMigrationInputSqlServerSqlMITasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ValidateMigrationInputSqlServerSqlMITasks client: %+v", err)
	}
	configureFunc(validateMigrationInputSqlServerSqlMITasksClient.Client)

	validateMigrationInputSqlServerSqlMiSyncTasksClient, err := validatemigrationinputsqlserversqlmisynctasks.NewValidateMigrationInputSqlServerSqlMiSyncTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ValidateMigrationInputSqlServerSqlMiSyncTasks client: %+v", err)
	}
	configureFunc(validateMigrationInputSqlServerSqlMiSyncTasksClient.Client)

	validateSyncMigrationInputSqlServerTasksClient, err := validatesyncmigrationinputsqlservertasks.NewValidateSyncMigrationInputSqlServerTasksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ValidateSyncMigrationInputSqlServerTasks client: %+v", err)
	}
	configureFunc(validateSyncMigrationInputSqlServerTasksClient.Client)

	return &Client{
		Commands:                                       commandsClient,
		Commons:                                        commonsClient,
		ConnectToSourceMySqlTasks:                      connectToSourceMySqlTasksClient,
		ConnectToSourcePostgreSqlSyncTasks:             connectToSourcePostgreSqlSyncTasksClient,
		ConnectToSourceSqlServerTasks:                  connectToSourceSqlServerTasksClient,
		ConnectToTargetAzureDbForMySqlTasks:            connectToTargetAzureDbForMySqlTasksClient,
		ConnectToTargetAzureDbForPostgreSqlSyncTasks:   connectToTargetAzureDbForPostgreSqlSyncTasksClient,
		ConnectToTargetSqlDbTasks:                      connectToTargetSqlDbTasksClient,
		ConnectToTargetSqlMITasks:                      connectToTargetSqlMITasksClient,
		ConnectToTargetSqlMiSyncTasks:                  connectToTargetSqlMiSyncTasksClient,
		ConnectToTargetSqlSqlDbSyncTasks:               connectToTargetSqlSqlDbSyncTasksClient,
		CustomOperation:                                customOperationClient,
		DELETE:                                         dELETEClient,
		FieResource:                                    fieResourceClient,
		FileResource:                                   fileResourceClient,
		GET:                                            gETClient,
		GetTdeCertificatesSqlTasks:                     getTdeCertificatesSqlTasksClient,
		GetUserTablesMySqlTasks:                        getUserTablesMySqlTasksClient,
		GetUserTablesSqlSyncTasks:                      getUserTablesSqlSyncTasksClient,
		GetUserTablesSqlTasks:                          getUserTablesSqlTasksClient,
		MigrateMySqlAzureDbForMySqlOfflineTasks:        migrateMySqlAzureDbForMySqlOfflineTasksClient,
		MigrateMySqlAzureDbForMySqlSyncTasks:           migrateMySqlAzureDbForMySqlSyncTasksClient,
		MigratePostgreSqlAzureDbForPostgreSqlSyncTasks: migratePostgreSqlAzureDbForPostgreSqlSyncTasksClient,
		MigrateSchemaSqlServerSqlDbTasks:               migrateSchemaSqlServerSqlDbTasksClient,
		MigrateSqlServerSqlDbSyncTasks:                 migrateSqlServerSqlDbSyncTasksClient,
		MigrateSqlServerSqlDbTasks:                     migrateSqlServerSqlDbTasksClient,
		MigrateSqlServerSqlMITasks:                     migrateSqlServerSqlMITasksClient,
		MigrateSqlServerSqlMiSyncTasks:                 migrateSqlServerSqlMiSyncTasksClient,
		MigrateSsisTasks:                               migrateSsisTasksClient,
		MongoDbTasks:                                   mongoDbTasksClient,
		OracleAzureDbPostgreSqlSyncTasks:               oracleAzureDbPostgreSqlSyncTasksClient,
		PATCH:                                          pATCHClient,
		POST:                                           pOSTClient,
		PUT:                                            pUTClient,
		ProjectResource:                                projectResourceClient,
		ServiceResource:                                serviceResourceClient,
		ServiceTaskResource:                            serviceTaskResourceClient,
		ServiceTasks:                                   serviceTasksClient,
		StandardOperation:                              standardOperationClient,
		TaskResource:                                   taskResourceClient,
		Tasks:                                          tasksClient,
		TasksCommons:                                   tasksCommonsClient,
		ValidateMigrationInputSqlServerSqlMITasks:      validateMigrationInputSqlServerSqlMITasksClient,
		ValidateMigrationInputSqlServerSqlMiSyncTasks: validateMigrationInputSqlServerSqlMiSyncTasksClient,
		ValidateSyncMigrationInputSqlServerTasks:      validateSyncMigrationInputSqlServerTasksClient,
	}, nil
}
