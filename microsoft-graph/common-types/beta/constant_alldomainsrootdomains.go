package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllDomainsRootDomains string

const (
	AllDomainsRootDomains_All                              AllDomainsRootDomains = "all"
	AllDomainsRootDomains_AllFederated                     AllDomainsRootDomains = "allFederated"
	AllDomainsRootDomains_AllManaged                       AllDomainsRootDomains = "allManaged"
	AllDomainsRootDomains_AllManagedAndEnumeratedFederated AllDomainsRootDomains = "allManagedAndEnumeratedFederated"
	AllDomainsRootDomains_Enumerated                       AllDomainsRootDomains = "enumerated"
	AllDomainsRootDomains_None                             AllDomainsRootDomains = "none"
)

func PossibleValuesForAllDomainsRootDomains() []string {
	return []string{
		string(AllDomainsRootDomains_All),
		string(AllDomainsRootDomains_AllFederated),
		string(AllDomainsRootDomains_AllManaged),
		string(AllDomainsRootDomains_AllManagedAndEnumeratedFederated),
		string(AllDomainsRootDomains_Enumerated),
		string(AllDomainsRootDomains_None),
	}
}

func (s *AllDomainsRootDomains) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAllDomainsRootDomains(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAllDomainsRootDomains(input string) (*AllDomainsRootDomains, error) {
	vals := map[string]AllDomainsRootDomains{
		"all":                              AllDomainsRootDomains_All,
		"allfederated":                     AllDomainsRootDomains_AllFederated,
		"allmanaged":                       AllDomainsRootDomains_AllManaged,
		"allmanagedandenumeratedfederated": AllDomainsRootDomains_AllManagedAndEnumeratedFederated,
		"enumerated":                       AllDomainsRootDomains_Enumerated,
		"none":                             AllDomainsRootDomains_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllDomainsRootDomains(input)
	return &out, nil
}
