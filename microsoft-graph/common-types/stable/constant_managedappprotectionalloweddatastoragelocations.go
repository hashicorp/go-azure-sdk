package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAppProtectionAllowedDataStorageLocations string

const (
	ManagedAppProtectionAllowedDataStorageLocations_Box                 ManagedAppProtectionAllowedDataStorageLocations = "box"
	ManagedAppProtectionAllowedDataStorageLocations_LocalStorage        ManagedAppProtectionAllowedDataStorageLocations = "localStorage"
	ManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness ManagedAppProtectionAllowedDataStorageLocations = "oneDriveForBusiness"
	ManagedAppProtectionAllowedDataStorageLocations_SharePoint          ManagedAppProtectionAllowedDataStorageLocations = "sharePoint"
)

func PossibleValuesForManagedAppProtectionAllowedDataStorageLocations() []string {
	return []string{
		string(ManagedAppProtectionAllowedDataStorageLocations_Box),
		string(ManagedAppProtectionAllowedDataStorageLocations_LocalStorage),
		string(ManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness),
		string(ManagedAppProtectionAllowedDataStorageLocations_SharePoint),
	}
}

func (s *ManagedAppProtectionAllowedDataStorageLocations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAppProtectionAllowedDataStorageLocations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAppProtectionAllowedDataStorageLocations(input string) (*ManagedAppProtectionAllowedDataStorageLocations, error) {
	vals := map[string]ManagedAppProtectionAllowedDataStorageLocations{
		"box":                 ManagedAppProtectionAllowedDataStorageLocations_Box,
		"localstorage":        ManagedAppProtectionAllowedDataStorageLocations_LocalStorage,
		"onedriveforbusiness": ManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness,
		"sharepoint":          ManagedAppProtectionAllowedDataStorageLocations_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAppProtectionAllowedDataStorageLocations(input)
	return &out, nil
}
