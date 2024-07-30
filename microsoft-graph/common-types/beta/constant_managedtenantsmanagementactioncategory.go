package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagementActionCategory string

const (
	ManagedTenantsManagementActionCategory_Custom   ManagedTenantsManagementActionCategory = "custom"
	ManagedTenantsManagementActionCategory_Data     ManagedTenantsManagementActionCategory = "data"
	ManagedTenantsManagementActionCategory_Devices  ManagedTenantsManagementActionCategory = "devices"
	ManagedTenantsManagementActionCategory_Identity ManagedTenantsManagementActionCategory = "identity"
)

func PossibleValuesForManagedTenantsManagementActionCategory() []string {
	return []string{
		string(ManagedTenantsManagementActionCategory_Custom),
		string(ManagedTenantsManagementActionCategory_Data),
		string(ManagedTenantsManagementActionCategory_Devices),
		string(ManagedTenantsManagementActionCategory_Identity),
	}
}

func (s *ManagedTenantsManagementActionCategory) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsManagementActionCategory(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsManagementActionCategory(input string) (*ManagedTenantsManagementActionCategory, error) {
	vals := map[string]ManagedTenantsManagementActionCategory{
		"custom":   ManagedTenantsManagementActionCategory_Custom,
		"data":     ManagedTenantsManagementActionCategory_Data,
		"devices":  ManagedTenantsManagementActionCategory_Devices,
		"identity": ManagedTenantsManagementActionCategory_Identity,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsManagementActionCategory(input)
	return &out, nil
}
