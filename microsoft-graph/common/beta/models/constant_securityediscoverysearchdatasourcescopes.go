package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoverySearchDataSourceScopes string

const (
	SecurityEdiscoverySearchDataSourceScopesallCaseCustodians              SecurityEdiscoverySearchDataSourceScopes = "AllCaseCustodians"
	SecurityEdiscoverySearchDataSourceScopesallCaseNoncustodialDataSources SecurityEdiscoverySearchDataSourceScopes = "AllCaseNoncustodialDataSources"
	SecurityEdiscoverySearchDataSourceScopesallTenantMailboxes             SecurityEdiscoverySearchDataSourceScopes = "AllTenantMailboxes"
	SecurityEdiscoverySearchDataSourceScopesallTenantSites                 SecurityEdiscoverySearchDataSourceScopes = "AllTenantSites"
	SecurityEdiscoverySearchDataSourceScopesnone                           SecurityEdiscoverySearchDataSourceScopes = "None"
)

func PossibleValuesForSecurityEdiscoverySearchDataSourceScopes() []string {
	return []string{
		string(SecurityEdiscoverySearchDataSourceScopesallCaseCustodians),
		string(SecurityEdiscoverySearchDataSourceScopesallCaseNoncustodialDataSources),
		string(SecurityEdiscoverySearchDataSourceScopesallTenantMailboxes),
		string(SecurityEdiscoverySearchDataSourceScopesallTenantSites),
		string(SecurityEdiscoverySearchDataSourceScopesnone),
	}
}

func parseSecurityEdiscoverySearchDataSourceScopes(input string) (*SecurityEdiscoverySearchDataSourceScopes, error) {
	vals := map[string]SecurityEdiscoverySearchDataSourceScopes{
		"allcasecustodians":              SecurityEdiscoverySearchDataSourceScopesallCaseCustodians,
		"allcasenoncustodialdatasources": SecurityEdiscoverySearchDataSourceScopesallCaseNoncustodialDataSources,
		"alltenantmailboxes":             SecurityEdiscoverySearchDataSourceScopesallTenantMailboxes,
		"alltenantsites":                 SecurityEdiscoverySearchDataSourceScopesallTenantSites,
		"none":                           SecurityEdiscoverySearchDataSourceScopesnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoverySearchDataSourceScopes(input)
	return &out, nil
}
