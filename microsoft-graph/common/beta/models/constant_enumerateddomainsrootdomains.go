package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnumeratedDomainsRootDomains string

const (
	EnumeratedDomainsRootDomainsall                              EnumeratedDomainsRootDomains = "All"
	EnumeratedDomainsRootDomainsallFederated                     EnumeratedDomainsRootDomains = "AllFederated"
	EnumeratedDomainsRootDomainsallManaged                       EnumeratedDomainsRootDomains = "AllManaged"
	EnumeratedDomainsRootDomainsallManagedAndEnumeratedFederated EnumeratedDomainsRootDomains = "AllManagedAndEnumeratedFederated"
	EnumeratedDomainsRootDomainsenumerated                       EnumeratedDomainsRootDomains = "Enumerated"
	EnumeratedDomainsRootDomainsnone                             EnumeratedDomainsRootDomains = "None"
)

func PossibleValuesForEnumeratedDomainsRootDomains() []string {
	return []string{
		string(EnumeratedDomainsRootDomainsall),
		string(EnumeratedDomainsRootDomainsallFederated),
		string(EnumeratedDomainsRootDomainsallManaged),
		string(EnumeratedDomainsRootDomainsallManagedAndEnumeratedFederated),
		string(EnumeratedDomainsRootDomainsenumerated),
		string(EnumeratedDomainsRootDomainsnone),
	}
}

func parseEnumeratedDomainsRootDomains(input string) (*EnumeratedDomainsRootDomains, error) {
	vals := map[string]EnumeratedDomainsRootDomains{
		"all":                              EnumeratedDomainsRootDomainsall,
		"allfederated":                     EnumeratedDomainsRootDomainsallFederated,
		"allmanaged":                       EnumeratedDomainsRootDomainsallManaged,
		"allmanagedandenumeratedfederated": EnumeratedDomainsRootDomainsallManagedAndEnumeratedFederated,
		"enumerated":                       EnumeratedDomainsRootDomainsenumerated,
		"none":                             EnumeratedDomainsRootDomainsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EnumeratedDomainsRootDomains(input)
	return &out, nil
}
