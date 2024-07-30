package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementTemplateStepCategory string

const (
	ManagedTenantsManagementTemplateStepCategory_Custom   ManagedTenantsManagementTemplateStepCategory = "custom"
	ManagedTenantsManagementTemplateStepCategory_Data     ManagedTenantsManagementTemplateStepCategory = "data"
	ManagedTenantsManagementTemplateStepCategory_Devices  ManagedTenantsManagementTemplateStepCategory = "devices"
	ManagedTenantsManagementTemplateStepCategory_Identity ManagedTenantsManagementTemplateStepCategory = "identity"
)

func PossibleValuesForManagedTenantsManagementTemplateStepCategory() []string {
	return []string{
		string(ManagedTenantsManagementTemplateStepCategory_Custom),
		string(ManagedTenantsManagementTemplateStepCategory_Data),
		string(ManagedTenantsManagementTemplateStepCategory_Devices),
		string(ManagedTenantsManagementTemplateStepCategory_Identity),
	}
}

func (s *ManagedTenantsManagementTemplateStepCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementTemplateStepCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementTemplateStepCategory(input string) (*ManagedTenantsManagementTemplateStepCategory, error) {
	vals := map[string]ManagedTenantsManagementTemplateStepCategory{
		"custom":   ManagedTenantsManagementTemplateStepCategory_Custom,
		"data":     ManagedTenantsManagementTemplateStepCategory_Data,
		"devices":  ManagedTenantsManagementTemplateStepCategory_Devices,
		"identity": ManagedTenantsManagementTemplateStepCategory_Identity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementTemplateStepCategory(input)
	return &out, nil
}
