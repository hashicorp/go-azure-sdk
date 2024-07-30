package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoverySourceCollectionDataSourceScopes string

const (
	EdiscoverySourceCollectionDataSourceScopes_AllCaseCustodians              EdiscoverySourceCollectionDataSourceScopes = "allCaseCustodians"
	EdiscoverySourceCollectionDataSourceScopes_AllCaseNoncustodialDataSources EdiscoverySourceCollectionDataSourceScopes = "allCaseNoncustodialDataSources"
	EdiscoverySourceCollectionDataSourceScopes_AllTenantMailboxes             EdiscoverySourceCollectionDataSourceScopes = "allTenantMailboxes"
	EdiscoverySourceCollectionDataSourceScopes_AllTenantSites                 EdiscoverySourceCollectionDataSourceScopes = "allTenantSites"
	EdiscoverySourceCollectionDataSourceScopes_None                           EdiscoverySourceCollectionDataSourceScopes = "none"
)

func PossibleValuesForEdiscoverySourceCollectionDataSourceScopes() []string {
	return []string{
		string(EdiscoverySourceCollectionDataSourceScopes_AllCaseCustodians),
		string(EdiscoverySourceCollectionDataSourceScopes_AllCaseNoncustodialDataSources),
		string(EdiscoverySourceCollectionDataSourceScopes_AllTenantMailboxes),
		string(EdiscoverySourceCollectionDataSourceScopes_AllTenantSites),
		string(EdiscoverySourceCollectionDataSourceScopes_None),
	}
}

func (s *EdiscoverySourceCollectionDataSourceScopes) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoverySourceCollectionDataSourceScopes(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoverySourceCollectionDataSourceScopes(input string) (*EdiscoverySourceCollectionDataSourceScopes, error) {
	vals := map[string]EdiscoverySourceCollectionDataSourceScopes{
		"allcasecustodians":              EdiscoverySourceCollectionDataSourceScopes_AllCaseCustodians,
		"allcasenoncustodialdatasources": EdiscoverySourceCollectionDataSourceScopes_AllCaseNoncustodialDataSources,
		"alltenantmailboxes":             EdiscoverySourceCollectionDataSourceScopes_AllTenantMailboxes,
		"alltenantsites":                 EdiscoverySourceCollectionDataSourceScopes_AllTenantSites,
		"none":                           EdiscoverySourceCollectionDataSourceScopes_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoverySourceCollectionDataSourceScopes(input)
	return &out, nil
}
