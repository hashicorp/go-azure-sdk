package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityEdiscoverySearchDataSourceScopes string

const (
	SecurityEdiscoverySearchDataSourceScopes_AllCaseCustodians              SecurityEdiscoverySearchDataSourceScopes = "allCaseCustodians"
	SecurityEdiscoverySearchDataSourceScopes_AllCaseNoncustodialDataSources SecurityEdiscoverySearchDataSourceScopes = "allCaseNoncustodialDataSources"
	SecurityEdiscoverySearchDataSourceScopes_AllTenantMailboxes             SecurityEdiscoverySearchDataSourceScopes = "allTenantMailboxes"
	SecurityEdiscoverySearchDataSourceScopes_AllTenantSites                 SecurityEdiscoverySearchDataSourceScopes = "allTenantSites"
	SecurityEdiscoverySearchDataSourceScopes_None                           SecurityEdiscoverySearchDataSourceScopes = "none"
)

func PossibleValuesForSecurityEdiscoverySearchDataSourceScopes() []string {
	return []string{
		string(SecurityEdiscoverySearchDataSourceScopes_AllCaseCustodians),
		string(SecurityEdiscoverySearchDataSourceScopes_AllCaseNoncustodialDataSources),
		string(SecurityEdiscoverySearchDataSourceScopes_AllTenantMailboxes),
		string(SecurityEdiscoverySearchDataSourceScopes_AllTenantSites),
		string(SecurityEdiscoverySearchDataSourceScopes_None),
	}
}

func (s *SecurityEdiscoverySearchDataSourceScopes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityEdiscoverySearchDataSourceScopes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityEdiscoverySearchDataSourceScopes(input string) (*SecurityEdiscoverySearchDataSourceScopes, error) {
	vals := map[string]SecurityEdiscoverySearchDataSourceScopes{
		"allcasecustodians":              SecurityEdiscoverySearchDataSourceScopes_AllCaseCustodians,
		"allcasenoncustodialdatasources": SecurityEdiscoverySearchDataSourceScopes_AllCaseNoncustodialDataSources,
		"alltenantmailboxes":             SecurityEdiscoverySearchDataSourceScopes_AllTenantMailboxes,
		"alltenantsites":                 SecurityEdiscoverySearchDataSourceScopes_AllTenantSites,
		"none":                           SecurityEdiscoverySearchDataSourceScopes_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityEdiscoverySearchDataSourceScopes(input)
	return &out, nil
}
