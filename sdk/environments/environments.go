package environments

import (
	"fmt"
	"strings"
)

// Environment represents a set of API configurations for a particular cloud.
type Environment struct {
	// The Azure AD endpoint for acquiring access tokens.
	AzureADEndpoint AzureADEndpoint

	// API definitions
	MSGraph             Api
	AADGraph            Api
	ResourceManager     Api
	BatchManagement     Api
	DataLake            Api
	KeyVault            Api
	OperationalInsights Api
	OSSRDBMS            Api
	ServiceBus          Api
	ServiceManagement   Api
	SQLDatabase         Api
	Storage             Api
	Synapse             Api
}

var (
	Global = Environment{
		AzureADEndpoint:     AzureADGlobal,
		MSGraph:             MsGraphGlobal,
		AADGraph:            AadGraphGlobal,
		ResourceManager:     ResourceManagerPublic,
		BatchManagement:     BatchManagementPublic,
		DataLake:            DataLakePublic,
		KeyVault:            KeyVaultPublic,
		OperationalInsights: OperationalInsightsPublic,
		OSSRDBMS:            OSSRDBMSPublic,
		ServiceBus:          ServiceBusPublic,
		ServiceManagement:   ServiceManagementPublic,
		SQLDatabase:         SQLDatabasePublic,
		Storage:             StoragePublic,
		Synapse:             SynapsePublic,
	}

	Public = Global

	China = Environment{
		AzureADEndpoint:     AzureADChina,
		MSGraph:             MsGraphChina,
		AADGraph:            AadGraphChina,
		ResourceManager:     ResourceManagerChina,
		BatchManagement:     BatchManagementChina,
		KeyVault:            KeyVaultChina,
		OperationalInsights: ApiUnavailable,
		OSSRDBMS:            OSSRDBMSChina,
		ServiceBus:          ServiceBusChina,
		ServiceManagement:   ServiceManagementChina,
		SQLDatabase:         SQLDatabaseChina,
		Storage:             StoragePublic,
		Synapse:             SynapsePublic,
	}

	USGovernmentL4 = Environment{
		AzureADEndpoint:     AzureADUSGov,
		MSGraph:             MsGraphUSGovL4,
		AADGraph:            AadGraphUSGov,
		ResourceManager:     ResourceManagerUSGov,
		BatchManagement:     BatchManagementUSGov,
		DataLake:            ApiUnavailable,
		KeyVault:            KeyVaultUSGov,
		OperationalInsights: OperationalInsightsUSGov,
		OSSRDBMS:            OSSRDBMSUSGov,
		ServiceBus:          ServiceBusUSGov,
		ServiceManagement:   ServiceManagementUSGov,
		SQLDatabase:         SQLDatabaseUSGov,
		Storage:             StoragePublic,
		Synapse:             ApiUnavailable,
	}

	USGovernmentL5 = Environment{
		AzureADEndpoint:     AzureADUSGov,
		MSGraph:             MsGraphUSGovL5,
		AADGraph:            AadGraphUSGov,
		ResourceManager:     ResourceManagerUSGov,
		BatchManagement:     BatchManagementUSGov,
		DataLake:            ApiUnavailable,
		KeyVault:            KeyVaultUSGov,
		OperationalInsights: OperationalInsightsUSGov,
		OSSRDBMS:            OSSRDBMSUSGov,
		ServiceBus:          ServiceBusUSGov,
		ServiceManagement:   ServiceManagementUSGov,
		SQLDatabase:         SQLDatabaseUSGov,
		Storage:             StoragePublic,
		Synapse:             ApiUnavailable,
	}

	Canary = Environment{
		AzureADEndpoint: AzureADGlobal,
		MSGraph:         MsGraphCanary,
	}
)

func FromNamed(env string) (Environment, error) {
	switch strings.ToLower(env) {
	case "public", "global":
		return Global, nil
	case "usgovernment", "usgovernmentl4":
		return USGovernmentL4, nil
	case "dod", "usgovernmentl5":
		return USGovernmentL5, nil
	case "canary":
		return Canary, nil
	case "china":
		return China, nil
	}

	return Environment{}, fmt.Errorf("invalid environment specified: %s", env)
}
