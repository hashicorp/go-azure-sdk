package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidatingDomainsRootDomains string

const (
	ValidatingDomainsRootDomains_All                              ValidatingDomainsRootDomains = "all"
	ValidatingDomainsRootDomains_AllFederated                     ValidatingDomainsRootDomains = "allFederated"
	ValidatingDomainsRootDomains_AllManaged                       ValidatingDomainsRootDomains = "allManaged"
	ValidatingDomainsRootDomains_AllManagedAndEnumeratedFederated ValidatingDomainsRootDomains = "allManagedAndEnumeratedFederated"
	ValidatingDomainsRootDomains_Enumerated                       ValidatingDomainsRootDomains = "enumerated"
	ValidatingDomainsRootDomains_None                             ValidatingDomainsRootDomains = "none"
)

func PossibleValuesForValidatingDomainsRootDomains() []string {
	return []string{
		string(ValidatingDomainsRootDomains_All),
		string(ValidatingDomainsRootDomains_AllFederated),
		string(ValidatingDomainsRootDomains_AllManaged),
		string(ValidatingDomainsRootDomains_AllManagedAndEnumeratedFederated),
		string(ValidatingDomainsRootDomains_Enumerated),
		string(ValidatingDomainsRootDomains_None),
	}
}

func (s *ValidatingDomainsRootDomains) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseValidatingDomainsRootDomains(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseValidatingDomainsRootDomains(input string) (*ValidatingDomainsRootDomains, error) {
	vals := map[string]ValidatingDomainsRootDomains{
		"all":                              ValidatingDomainsRootDomains_All,
		"allfederated":                     ValidatingDomainsRootDomains_AllFederated,
		"allmanaged":                       ValidatingDomainsRootDomains_AllManaged,
		"allmanagedandenumeratedfederated": ValidatingDomainsRootDomains_AllManagedAndEnumeratedFederated,
		"enumerated":                       ValidatingDomainsRootDomains_Enumerated,
		"none":                             ValidatingDomainsRootDomains_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValidatingDomainsRootDomains(input)
	return &out, nil
}
