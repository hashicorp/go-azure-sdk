package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllDomainsRootDomains string

const (
	AllDomainsRootDomainsall                              AllDomainsRootDomains = "All"
	AllDomainsRootDomainsallFederated                     AllDomainsRootDomains = "AllFederated"
	AllDomainsRootDomainsallManaged                       AllDomainsRootDomains = "AllManaged"
	AllDomainsRootDomainsallManagedAndEnumeratedFederated AllDomainsRootDomains = "AllManagedAndEnumeratedFederated"
	AllDomainsRootDomainsenumerated                       AllDomainsRootDomains = "Enumerated"
	AllDomainsRootDomainsnone                             AllDomainsRootDomains = "None"
)

func PossibleValuesForAllDomainsRootDomains() []string {
	return []string{
		string(AllDomainsRootDomainsall),
		string(AllDomainsRootDomainsallFederated),
		string(AllDomainsRootDomainsallManaged),
		string(AllDomainsRootDomainsallManagedAndEnumeratedFederated),
		string(AllDomainsRootDomainsenumerated),
		string(AllDomainsRootDomainsnone),
	}
}

func parseAllDomainsRootDomains(input string) (*AllDomainsRootDomains, error) {
	vals := map[string]AllDomainsRootDomains{
		"all":                              AllDomainsRootDomainsall,
		"allfederated":                     AllDomainsRootDomainsallFederated,
		"allmanaged":                       AllDomainsRootDomainsallManaged,
		"allmanagedandenumeratedfederated": AllDomainsRootDomainsallManagedAndEnumeratedFederated,
		"enumerated":                       AllDomainsRootDomainsenumerated,
		"none":                             AllDomainsRootDomainsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AllDomainsRootDomains(input)
	return &out, nil
}
