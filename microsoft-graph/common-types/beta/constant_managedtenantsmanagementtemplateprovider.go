package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateProvider string

const (
	ManagedTenantsManagementTemplateProvider_Community        ManagedTenantsManagementTemplateProvider = "community"
	ManagedTenantsManagementTemplateProvider_IndirectProvider ManagedTenantsManagementTemplateProvider = "indirectProvider"
	ManagedTenantsManagementTemplateProvider_Microsoft        ManagedTenantsManagementTemplateProvider = "microsoft"
	ManagedTenantsManagementTemplateProvider_Self             ManagedTenantsManagementTemplateProvider = "self"
)

func PossibleValuesForManagedTenantsManagementTemplateProvider() []string {
	return []string{
		string(ManagedTenantsManagementTemplateProvider_Community),
		string(ManagedTenantsManagementTemplateProvider_IndirectProvider),
		string(ManagedTenantsManagementTemplateProvider_Microsoft),
		string(ManagedTenantsManagementTemplateProvider_Self),
	}
}

func (s *ManagedTenantsManagementTemplateProvider) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementTemplateProvider(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementTemplateProvider(input string) (*ManagedTenantsManagementTemplateProvider, error) {
	vals := map[string]ManagedTenantsManagementTemplateProvider{
		"community":        ManagedTenantsManagementTemplateProvider_Community,
		"indirectprovider": ManagedTenantsManagementTemplateProvider_IndirectProvider,
		"microsoft":        ManagedTenantsManagementTemplateProvider_Microsoft,
		"self":             ManagedTenantsManagementTemplateProvider_Self,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateProvider(input)
	return &out, nil
}
