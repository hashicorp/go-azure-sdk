package environments

/*
	Name:                         "AzurePublicCloud",
// TODO: we probly need these two as well?!
		ResourceManagerEndpoint:      "https://management.azure.com/",
		ActiveDirectoryEndpoint:      "https://login.microsoftonline.com/",

		// TODO: then what about these?
		TokenAudience:                "https://management.azure.com/",
		GalleryEndpoint:              "https://gallery.azure.com/",
		ResourceManagerVMDNSSuffix:   "cloudapp.azure.com",

		KeyVaultEndpoint:             "https://vault.azure.net/",
		MicrosoftGraphEndpoint:       "https://graph.microsoft.com/",
*/

func AzurePublic() Environment {
	env := baseEnvironmentWithName("Public")

	env.Authorization = &Authorization{
		Audiences: []string{
			"https://management.core.windows.net/",
			"https://management.azure.com/",
		},
		LoginEndpoint: "https://login.microsoftonline.com",
	}
	env.ResourceManager = ResourceManagerAPI("https://management.azure.com")
	env.MicrosoftGraph = MicrosoftGraphAPI("https://graph.microsoft.com/")

	env.ApiManagement = ApiManagementAPI("azure-api.net")
	env.Batch = BatchAPI("https://batch.core.windows.net/")
	env.ContainerRegistry = ContainerRegistryAPI("azurecr.io")
	env.CosmosDB = CosmosDBAPI("documents.azure.com")
	env.DataLake = DataLakeAPI("azuredatalakestore.net")
	env.KeyVault = KeyVaultAPI("vault.azure.net")
	env.ManagedHSM = ManagedHSMAPI("https://managedhsm.azure.net", "managedhsm.azure.net")
	env.MariaDB = MariaDBAPI("mariadb.database.azure.com")
	env.MySql = MySqlAPI("mysql.database.azure.com")
	env.OperationalInsights = OperationalInsightsAPI()
	env.Postgresql = PostgresqlAPI("postgres.database.azure.com")
	env.ServiceBus = ServiceBusAPI("servicebus.windows.net")
	env.Sql = SqlAPI("database.windows.net")
	env.Storage = StorageAPI("core.windows.net")
	env.Synapse = SynapseAPI("dev.azuresynapse.net")
	env.TrafficManager = TrafficManagerAPI("trafficmanager.net")

	return env
}
