package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnumeratedDomainsRootDomains string

const (
	EnumeratedDomainsRootDomains_All                              EnumeratedDomainsRootDomains = "all"
	EnumeratedDomainsRootDomains_AllFederated                     EnumeratedDomainsRootDomains = "allFederated"
	EnumeratedDomainsRootDomains_AllManaged                       EnumeratedDomainsRootDomains = "allManaged"
	EnumeratedDomainsRootDomains_AllManagedAndEnumeratedFederated EnumeratedDomainsRootDomains = "allManagedAndEnumeratedFederated"
	EnumeratedDomainsRootDomains_Enumerated                       EnumeratedDomainsRootDomains = "enumerated"
	EnumeratedDomainsRootDomains_None                             EnumeratedDomainsRootDomains = "none"
)

func PossibleValuesForEnumeratedDomainsRootDomains() []string {
	return []string{
		string(EnumeratedDomainsRootDomains_All),
		string(EnumeratedDomainsRootDomains_AllFederated),
		string(EnumeratedDomainsRootDomains_AllManaged),
		string(EnumeratedDomainsRootDomains_AllManagedAndEnumeratedFederated),
		string(EnumeratedDomainsRootDomains_Enumerated),
		string(EnumeratedDomainsRootDomains_None),
	}
}

func (s *EnumeratedDomainsRootDomains) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEnumeratedDomainsRootDomains(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEnumeratedDomainsRootDomains(input string) (*EnumeratedDomainsRootDomains, error) {
	vals := map[string]EnumeratedDomainsRootDomains{
		"all":                              EnumeratedDomainsRootDomains_All,
		"allfederated":                     EnumeratedDomainsRootDomains_AllFederated,
		"allmanaged":                       EnumeratedDomainsRootDomains_AllManaged,
		"allmanagedandenumeratedfederated": EnumeratedDomainsRootDomains_AllManagedAndEnumeratedFederated,
		"enumerated":                       EnumeratedDomainsRootDomains_Enumerated,
		"none":                             EnumeratedDomainsRootDomains_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnumeratedDomainsRootDomains(input)
	return &out, nil
}
