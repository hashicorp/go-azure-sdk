package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllowedContentLevel string

const (
	AllowedContentLevelHigh   AllowedContentLevel = "High"
	AllowedContentLevelLow    AllowedContentLevel = "Low"
	AllowedContentLevelMedium AllowedContentLevel = "Medium"
)

func PossibleValuesForAllowedContentLevel() []string {
	return []string{
		string(AllowedContentLevelHigh),
		string(AllowedContentLevelLow),
		string(AllowedContentLevelMedium),
	}
}

func (s *AllowedContentLevel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllowedContentLevel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllowedContentLevel(input string) (*AllowedContentLevel, error) {
	vals := map[string]AllowedContentLevel{
		"high":   AllowedContentLevelHigh,
		"low":    AllowedContentLevelLow,
		"medium": AllowedContentLevelMedium,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllowedContentLevel(input)
	return &out, nil
}

type ConnectionAuthType string

const (
	ConnectionAuthTypeAAD              ConnectionAuthType = "AAD"
	ConnectionAuthTypeAccessKey        ConnectionAuthType = "AccessKey"
	ConnectionAuthTypeAccountKey       ConnectionAuthType = "AccountKey"
	ConnectionAuthTypeApiKey           ConnectionAuthType = "ApiKey"
	ConnectionAuthTypeCustomKeys       ConnectionAuthType = "CustomKeys"
	ConnectionAuthTypeManagedIdentity  ConnectionAuthType = "ManagedIdentity"
	ConnectionAuthTypeNone             ConnectionAuthType = "None"
	ConnectionAuthTypeOAuthTwo         ConnectionAuthType = "OAuth2"
	ConnectionAuthTypePAT              ConnectionAuthType = "PAT"
	ConnectionAuthTypeSAS              ConnectionAuthType = "SAS"
	ConnectionAuthTypeServicePrincipal ConnectionAuthType = "ServicePrincipal"
	ConnectionAuthTypeUsernamePassword ConnectionAuthType = "UsernamePassword"
)

func PossibleValuesForConnectionAuthType() []string {
	return []string{
		string(ConnectionAuthTypeAAD),
		string(ConnectionAuthTypeAccessKey),
		string(ConnectionAuthTypeAccountKey),
		string(ConnectionAuthTypeApiKey),
		string(ConnectionAuthTypeCustomKeys),
		string(ConnectionAuthTypeManagedIdentity),
		string(ConnectionAuthTypeNone),
		string(ConnectionAuthTypeOAuthTwo),
		string(ConnectionAuthTypePAT),
		string(ConnectionAuthTypeSAS),
		string(ConnectionAuthTypeServicePrincipal),
		string(ConnectionAuthTypeUsernamePassword),
	}
}

func (s *ConnectionAuthType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectionAuthType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectionAuthType(input string) (*ConnectionAuthType, error) {
	vals := map[string]ConnectionAuthType{
		"aad":              ConnectionAuthTypeAAD,
		"accesskey":        ConnectionAuthTypeAccessKey,
		"accountkey":       ConnectionAuthTypeAccountKey,
		"apikey":           ConnectionAuthTypeApiKey,
		"customkeys":       ConnectionAuthTypeCustomKeys,
		"managedidentity":  ConnectionAuthTypeManagedIdentity,
		"none":             ConnectionAuthTypeNone,
		"oauth2":           ConnectionAuthTypeOAuthTwo,
		"pat":              ConnectionAuthTypePAT,
		"sas":              ConnectionAuthTypeSAS,
		"serviceprincipal": ConnectionAuthTypeServicePrincipal,
		"usernamepassword": ConnectionAuthTypeUsernamePassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionAuthType(input)
	return &out, nil
}

type ConnectionCategory string

const (
	ConnectionCategoryADLSGenTwo               ConnectionCategory = "ADLSGen2"
	ConnectionCategoryAIServices               ConnectionCategory = "AIServices"
	ConnectionCategoryAmazonMws                ConnectionCategory = "AmazonMws"
	ConnectionCategoryAmazonRdsForOracle       ConnectionCategory = "AmazonRdsForOracle"
	ConnectionCategoryAmazonRdsForSqlServer    ConnectionCategory = "AmazonRdsForSqlServer"
	ConnectionCategoryAmazonRedshift           ConnectionCategory = "AmazonRedshift"
	ConnectionCategoryAmazonSThreeCompatible   ConnectionCategory = "AmazonS3Compatible"
	ConnectionCategoryApiKey                   ConnectionCategory = "ApiKey"
	ConnectionCategoryAzureBlob                ConnectionCategory = "AzureBlob"
	ConnectionCategoryAzureDataExplorer        ConnectionCategory = "AzureDataExplorer"
	ConnectionCategoryAzureDatabricksDeltaLake ConnectionCategory = "AzureDatabricksDeltaLake"
	ConnectionCategoryAzureMariaDb             ConnectionCategory = "AzureMariaDb"
	ConnectionCategoryAzureMySqlDb             ConnectionCategory = "AzureMySqlDb"
	ConnectionCategoryAzureOneLake             ConnectionCategory = "AzureOneLake"
	ConnectionCategoryAzureOpenAI              ConnectionCategory = "AzureOpenAI"
	ConnectionCategoryAzurePostgresDb          ConnectionCategory = "AzurePostgresDb"
	ConnectionCategoryAzureSqlDb               ConnectionCategory = "AzureSqlDb"
	ConnectionCategoryAzureSqlMi               ConnectionCategory = "AzureSqlMi"
	ConnectionCategoryAzureSynapseAnalytics    ConnectionCategory = "AzureSynapseAnalytics"
	ConnectionCategoryAzureTableStorage        ConnectionCategory = "AzureTableStorage"
	ConnectionCategoryBingLLMSearch            ConnectionCategory = "BingLLMSearch"
	ConnectionCategoryCassandra                ConnectionCategory = "Cassandra"
	ConnectionCategoryCognitiveSearch          ConnectionCategory = "CognitiveSearch"
	ConnectionCategoryCognitiveService         ConnectionCategory = "CognitiveService"
	ConnectionCategoryConcur                   ConnectionCategory = "Concur"
	ConnectionCategoryContainerRegistry        ConnectionCategory = "ContainerRegistry"
	ConnectionCategoryCosmosDb                 ConnectionCategory = "CosmosDb"
	ConnectionCategoryCosmosDbMongoDbApi       ConnectionCategory = "CosmosDbMongoDbApi"
	ConnectionCategoryCouchbase                ConnectionCategory = "Couchbase"
	ConnectionCategoryCustomKeys               ConnectionCategory = "CustomKeys"
	ConnectionCategoryDbTwo                    ConnectionCategory = "Db2"
	ConnectionCategoryDrill                    ConnectionCategory = "Drill"
	ConnectionCategoryDynamics                 ConnectionCategory = "Dynamics"
	ConnectionCategoryDynamicsAx               ConnectionCategory = "DynamicsAx"
	ConnectionCategoryDynamicsCrm              ConnectionCategory = "DynamicsCrm"
	ConnectionCategoryElasticsearch            ConnectionCategory = "Elasticsearch"
	ConnectionCategoryEloqua                   ConnectionCategory = "Eloqua"
	ConnectionCategoryFileServer               ConnectionCategory = "FileServer"
	ConnectionCategoryFtpServer                ConnectionCategory = "FtpServer"
	ConnectionCategoryGenericContainerRegistry ConnectionCategory = "GenericContainerRegistry"
	ConnectionCategoryGenericHTTP              ConnectionCategory = "GenericHttp"
	ConnectionCategoryGenericRest              ConnectionCategory = "GenericRest"
	ConnectionCategoryGit                      ConnectionCategory = "Git"
	ConnectionCategoryGoogleAdWords            ConnectionCategory = "GoogleAdWords"
	ConnectionCategoryGoogleBigQuery           ConnectionCategory = "GoogleBigQuery"
	ConnectionCategoryGoogleCloudStorage       ConnectionCategory = "GoogleCloudStorage"
	ConnectionCategoryGreenplum                ConnectionCategory = "Greenplum"
	ConnectionCategoryHbase                    ConnectionCategory = "Hbase"
	ConnectionCategoryHdfs                     ConnectionCategory = "Hdfs"
	ConnectionCategoryHive                     ConnectionCategory = "Hive"
	ConnectionCategoryHubspot                  ConnectionCategory = "Hubspot"
	ConnectionCategoryImpala                   ConnectionCategory = "Impala"
	ConnectionCategoryInformix                 ConnectionCategory = "Informix"
	ConnectionCategoryJira                     ConnectionCategory = "Jira"
	ConnectionCategoryMagento                  ConnectionCategory = "Magento"
	ConnectionCategoryManagedOnlineEndpoint    ConnectionCategory = "ManagedOnlineEndpoint"
	ConnectionCategoryMariaDb                  ConnectionCategory = "MariaDb"
	ConnectionCategoryMarketo                  ConnectionCategory = "Marketo"
	ConnectionCategoryMicrosoftAccess          ConnectionCategory = "MicrosoftAccess"
	ConnectionCategoryMongoDbAtlas             ConnectionCategory = "MongoDbAtlas"
	ConnectionCategoryMongoDbVTwo              ConnectionCategory = "MongoDbV2"
	ConnectionCategoryMySql                    ConnectionCategory = "MySql"
	ConnectionCategoryNetezza                  ConnectionCategory = "Netezza"
	ConnectionCategoryODataRest                ConnectionCategory = "ODataRest"
	ConnectionCategoryOdbc                     ConnectionCategory = "Odbc"
	ConnectionCategoryOfficeThreeSixFive       ConnectionCategory = "Office365"
	ConnectionCategoryOpenAI                   ConnectionCategory = "OpenAI"
	ConnectionCategoryOracle                   ConnectionCategory = "Oracle"
	ConnectionCategoryOracleCloudStorage       ConnectionCategory = "OracleCloudStorage"
	ConnectionCategoryOracleServiceCloud       ConnectionCategory = "OracleServiceCloud"
	ConnectionCategoryPayPal                   ConnectionCategory = "PayPal"
	ConnectionCategoryPhoenix                  ConnectionCategory = "Phoenix"
	ConnectionCategoryPinecone                 ConnectionCategory = "Pinecone"
	ConnectionCategoryPostgreSql               ConnectionCategory = "PostgreSql"
	ConnectionCategoryPresto                   ConnectionCategory = "Presto"
	ConnectionCategoryPythonFeed               ConnectionCategory = "PythonFeed"
	ConnectionCategoryQuickBooks               ConnectionCategory = "QuickBooks"
	ConnectionCategoryRedis                    ConnectionCategory = "Redis"
	ConnectionCategoryResponsys                ConnectionCategory = "Responsys"
	ConnectionCategorySThree                   ConnectionCategory = "S3"
	ConnectionCategorySalesforce               ConnectionCategory = "Salesforce"
	ConnectionCategorySalesforceMarketingCloud ConnectionCategory = "SalesforceMarketingCloud"
	ConnectionCategorySalesforceServiceCloud   ConnectionCategory = "SalesforceServiceCloud"
	ConnectionCategorySapBw                    ConnectionCategory = "SapBw"
	ConnectionCategorySapCloudForCustomer      ConnectionCategory = "SapCloudForCustomer"
	ConnectionCategorySapEcc                   ConnectionCategory = "SapEcc"
	ConnectionCategorySapHana                  ConnectionCategory = "SapHana"
	ConnectionCategorySapOpenHub               ConnectionCategory = "SapOpenHub"
	ConnectionCategorySapTable                 ConnectionCategory = "SapTable"
	ConnectionCategorySerp                     ConnectionCategory = "Serp"
	ConnectionCategoryServerless               ConnectionCategory = "Serverless"
	ConnectionCategoryServiceNow               ConnectionCategory = "ServiceNow"
	ConnectionCategorySftp                     ConnectionCategory = "Sftp"
	ConnectionCategorySharePointOnlineList     ConnectionCategory = "SharePointOnlineList"
	ConnectionCategoryShopify                  ConnectionCategory = "Shopify"
	ConnectionCategorySnowflake                ConnectionCategory = "Snowflake"
	ConnectionCategorySpark                    ConnectionCategory = "Spark"
	ConnectionCategorySqlServer                ConnectionCategory = "SqlServer"
	ConnectionCategorySquare                   ConnectionCategory = "Square"
	ConnectionCategorySybase                   ConnectionCategory = "Sybase"
	ConnectionCategoryTeradata                 ConnectionCategory = "Teradata"
	ConnectionCategoryVertica                  ConnectionCategory = "Vertica"
	ConnectionCategoryWebTable                 ConnectionCategory = "WebTable"
	ConnectionCategoryXero                     ConnectionCategory = "Xero"
	ConnectionCategoryZoho                     ConnectionCategory = "Zoho"
)

func PossibleValuesForConnectionCategory() []string {
	return []string{
		string(ConnectionCategoryADLSGenTwo),
		string(ConnectionCategoryAIServices),
		string(ConnectionCategoryAmazonMws),
		string(ConnectionCategoryAmazonRdsForOracle),
		string(ConnectionCategoryAmazonRdsForSqlServer),
		string(ConnectionCategoryAmazonRedshift),
		string(ConnectionCategoryAmazonSThreeCompatible),
		string(ConnectionCategoryApiKey),
		string(ConnectionCategoryAzureBlob),
		string(ConnectionCategoryAzureDataExplorer),
		string(ConnectionCategoryAzureDatabricksDeltaLake),
		string(ConnectionCategoryAzureMariaDb),
		string(ConnectionCategoryAzureMySqlDb),
		string(ConnectionCategoryAzureOneLake),
		string(ConnectionCategoryAzureOpenAI),
		string(ConnectionCategoryAzurePostgresDb),
		string(ConnectionCategoryAzureSqlDb),
		string(ConnectionCategoryAzureSqlMi),
		string(ConnectionCategoryAzureSynapseAnalytics),
		string(ConnectionCategoryAzureTableStorage),
		string(ConnectionCategoryBingLLMSearch),
		string(ConnectionCategoryCassandra),
		string(ConnectionCategoryCognitiveSearch),
		string(ConnectionCategoryCognitiveService),
		string(ConnectionCategoryConcur),
		string(ConnectionCategoryContainerRegistry),
		string(ConnectionCategoryCosmosDb),
		string(ConnectionCategoryCosmosDbMongoDbApi),
		string(ConnectionCategoryCouchbase),
		string(ConnectionCategoryCustomKeys),
		string(ConnectionCategoryDbTwo),
		string(ConnectionCategoryDrill),
		string(ConnectionCategoryDynamics),
		string(ConnectionCategoryDynamicsAx),
		string(ConnectionCategoryDynamicsCrm),
		string(ConnectionCategoryElasticsearch),
		string(ConnectionCategoryEloqua),
		string(ConnectionCategoryFileServer),
		string(ConnectionCategoryFtpServer),
		string(ConnectionCategoryGenericContainerRegistry),
		string(ConnectionCategoryGenericHTTP),
		string(ConnectionCategoryGenericRest),
		string(ConnectionCategoryGit),
		string(ConnectionCategoryGoogleAdWords),
		string(ConnectionCategoryGoogleBigQuery),
		string(ConnectionCategoryGoogleCloudStorage),
		string(ConnectionCategoryGreenplum),
		string(ConnectionCategoryHbase),
		string(ConnectionCategoryHdfs),
		string(ConnectionCategoryHive),
		string(ConnectionCategoryHubspot),
		string(ConnectionCategoryImpala),
		string(ConnectionCategoryInformix),
		string(ConnectionCategoryJira),
		string(ConnectionCategoryMagento),
		string(ConnectionCategoryManagedOnlineEndpoint),
		string(ConnectionCategoryMariaDb),
		string(ConnectionCategoryMarketo),
		string(ConnectionCategoryMicrosoftAccess),
		string(ConnectionCategoryMongoDbAtlas),
		string(ConnectionCategoryMongoDbVTwo),
		string(ConnectionCategoryMySql),
		string(ConnectionCategoryNetezza),
		string(ConnectionCategoryODataRest),
		string(ConnectionCategoryOdbc),
		string(ConnectionCategoryOfficeThreeSixFive),
		string(ConnectionCategoryOpenAI),
		string(ConnectionCategoryOracle),
		string(ConnectionCategoryOracleCloudStorage),
		string(ConnectionCategoryOracleServiceCloud),
		string(ConnectionCategoryPayPal),
		string(ConnectionCategoryPhoenix),
		string(ConnectionCategoryPinecone),
		string(ConnectionCategoryPostgreSql),
		string(ConnectionCategoryPresto),
		string(ConnectionCategoryPythonFeed),
		string(ConnectionCategoryQuickBooks),
		string(ConnectionCategoryRedis),
		string(ConnectionCategoryResponsys),
		string(ConnectionCategorySThree),
		string(ConnectionCategorySalesforce),
		string(ConnectionCategorySalesforceMarketingCloud),
		string(ConnectionCategorySalesforceServiceCloud),
		string(ConnectionCategorySapBw),
		string(ConnectionCategorySapCloudForCustomer),
		string(ConnectionCategorySapEcc),
		string(ConnectionCategorySapHana),
		string(ConnectionCategorySapOpenHub),
		string(ConnectionCategorySapTable),
		string(ConnectionCategorySerp),
		string(ConnectionCategoryServerless),
		string(ConnectionCategoryServiceNow),
		string(ConnectionCategorySftp),
		string(ConnectionCategorySharePointOnlineList),
		string(ConnectionCategoryShopify),
		string(ConnectionCategorySnowflake),
		string(ConnectionCategorySpark),
		string(ConnectionCategorySqlServer),
		string(ConnectionCategorySquare),
		string(ConnectionCategorySybase),
		string(ConnectionCategoryTeradata),
		string(ConnectionCategoryVertica),
		string(ConnectionCategoryWebTable),
		string(ConnectionCategoryXero),
		string(ConnectionCategoryZoho),
	}
}

func (s *ConnectionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectionCategory(input string) (*ConnectionCategory, error) {
	vals := map[string]ConnectionCategory{
		"adlsgen2":                 ConnectionCategoryADLSGenTwo,
		"aiservices":               ConnectionCategoryAIServices,
		"amazonmws":                ConnectionCategoryAmazonMws,
		"amazonrdsfororacle":       ConnectionCategoryAmazonRdsForOracle,
		"amazonrdsforsqlserver":    ConnectionCategoryAmazonRdsForSqlServer,
		"amazonredshift":           ConnectionCategoryAmazonRedshift,
		"amazons3compatible":       ConnectionCategoryAmazonSThreeCompatible,
		"apikey":                   ConnectionCategoryApiKey,
		"azureblob":                ConnectionCategoryAzureBlob,
		"azuredataexplorer":        ConnectionCategoryAzureDataExplorer,
		"azuredatabricksdeltalake": ConnectionCategoryAzureDatabricksDeltaLake,
		"azuremariadb":             ConnectionCategoryAzureMariaDb,
		"azuremysqldb":             ConnectionCategoryAzureMySqlDb,
		"azureonelake":             ConnectionCategoryAzureOneLake,
		"azureopenai":              ConnectionCategoryAzureOpenAI,
		"azurepostgresdb":          ConnectionCategoryAzurePostgresDb,
		"azuresqldb":               ConnectionCategoryAzureSqlDb,
		"azuresqlmi":               ConnectionCategoryAzureSqlMi,
		"azuresynapseanalytics":    ConnectionCategoryAzureSynapseAnalytics,
		"azuretablestorage":        ConnectionCategoryAzureTableStorage,
		"bingllmsearch":            ConnectionCategoryBingLLMSearch,
		"cassandra":                ConnectionCategoryCassandra,
		"cognitivesearch":          ConnectionCategoryCognitiveSearch,
		"cognitiveservice":         ConnectionCategoryCognitiveService,
		"concur":                   ConnectionCategoryConcur,
		"containerregistry":        ConnectionCategoryContainerRegistry,
		"cosmosdb":                 ConnectionCategoryCosmosDb,
		"cosmosdbmongodbapi":       ConnectionCategoryCosmosDbMongoDbApi,
		"couchbase":                ConnectionCategoryCouchbase,
		"customkeys":               ConnectionCategoryCustomKeys,
		"db2":                      ConnectionCategoryDbTwo,
		"drill":                    ConnectionCategoryDrill,
		"dynamics":                 ConnectionCategoryDynamics,
		"dynamicsax":               ConnectionCategoryDynamicsAx,
		"dynamicscrm":              ConnectionCategoryDynamicsCrm,
		"elasticsearch":            ConnectionCategoryElasticsearch,
		"eloqua":                   ConnectionCategoryEloqua,
		"fileserver":               ConnectionCategoryFileServer,
		"ftpserver":                ConnectionCategoryFtpServer,
		"genericcontainerregistry": ConnectionCategoryGenericContainerRegistry,
		"generichttp":              ConnectionCategoryGenericHTTP,
		"genericrest":              ConnectionCategoryGenericRest,
		"git":                      ConnectionCategoryGit,
		"googleadwords":            ConnectionCategoryGoogleAdWords,
		"googlebigquery":           ConnectionCategoryGoogleBigQuery,
		"googlecloudstorage":       ConnectionCategoryGoogleCloudStorage,
		"greenplum":                ConnectionCategoryGreenplum,
		"hbase":                    ConnectionCategoryHbase,
		"hdfs":                     ConnectionCategoryHdfs,
		"hive":                     ConnectionCategoryHive,
		"hubspot":                  ConnectionCategoryHubspot,
		"impala":                   ConnectionCategoryImpala,
		"informix":                 ConnectionCategoryInformix,
		"jira":                     ConnectionCategoryJira,
		"magento":                  ConnectionCategoryMagento,
		"managedonlineendpoint":    ConnectionCategoryManagedOnlineEndpoint,
		"mariadb":                  ConnectionCategoryMariaDb,
		"marketo":                  ConnectionCategoryMarketo,
		"microsoftaccess":          ConnectionCategoryMicrosoftAccess,
		"mongodbatlas":             ConnectionCategoryMongoDbAtlas,
		"mongodbv2":                ConnectionCategoryMongoDbVTwo,
		"mysql":                    ConnectionCategoryMySql,
		"netezza":                  ConnectionCategoryNetezza,
		"odatarest":                ConnectionCategoryODataRest,
		"odbc":                     ConnectionCategoryOdbc,
		"office365":                ConnectionCategoryOfficeThreeSixFive,
		"openai":                   ConnectionCategoryOpenAI,
		"oracle":                   ConnectionCategoryOracle,
		"oraclecloudstorage":       ConnectionCategoryOracleCloudStorage,
		"oracleservicecloud":       ConnectionCategoryOracleServiceCloud,
		"paypal":                   ConnectionCategoryPayPal,
		"phoenix":                  ConnectionCategoryPhoenix,
		"pinecone":                 ConnectionCategoryPinecone,
		"postgresql":               ConnectionCategoryPostgreSql,
		"presto":                   ConnectionCategoryPresto,
		"pythonfeed":               ConnectionCategoryPythonFeed,
		"quickbooks":               ConnectionCategoryQuickBooks,
		"redis":                    ConnectionCategoryRedis,
		"responsys":                ConnectionCategoryResponsys,
		"s3":                       ConnectionCategorySThree,
		"salesforce":               ConnectionCategorySalesforce,
		"salesforcemarketingcloud": ConnectionCategorySalesforceMarketingCloud,
		"salesforceservicecloud":   ConnectionCategorySalesforceServiceCloud,
		"sapbw":                    ConnectionCategorySapBw,
		"sapcloudforcustomer":      ConnectionCategorySapCloudForCustomer,
		"sapecc":                   ConnectionCategorySapEcc,
		"saphana":                  ConnectionCategorySapHana,
		"sapopenhub":               ConnectionCategorySapOpenHub,
		"saptable":                 ConnectionCategorySapTable,
		"serp":                     ConnectionCategorySerp,
		"serverless":               ConnectionCategoryServerless,
		"servicenow":               ConnectionCategoryServiceNow,
		"sftp":                     ConnectionCategorySftp,
		"sharepointonlinelist":     ConnectionCategorySharePointOnlineList,
		"shopify":                  ConnectionCategoryShopify,
		"snowflake":                ConnectionCategorySnowflake,
		"spark":                    ConnectionCategorySpark,
		"sqlserver":                ConnectionCategorySqlServer,
		"square":                   ConnectionCategorySquare,
		"sybase":                   ConnectionCategorySybase,
		"teradata":                 ConnectionCategoryTeradata,
		"vertica":                  ConnectionCategoryVertica,
		"webtable":                 ConnectionCategoryWebTable,
		"xero":                     ConnectionCategoryXero,
		"zoho":                     ConnectionCategoryZoho,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionCategory(input)
	return &out, nil
}

type ConnectionGroup string

const (
	ConnectionGroupAzure           ConnectionGroup = "Azure"
	ConnectionGroupAzureAI         ConnectionGroup = "AzureAI"
	ConnectionGroupDatabase        ConnectionGroup = "Database"
	ConnectionGroupFile            ConnectionGroup = "File"
	ConnectionGroupGenericProtocol ConnectionGroup = "GenericProtocol"
	ConnectionGroupNoSQL           ConnectionGroup = "NoSQL"
	ConnectionGroupServicesAndApps ConnectionGroup = "ServicesAndApps"
)

func PossibleValuesForConnectionGroup() []string {
	return []string{
		string(ConnectionGroupAzure),
		string(ConnectionGroupAzureAI),
		string(ConnectionGroupDatabase),
		string(ConnectionGroupFile),
		string(ConnectionGroupGenericProtocol),
		string(ConnectionGroupNoSQL),
		string(ConnectionGroupServicesAndApps),
	}
}

func (s *ConnectionGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectionGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectionGroup(input string) (*ConnectionGroup, error) {
	vals := map[string]ConnectionGroup{
		"azure":           ConnectionGroupAzure,
		"azureai":         ConnectionGroupAzureAI,
		"database":        ConnectionGroupDatabase,
		"file":            ConnectionGroupFile,
		"genericprotocol": ConnectionGroupGenericProtocol,
		"nosql":           ConnectionGroupNoSQL,
		"servicesandapps": ConnectionGroupServicesAndApps,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionGroup(input)
	return &out, nil
}

type DefaultResourceProvisioningState string

const (
	DefaultResourceProvisioningStateAccepted   DefaultResourceProvisioningState = "Accepted"
	DefaultResourceProvisioningStateCanceled   DefaultResourceProvisioningState = "Canceled"
	DefaultResourceProvisioningStateCreating   DefaultResourceProvisioningState = "Creating"
	DefaultResourceProvisioningStateDeleting   DefaultResourceProvisioningState = "Deleting"
	DefaultResourceProvisioningStateDisabled   DefaultResourceProvisioningState = "Disabled"
	DefaultResourceProvisioningStateFailed     DefaultResourceProvisioningState = "Failed"
	DefaultResourceProvisioningStateNotStarted DefaultResourceProvisioningState = "NotStarted"
	DefaultResourceProvisioningStateScaling    DefaultResourceProvisioningState = "Scaling"
	DefaultResourceProvisioningStateSucceeded  DefaultResourceProvisioningState = "Succeeded"
	DefaultResourceProvisioningStateUpdating   DefaultResourceProvisioningState = "Updating"
)

func PossibleValuesForDefaultResourceProvisioningState() []string {
	return []string{
		string(DefaultResourceProvisioningStateAccepted),
		string(DefaultResourceProvisioningStateCanceled),
		string(DefaultResourceProvisioningStateCreating),
		string(DefaultResourceProvisioningStateDeleting),
		string(DefaultResourceProvisioningStateDisabled),
		string(DefaultResourceProvisioningStateFailed),
		string(DefaultResourceProvisioningStateNotStarted),
		string(DefaultResourceProvisioningStateScaling),
		string(DefaultResourceProvisioningStateSucceeded),
		string(DefaultResourceProvisioningStateUpdating),
	}
}

func (s *DefaultResourceProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultResourceProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultResourceProvisioningState(input string) (*DefaultResourceProvisioningState, error) {
	vals := map[string]DefaultResourceProvisioningState{
		"accepted":   DefaultResourceProvisioningStateAccepted,
		"canceled":   DefaultResourceProvisioningStateCanceled,
		"creating":   DefaultResourceProvisioningStateCreating,
		"deleting":   DefaultResourceProvisioningStateDeleting,
		"disabled":   DefaultResourceProvisioningStateDisabled,
		"failed":     DefaultResourceProvisioningStateFailed,
		"notstarted": DefaultResourceProvisioningStateNotStarted,
		"scaling":    DefaultResourceProvisioningStateScaling,
		"succeeded":  DefaultResourceProvisioningStateSucceeded,
		"updating":   DefaultResourceProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultResourceProvisioningState(input)
	return &out, nil
}

type DeploymentModelVersionUpgradeOption string

const (
	DeploymentModelVersionUpgradeOptionNoAutoUpgrade                  DeploymentModelVersionUpgradeOption = "NoAutoUpgrade"
	DeploymentModelVersionUpgradeOptionOnceCurrentVersionExpired      DeploymentModelVersionUpgradeOption = "OnceCurrentVersionExpired"
	DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable DeploymentModelVersionUpgradeOption = "OnceNewDefaultVersionAvailable"
)

func PossibleValuesForDeploymentModelVersionUpgradeOption() []string {
	return []string{
		string(DeploymentModelVersionUpgradeOptionNoAutoUpgrade),
		string(DeploymentModelVersionUpgradeOptionOnceCurrentVersionExpired),
		string(DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable),
	}
}

func (s *DeploymentModelVersionUpgradeOption) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeploymentModelVersionUpgradeOption(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeploymentModelVersionUpgradeOption(input string) (*DeploymentModelVersionUpgradeOption, error) {
	vals := map[string]DeploymentModelVersionUpgradeOption{
		"noautoupgrade":                  DeploymentModelVersionUpgradeOptionNoAutoUpgrade,
		"oncecurrentversionexpired":      DeploymentModelVersionUpgradeOptionOnceCurrentVersionExpired,
		"oncenewdefaultversionavailable": DeploymentModelVersionUpgradeOptionOnceNewDefaultVersionAvailable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeploymentModelVersionUpgradeOption(input)
	return &out, nil
}

type EndpointComputeType string

const (
	EndpointComputeTypeAzureMLCompute EndpointComputeType = "AzureMLCompute"
	EndpointComputeTypeKubernetes     EndpointComputeType = "Kubernetes"
	EndpointComputeTypeManaged        EndpointComputeType = "Managed"
)

func PossibleValuesForEndpointComputeType() []string {
	return []string{
		string(EndpointComputeTypeAzureMLCompute),
		string(EndpointComputeTypeKubernetes),
		string(EndpointComputeTypeManaged),
	}
}

func (s *EndpointComputeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEndpointComputeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEndpointComputeType(input string) (*EndpointComputeType, error) {
	vals := map[string]EndpointComputeType{
		"azuremlcompute": EndpointComputeTypeAzureMLCompute,
		"kubernetes":     EndpointComputeTypeKubernetes,
		"managed":        EndpointComputeTypeManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndpointComputeType(input)
	return &out, nil
}

type ManagedPERequirement string

const (
	ManagedPERequirementNotApplicable ManagedPERequirement = "NotApplicable"
	ManagedPERequirementNotRequired   ManagedPERequirement = "NotRequired"
	ManagedPERequirementRequired      ManagedPERequirement = "Required"
)

func PossibleValuesForManagedPERequirement() []string {
	return []string{
		string(ManagedPERequirementNotApplicable),
		string(ManagedPERequirementNotRequired),
		string(ManagedPERequirementRequired),
	}
}

func (s *ManagedPERequirement) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedPERequirement(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedPERequirement(input string) (*ManagedPERequirement, error) {
	vals := map[string]ManagedPERequirement{
		"notapplicable": ManagedPERequirementNotApplicable,
		"notrequired":   ManagedPERequirementNotRequired,
		"required":      ManagedPERequirementRequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedPERequirement(input)
	return &out, nil
}

type ManagedPEStatus string

const (
	ManagedPEStatusActive        ManagedPEStatus = "Active"
	ManagedPEStatusInactive      ManagedPEStatus = "Inactive"
	ManagedPEStatusNotApplicable ManagedPEStatus = "NotApplicable"
)

func PossibleValuesForManagedPEStatus() []string {
	return []string{
		string(ManagedPEStatusActive),
		string(ManagedPEStatusInactive),
		string(ManagedPEStatusNotApplicable),
	}
}

func (s *ManagedPEStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedPEStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedPEStatus(input string) (*ManagedPEStatus, error) {
	vals := map[string]ManagedPEStatus{
		"active":        ManagedPEStatusActive,
		"inactive":      ManagedPEStatusInactive,
		"notapplicable": ManagedPEStatusNotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedPEStatus(input)
	return &out, nil
}

type ModelLifecycleStatus string

const (
	ModelLifecycleStatusGenerallyAvailable ModelLifecycleStatus = "GenerallyAvailable"
	ModelLifecycleStatusPreview            ModelLifecycleStatus = "Preview"
)

func PossibleValuesForModelLifecycleStatus() []string {
	return []string{
		string(ModelLifecycleStatusGenerallyAvailable),
		string(ModelLifecycleStatusPreview),
	}
}

func (s *ModelLifecycleStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseModelLifecycleStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseModelLifecycleStatus(input string) (*ModelLifecycleStatus, error) {
	vals := map[string]ModelLifecycleStatus{
		"generallyavailable": ModelLifecycleStatusGenerallyAvailable,
		"preview":            ModelLifecycleStatusPreview,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ModelLifecycleStatus(input)
	return &out, nil
}

type RaiPolicyContentSource string

const (
	RaiPolicyContentSourceCompletion RaiPolicyContentSource = "Completion"
	RaiPolicyContentSourcePrompt     RaiPolicyContentSource = "Prompt"
)

func PossibleValuesForRaiPolicyContentSource() []string {
	return []string{
		string(RaiPolicyContentSourceCompletion),
		string(RaiPolicyContentSourcePrompt),
	}
}

func (s *RaiPolicyContentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRaiPolicyContentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRaiPolicyContentSource(input string) (*RaiPolicyContentSource, error) {
	vals := map[string]RaiPolicyContentSource{
		"completion": RaiPolicyContentSourceCompletion,
		"prompt":     RaiPolicyContentSourcePrompt,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RaiPolicyContentSource(input)
	return &out, nil
}

type RaiPolicyMode string

const (
	RaiPolicyModeBlocking RaiPolicyMode = "Blocking"
	RaiPolicyModeDefault  RaiPolicyMode = "Default"
	RaiPolicyModeDeferred RaiPolicyMode = "Deferred"
)

func PossibleValuesForRaiPolicyMode() []string {
	return []string{
		string(RaiPolicyModeBlocking),
		string(RaiPolicyModeDefault),
		string(RaiPolicyModeDeferred),
	}
}

func (s *RaiPolicyMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRaiPolicyMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRaiPolicyMode(input string) (*RaiPolicyMode, error) {
	vals := map[string]RaiPolicyMode{
		"blocking": RaiPolicyModeBlocking,
		"default":  RaiPolicyModeDefault,
		"deferred": RaiPolicyModeDeferred,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RaiPolicyMode(input)
	return &out, nil
}

type RaiPolicyType string

const (
	RaiPolicyTypeSystemManaged RaiPolicyType = "SystemManaged"
	RaiPolicyTypeUserManaged   RaiPolicyType = "UserManaged"
)

func PossibleValuesForRaiPolicyType() []string {
	return []string{
		string(RaiPolicyTypeSystemManaged),
		string(RaiPolicyTypeUserManaged),
	}
}

func (s *RaiPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRaiPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRaiPolicyType(input string) (*RaiPolicyType, error) {
	vals := map[string]RaiPolicyType{
		"systemmanaged": RaiPolicyTypeSystemManaged,
		"usermanaged":   RaiPolicyTypeUserManaged,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RaiPolicyType(input)
	return &out, nil
}
