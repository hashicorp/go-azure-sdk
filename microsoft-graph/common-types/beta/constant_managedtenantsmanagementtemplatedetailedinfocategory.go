package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateDetailedInfoCategory string

const (
	ManagedTenantsManagementTemplateDetailedInfoCategory_Custom   ManagedTenantsManagementTemplateDetailedInfoCategory = "custom"
	ManagedTenantsManagementTemplateDetailedInfoCategory_Data     ManagedTenantsManagementTemplateDetailedInfoCategory = "data"
	ManagedTenantsManagementTemplateDetailedInfoCategory_Devices  ManagedTenantsManagementTemplateDetailedInfoCategory = "devices"
	ManagedTenantsManagementTemplateDetailedInfoCategory_Identity ManagedTenantsManagementTemplateDetailedInfoCategory = "identity"
)

func PossibleValuesForManagedTenantsManagementTemplateDetailedInfoCategory() []string {
	return []string{
		string(ManagedTenantsManagementTemplateDetailedInfoCategory_Custom),
		string(ManagedTenantsManagementTemplateDetailedInfoCategory_Data),
		string(ManagedTenantsManagementTemplateDetailedInfoCategory_Devices),
		string(ManagedTenantsManagementTemplateDetailedInfoCategory_Identity),
	}
}

func (s *ManagedTenantsManagementTemplateDetailedInfoCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementTemplateDetailedInfoCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementTemplateDetailedInfoCategory(input string) (*ManagedTenantsManagementTemplateDetailedInfoCategory, error) {
	vals := map[string]ManagedTenantsManagementTemplateDetailedInfoCategory{
		"custom":   ManagedTenantsManagementTemplateDetailedInfoCategory_Custom,
		"data":     ManagedTenantsManagementTemplateDetailedInfoCategory_Data,
		"devices":  ManagedTenantsManagementTemplateDetailedInfoCategory_Devices,
		"identity": ManagedTenantsManagementTemplateDetailedInfoCategory_Identity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateDetailedInfoCategory(input)
	return &out, nil
}
