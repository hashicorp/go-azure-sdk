package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultManagedAppProtectionAllowedDataStorageLocations string

const (
	DefaultManagedAppProtectionAllowedDataStorageLocations_Box                 DefaultManagedAppProtectionAllowedDataStorageLocations = "box"
	DefaultManagedAppProtectionAllowedDataStorageLocations_LocalStorage        DefaultManagedAppProtectionAllowedDataStorageLocations = "localStorage"
	DefaultManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness DefaultManagedAppProtectionAllowedDataStorageLocations = "oneDriveForBusiness"
	DefaultManagedAppProtectionAllowedDataStorageLocations_PhotoLibrary        DefaultManagedAppProtectionAllowedDataStorageLocations = "photoLibrary"
	DefaultManagedAppProtectionAllowedDataStorageLocations_SharePoint          DefaultManagedAppProtectionAllowedDataStorageLocations = "sharePoint"
)

func PossibleValuesForDefaultManagedAppProtectionAllowedDataStorageLocations() []string {
	return []string{
		string(DefaultManagedAppProtectionAllowedDataStorageLocations_Box),
		string(DefaultManagedAppProtectionAllowedDataStorageLocations_LocalStorage),
		string(DefaultManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness),
		string(DefaultManagedAppProtectionAllowedDataStorageLocations_PhotoLibrary),
		string(DefaultManagedAppProtectionAllowedDataStorageLocations_SharePoint),
	}
}

func (s *DefaultManagedAppProtectionAllowedDataStorageLocations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefaultManagedAppProtectionAllowedDataStorageLocations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefaultManagedAppProtectionAllowedDataStorageLocations(input string) (*DefaultManagedAppProtectionAllowedDataStorageLocations, error) {
	vals := map[string]DefaultManagedAppProtectionAllowedDataStorageLocations{
		"box":                 DefaultManagedAppProtectionAllowedDataStorageLocations_Box,
		"localstorage":        DefaultManagedAppProtectionAllowedDataStorageLocations_LocalStorage,
		"onedriveforbusiness": DefaultManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness,
		"photolibrary":        DefaultManagedAppProtectionAllowedDataStorageLocations_PhotoLibrary,
		"sharepoint":          DefaultManagedAppProtectionAllowedDataStorageLocations_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefaultManagedAppProtectionAllowedDataStorageLocations(input)
	return &out, nil
}
