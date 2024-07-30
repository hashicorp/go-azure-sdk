package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetedManagedAppProtectionAllowedDataStorageLocations string

const (
	TargetedManagedAppProtectionAllowedDataStorageLocations_Box                 TargetedManagedAppProtectionAllowedDataStorageLocations = "box"
	TargetedManagedAppProtectionAllowedDataStorageLocations_LocalStorage        TargetedManagedAppProtectionAllowedDataStorageLocations = "localStorage"
	TargetedManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness TargetedManagedAppProtectionAllowedDataStorageLocations = "oneDriveForBusiness"
	TargetedManagedAppProtectionAllowedDataStorageLocations_PhotoLibrary        TargetedManagedAppProtectionAllowedDataStorageLocations = "photoLibrary"
	TargetedManagedAppProtectionAllowedDataStorageLocations_SharePoint          TargetedManagedAppProtectionAllowedDataStorageLocations = "sharePoint"
)

func PossibleValuesForTargetedManagedAppProtectionAllowedDataStorageLocations() []string {
	return []string{
		string(TargetedManagedAppProtectionAllowedDataStorageLocations_Box),
		string(TargetedManagedAppProtectionAllowedDataStorageLocations_LocalStorage),
		string(TargetedManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness),
		string(TargetedManagedAppProtectionAllowedDataStorageLocations_PhotoLibrary),
		string(TargetedManagedAppProtectionAllowedDataStorageLocations_SharePoint),
	}
}

func (s *TargetedManagedAppProtectionAllowedDataStorageLocations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTargetedManagedAppProtectionAllowedDataStorageLocations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTargetedManagedAppProtectionAllowedDataStorageLocations(input string) (*TargetedManagedAppProtectionAllowedDataStorageLocations, error) {
	vals := map[string]TargetedManagedAppProtectionAllowedDataStorageLocations{
		"box":                 TargetedManagedAppProtectionAllowedDataStorageLocations_Box,
		"localstorage":        TargetedManagedAppProtectionAllowedDataStorageLocations_LocalStorage,
		"onedriveforbusiness": TargetedManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness,
		"photolibrary":        TargetedManagedAppProtectionAllowedDataStorageLocations_PhotoLibrary,
		"sharepoint":          TargetedManagedAppProtectionAllowedDataStorageLocations_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TargetedManagedAppProtectionAllowedDataStorageLocations(input)
	return &out, nil
}
