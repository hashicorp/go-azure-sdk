package environments

import (
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/azurecli"
)

const azureCliMinimumVersion = "2.0.21"

type azureCliSchema struct {
	IsActive bool   `json:"isActive"`
	Name     string `json:"name"`
	Profile  string `json:"profile"`

	Endpoints struct {
		ActiveDirectory                       string `json:"activeDirectory"`
		ActiveDirectoryDataLakeResourceID     string `json:"activeDirectoryDataLakeResourceId"`
		ActiveDirectoryGraphResourceID        string `json:"activeDirectoryGraphResourceId"`
		ActiveDirectoryResourceID             string `json:"activeDirectoryResourceId"`
		AppInsightsResourceID                 string `json:"appInsightsResourceId"`
		AppInsightsTelemetryChannelResourceID string `json:"appInsightsTelemetryChannelResourceId"`
		AttestationResourceID                 string `json:"attestationResourceId"`
		AZMirrorStorageAccountResourceID      string `json:"azmirrorStorageAccountResourceId"`
		BatchResourceID                       string `json:"batchResourceId"`
		Gallery                               string `json:"gallery"`
		LogAnalyticsResourceID                string `json:"logAnalyticsResourceId"`
		Management                            string `json:"management"`
		MediaResourceID                       string `json:"mediaResourceId"`
		MicrosoftGraphResourceID              string `json:"microsoftGraphResourceId"`
		OSSRDBMSResourceID                    string `json:"ossrdbmsResourceId"`
		Portal                                string `json:"portal"`
		ResourceManager                       string `json:"resourceManager"`
		SqlManagement                         string `json:"sqlManagement"`
		SynapseAnalyticsResourceID            string `json:"synapseAnalyticsResourceId"`
		VMImageAliasDoc                       string `json:"vmImageAliasDoc"`
	} `json:"endpoints"`

	Suffixes struct {
		AcrLoginServerEndpoint                      string `json:"acrLoginServerEndpoint"`
		AttestationEndpoint                         string `json:"attestationEndpoint"`
		AzureDataLakeAnalyticsCatalogAndJobEndpoint string `json:"azureDatalakeAnalyticsCatalogAndJobEndpoint"`
		AzureDataLakeStoreFileSystemEndpoint        string `json:"azureDatalakeStoreFileSystemEndpoint"`
		KeyVaultDNS                                 string `json:"keyvaultDns"`
		MariadbServerEndpoint                       string `json:"mariadbServerEndpoint"`
		MHSMDNS                                     string `json:"mhsmDns"`
		MysqlServerEndpoint                         string `json:"mysqlServerEndpoint"`
		PostgresqlServerEndpoint                    string `json:"postgresqlServerEndpoint"`
		SqlServerHostname                           string `json:"sqlServerHostname"`
		StorageEndpoint                             string `json:"storageEndpoint"`
		StorageSyncEndpoint                         string `json:"storageSyncEndpoint"`
		SynapseAnalyticsEndpoint                    string `json:"synapseAnalyticsEndpoint"`
	} `json:"suffixes"`
}

func FromCliName(environmentName string) (*Environment, error) {
	// check az-cli minimum version
	if err := azurecli.CheckAzVersion(azureCliMinimumVersion, nil); err != nil {
		return nil, err
	}

	var metadata azureCliSchema
	if err := azurecli.JSONUnmarshalAzCmd(&metadata, "cloud", "show", "--name", environmentName); err != nil {
		return nil, err
	}

	return &Environment{
		AzureADEndpoint: azureAdEndpoint(metadata.Endpoints.ActiveDirectory),

		AADGraph: Api{
			AppId:    PublishedApis["AzureActiveDirectoryGraph"],
			Endpoint: apiEndpoint(metadata.Endpoints.ActiveDirectoryGraphResourceID),
		},
		BatchManagement: Api{
			AppId:    PublishedApis["AzureBatch"],
			Endpoint: apiEndpoint(metadata.Endpoints.BatchResourceID),
		},
		DataLake: Api{
			AppId:    PublishedApis["AzureDataLake"],
			Endpoint: apiEndpoint(metadata.Endpoints.ActiveDirectoryDataLakeResourceID),
		},
		KeyVault: Api{
			AppId:    PublishedApis["AzureKeyVault"],
			Endpoint: apiEndpoint(fmt.Sprintf("https://%s", metadata.Suffixes.KeyVaultDNS)),
		},
		MSGraph: Api{
			AppId:    PublishedApis["MicrosoftGraph"],
			Endpoint: apiEndpoint(metadata.Endpoints.MicrosoftGraphResourceID),
		},
		OperationalInsights: Api{
			AppId:    PublishedApis["LogAnalytics"],
			Endpoint: apiEndpoint(metadata.Endpoints.LogAnalyticsResourceID),
		},
		OSSRDBMS: Api{
			AppId:    PublishedApis["OssRdbms"],
			Endpoint: apiEndpoint(metadata.Endpoints.OSSRDBMSResourceID),
		},
		ResourceManager: Api{
			AppId:    PublishedApis["AzureServiceManagement"],
			Endpoint: apiEndpoint(metadata.Endpoints.ResourceManager),
		},
		ServiceManagement: Api{
			AppId:    PublishedApis["AzureServiceManagement"],
			Endpoint: apiEndpoint(metadata.Endpoints.Management),
		},
		SQLDatabase: Api{
			AppId:    PublishedApis["AzureSqlDatabase"],
			Endpoint: apiEndpoint(fmt.Sprintf("https://%s", metadata.Endpoints.SqlManagement)),
		},
		Storage: Api{
			AppId:    PublishedApis["AzureStorage"],
			Endpoint: apiEndpoint(fmt.Sprintf("https://%s", metadata.Suffixes.StorageEndpoint)),
		},
		Synapse: Api{
			AppId:    PublishedApis["AzureSynapseGateway"],
			Endpoint: apiEndpoint(metadata.Endpoints.SynapseAnalyticsResourceID),
		},

		ServiceBus: ApiUnavailable,
	}, nil
}
