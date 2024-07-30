package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateCategory string

const (
	ManagedTenantsManagementTemplateCategory_Custom   ManagedTenantsManagementTemplateCategory = "custom"
	ManagedTenantsManagementTemplateCategory_Data     ManagedTenantsManagementTemplateCategory = "data"
	ManagedTenantsManagementTemplateCategory_Devices  ManagedTenantsManagementTemplateCategory = "devices"
	ManagedTenantsManagementTemplateCategory_Identity ManagedTenantsManagementTemplateCategory = "identity"
)

func PossibleValuesForManagedTenantsManagementTemplateCategory() []string {
	return []string{
		string(ManagedTenantsManagementTemplateCategory_Custom),
		string(ManagedTenantsManagementTemplateCategory_Data),
		string(ManagedTenantsManagementTemplateCategory_Devices),
		string(ManagedTenantsManagementTemplateCategory_Identity),
	}
}

func (s *ManagedTenantsManagementTemplateCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementTemplateCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementTemplateCategory(input string) (*ManagedTenantsManagementTemplateCategory, error) {
	vals := map[string]ManagedTenantsManagementTemplateCategory{
		"custom":   ManagedTenantsManagementTemplateCategory_Custom,
		"data":     ManagedTenantsManagementTemplateCategory_Data,
		"devices":  ManagedTenantsManagementTemplateCategory_Devices,
		"identity": ManagedTenantsManagementTemplateCategory_Identity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateCategory(input)
	return &out, nil
}
