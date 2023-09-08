package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateProvider string

const (
	ManagedTenantsManagementTemplateProvidercommunity        ManagedTenantsManagementTemplateProvider = "Community"
	ManagedTenantsManagementTemplateProviderindirectProvider ManagedTenantsManagementTemplateProvider = "IndirectProvider"
	ManagedTenantsManagementTemplateProvidermicrosoft        ManagedTenantsManagementTemplateProvider = "Microsoft"
	ManagedTenantsManagementTemplateProviderself             ManagedTenantsManagementTemplateProvider = "Self"
)

func PossibleValuesForManagedTenantsManagementTemplateProvider() []string {
	return []string{
		string(ManagedTenantsManagementTemplateProvidercommunity),
		string(ManagedTenantsManagementTemplateProviderindirectProvider),
		string(ManagedTenantsManagementTemplateProvidermicrosoft),
		string(ManagedTenantsManagementTemplateProviderself),
	}
}

func parseManagedTenantsManagementTemplateProvider(input string) (*ManagedTenantsManagementTemplateProvider, error) {
	vals := map[string]ManagedTenantsManagementTemplateProvider{
		"community":        ManagedTenantsManagementTemplateProvidercommunity,
		"indirectprovider": ManagedTenantsManagementTemplateProviderindirectProvider,
		"microsoft":        ManagedTenantsManagementTemplateProvidermicrosoft,
		"self":             ManagedTenantsManagementTemplateProviderself,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateProvider(input)
	return &out, nil
}
