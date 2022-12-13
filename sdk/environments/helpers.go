package environments

import "github.com/hashicorp/go-azure-helpers/lang/pointer"

func applicationIdOnly(name, applicationId string) Api {
	return &ApiEndpoint{
		domainSuffix:        nil,
		endpoint:            nil,
		microsoftGraphAppId: applicationId,
		name:                name,
		resourceIdentifier:  nil,
	}
}

func ApiManagementAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: apiManagementAppId,
		name:                "ApiManagement",
		resourceIdentifier:  nil,
	}
}

func BatchAPI(endpoint string) Api {
	return &ApiEndpoint{
		domainSuffix:        nil,
		endpoint:            pointer.To(endpoint),
		microsoftGraphAppId: batchAppId,
		name:                "Batch",
		resourceIdentifier:  pointer.To("https://batch.core.windows.net/"),
	}
}

func ContainerRegistryAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: containerRegistryAppId,
		name:                "ContainerRegistry",
		resourceIdentifier:  pointer.To("https://api.loganalytics.io"),
	}
}

func CosmosDBAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: cosmosDBAppId,
		name:                "AzureCosmosDb",
		resourceIdentifier:  pointer.To("https://cosmos.azure.com"),
	}
}

func DataLakeAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: dataLakeAppId,
		name:                "DataLake",
		resourceIdentifier:  pointer.To("https://datalake.azure.net/"),
	}
}

func KeyVaultAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: keyVaultAppId,
		name:                "AzureKeyVault",
		resourceIdentifier:  pointer.To("https://vault.azure.net"),
	}
}

func ManagedHSMAPI(endpoint, domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            pointer.To(endpoint),
		microsoftGraphAppId: managedHSMAppId,
		name:                "ManagedHSM",
		resourceIdentifier:  pointer.To("https://managedhsm.azure.net"),
	}
}

func MariaDBAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: ossRDBMSAppId,
		name:                "OssRdbms",
		resourceIdentifier:  pointer.To("https://ossrdbms-aad.database.windows.net"),
	}
}

func MicrosoftGraphAPI(endpoint string) Api {
	return &ApiEndpoint{
		domainSuffix:        nil,
		endpoint:            pointer.To(endpoint),
		microsoftGraphAppId: microsoftGraphAppId,
		name:                "MicrosoftGraph",
		resourceIdentifier:  pointer.To("https://graph.microsoft.com/"),
	}
}

func MySqlAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: ossRDBMSAppId,
		name:                "OssRdbms",
		resourceIdentifier:  pointer.To("https://ossrdbms-aad.database.windows.net"),
	}
}

func OperationalInsightsAPI() Api {
	return &ApiEndpoint{
		domainSuffix:        nil,
		endpoint:            nil,
		microsoftGraphAppId: logAnalyticsAppId,
		name:                "OperationalInsights",
		resourceIdentifier:  pointer.To("https://api.loganalytics.io"),
	}
}

func PostgresqlAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: ossRDBMSAppId,
		name:                "OssRdbms",
		resourceIdentifier:  pointer.To("https://ossrdbms-aad.database.windows.net"),
	}
}

func ResourceManagerAPI(endpoint string) Api {
	return &ApiEndpoint{
		domainSuffix:        nil,
		endpoint:            pointer.To(endpoint),
		microsoftGraphAppId: "TODO", // TODO: impl. me
		name:                "ResourceManager",
		resourceIdentifier:  nil,
	}
}

func ServiceBusAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: serviceBusAppId,
		name:                "ServiceBus",
		resourceIdentifier:  pointer.To("https://servicebus.azure.net/"),
	}
}

func SqlAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: sqlDatabaseAppId,
		name:                "AzureSqlDatabase",
		resourceIdentifier:  pointer.To("https://database.windows.net/"),
	}
}

func StorageAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: storageAppId,
		name:                "AzureStorage",
		resourceIdentifier:  pointer.To("https://storage.azure.com/"),
	}
}

func SynapseAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: synapseAppId,
		name:                "Synapse",
		resourceIdentifier:  pointer.To("https://dev.azuresynapse.net"),
	}
}

func TrafficManagerAPI(domainSuffix string) Api {
	return &ApiEndpoint{
		domainSuffix:        pointer.To(domainSuffix),
		endpoint:            nil,
		microsoftGraphAppId: trafficManagerAppId,
		name:                "TrafficManager",
		resourceIdentifier:  nil,
	}
}
