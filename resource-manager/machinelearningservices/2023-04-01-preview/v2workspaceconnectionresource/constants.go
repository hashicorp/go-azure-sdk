package v2workspaceconnectionresource

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionAuthType string

const (
	ConnectionAuthTypeAccessKey        ConnectionAuthType = "AccessKey"
	ConnectionAuthTypeManagedIdentity  ConnectionAuthType = "ManagedIdentity"
	ConnectionAuthTypeNone             ConnectionAuthType = "None"
	ConnectionAuthTypePAT              ConnectionAuthType = "PAT"
	ConnectionAuthTypeSAS              ConnectionAuthType = "SAS"
	ConnectionAuthTypeServicePrincipal ConnectionAuthType = "ServicePrincipal"
	ConnectionAuthTypeUsernamePassword ConnectionAuthType = "UsernamePassword"
)

func PossibleValuesForConnectionAuthType() []string {
	return []string{
		string(ConnectionAuthTypeAccessKey),
		string(ConnectionAuthTypeManagedIdentity),
		string(ConnectionAuthTypeNone),
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
		"accesskey":        ConnectionAuthTypeAccessKey,
		"managedidentity":  ConnectionAuthTypeManagedIdentity,
		"none":             ConnectionAuthTypeNone,
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
	ConnectionCategoryAzureDataLakeGenTwo   ConnectionCategory = "AzureDataLakeGen2"
	ConnectionCategoryAzureMySqlDb          ConnectionCategory = "AzureMySqlDb"
	ConnectionCategoryAzurePostgresDb       ConnectionCategory = "AzurePostgresDb"
	ConnectionCategoryAzureSqlDb            ConnectionCategory = "AzureSqlDb"
	ConnectionCategoryAzureSynapseAnalytics ConnectionCategory = "AzureSynapseAnalytics"
	ConnectionCategoryContainerRegistry     ConnectionCategory = "ContainerRegistry"
	ConnectionCategoryFeatureStore          ConnectionCategory = "FeatureStore"
	ConnectionCategoryGit                   ConnectionCategory = "Git"
	ConnectionCategoryPythonFeed            ConnectionCategory = "PythonFeed"
	ConnectionCategoryRedis                 ConnectionCategory = "Redis"
	ConnectionCategorySThree                ConnectionCategory = "S3"
	ConnectionCategorySnowflake             ConnectionCategory = "Snowflake"
)

func PossibleValuesForConnectionCategory() []string {
	return []string{
		string(ConnectionCategoryAzureDataLakeGenTwo),
		string(ConnectionCategoryAzureMySqlDb),
		string(ConnectionCategoryAzurePostgresDb),
		string(ConnectionCategoryAzureSqlDb),
		string(ConnectionCategoryAzureSynapseAnalytics),
		string(ConnectionCategoryContainerRegistry),
		string(ConnectionCategoryFeatureStore),
		string(ConnectionCategoryGit),
		string(ConnectionCategoryPythonFeed),
		string(ConnectionCategoryRedis),
		string(ConnectionCategorySThree),
		string(ConnectionCategorySnowflake),
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
		"azuredatalakegen2":     ConnectionCategoryAzureDataLakeGenTwo,
		"azuremysqldb":          ConnectionCategoryAzureMySqlDb,
		"azurepostgresdb":       ConnectionCategoryAzurePostgresDb,
		"azuresqldb":            ConnectionCategoryAzureSqlDb,
		"azuresynapseanalytics": ConnectionCategoryAzureSynapseAnalytics,
		"containerregistry":     ConnectionCategoryContainerRegistry,
		"featurestore":          ConnectionCategoryFeatureStore,
		"git":                   ConnectionCategoryGit,
		"pythonfeed":            ConnectionCategoryPythonFeed,
		"redis":                 ConnectionCategoryRedis,
		"s3":                    ConnectionCategorySThree,
		"snowflake":             ConnectionCategorySnowflake,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionCategory(input)
	return &out, nil
}

type ValueFormat string

const (
	ValueFormatJSON ValueFormat = "JSON"
)

func PossibleValuesForValueFormat() []string {
	return []string{
		string(ValueFormatJSON),
	}
}

func (s *ValueFormat) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseValueFormat(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseValueFormat(input string) (*ValueFormat, error) {
	vals := map[string]ValueFormat{
		"json": ValueFormatJSON,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValueFormat(input)
	return &out, nil
}
