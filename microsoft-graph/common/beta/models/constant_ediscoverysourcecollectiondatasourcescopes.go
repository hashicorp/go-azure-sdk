package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoverySourceCollectionDataSourceScopes string

const (
	EdiscoverySourceCollectionDataSourceScopesallCaseCustodians              EdiscoverySourceCollectionDataSourceScopes = "AllCaseCustodians"
	EdiscoverySourceCollectionDataSourceScopesallCaseNoncustodialDataSources EdiscoverySourceCollectionDataSourceScopes = "AllCaseNoncustodialDataSources"
	EdiscoverySourceCollectionDataSourceScopesallTenantMailboxes             EdiscoverySourceCollectionDataSourceScopes = "AllTenantMailboxes"
	EdiscoverySourceCollectionDataSourceScopesallTenantSites                 EdiscoverySourceCollectionDataSourceScopes = "AllTenantSites"
	EdiscoverySourceCollectionDataSourceScopesnone                           EdiscoverySourceCollectionDataSourceScopes = "None"
)

func PossibleValuesForEdiscoverySourceCollectionDataSourceScopes() []string {
	return []string{
		string(EdiscoverySourceCollectionDataSourceScopesallCaseCustodians),
		string(EdiscoverySourceCollectionDataSourceScopesallCaseNoncustodialDataSources),
		string(EdiscoverySourceCollectionDataSourceScopesallTenantMailboxes),
		string(EdiscoverySourceCollectionDataSourceScopesallTenantSites),
		string(EdiscoverySourceCollectionDataSourceScopesnone),
	}
}

func parseEdiscoverySourceCollectionDataSourceScopes(input string) (*EdiscoverySourceCollectionDataSourceScopes, error) {
	vals := map[string]EdiscoverySourceCollectionDataSourceScopes{
		"allcasecustodians":              EdiscoverySourceCollectionDataSourceScopesallCaseCustodians,
		"allcasenoncustodialdatasources": EdiscoverySourceCollectionDataSourceScopesallCaseNoncustodialDataSources,
		"alltenantmailboxes":             EdiscoverySourceCollectionDataSourceScopesallTenantMailboxes,
		"alltenantsites":                 EdiscoverySourceCollectionDataSourceScopesallTenantSites,
		"none":                           EdiscoverySourceCollectionDataSourceScopesnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoverySourceCollectionDataSourceScopes(input)
	return &out, nil
}
