package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ValidatingDomainsRootDomains string

const (
	ValidatingDomainsRootDomainsall                              ValidatingDomainsRootDomains = "All"
	ValidatingDomainsRootDomainsallFederated                     ValidatingDomainsRootDomains = "AllFederated"
	ValidatingDomainsRootDomainsallManaged                       ValidatingDomainsRootDomains = "AllManaged"
	ValidatingDomainsRootDomainsallManagedAndEnumeratedFederated ValidatingDomainsRootDomains = "AllManagedAndEnumeratedFederated"
	ValidatingDomainsRootDomainsenumerated                       ValidatingDomainsRootDomains = "Enumerated"
	ValidatingDomainsRootDomainsnone                             ValidatingDomainsRootDomains = "None"
)

func PossibleValuesForValidatingDomainsRootDomains() []string {
	return []string{
		string(ValidatingDomainsRootDomainsall),
		string(ValidatingDomainsRootDomainsallFederated),
		string(ValidatingDomainsRootDomainsallManaged),
		string(ValidatingDomainsRootDomainsallManagedAndEnumeratedFederated),
		string(ValidatingDomainsRootDomainsenumerated),
		string(ValidatingDomainsRootDomainsnone),
	}
}

func parseValidatingDomainsRootDomains(input string) (*ValidatingDomainsRootDomains, error) {
	vals := map[string]ValidatingDomainsRootDomains{
		"all":                              ValidatingDomainsRootDomainsall,
		"allfederated":                     ValidatingDomainsRootDomainsallFederated,
		"allmanaged":                       ValidatingDomainsRootDomainsallManaged,
		"allmanagedandenumeratedfederated": ValidatingDomainsRootDomainsallManagedAndEnumeratedFederated,
		"enumerated":                       ValidatingDomainsRootDomainsenumerated,
		"none":                             ValidatingDomainsRootDomainsnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ValidatingDomainsRootDomains(input)
	return &out, nil
}
