package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedAppProtectionAllowedDataStorageLocations string

const (
	AndroidManagedAppProtectionAllowedDataStorageLocations_Box                 AndroidManagedAppProtectionAllowedDataStorageLocations = "box"
	AndroidManagedAppProtectionAllowedDataStorageLocations_LocalStorage        AndroidManagedAppProtectionAllowedDataStorageLocations = "localStorage"
	AndroidManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness AndroidManagedAppProtectionAllowedDataStorageLocations = "oneDriveForBusiness"
	AndroidManagedAppProtectionAllowedDataStorageLocations_PhotoLibrary        AndroidManagedAppProtectionAllowedDataStorageLocations = "photoLibrary"
	AndroidManagedAppProtectionAllowedDataStorageLocations_SharePoint          AndroidManagedAppProtectionAllowedDataStorageLocations = "sharePoint"
)

func PossibleValuesForAndroidManagedAppProtectionAllowedDataStorageLocations() []string {
	return []string{
		string(AndroidManagedAppProtectionAllowedDataStorageLocations_Box),
		string(AndroidManagedAppProtectionAllowedDataStorageLocations_LocalStorage),
		string(AndroidManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness),
		string(AndroidManagedAppProtectionAllowedDataStorageLocations_PhotoLibrary),
		string(AndroidManagedAppProtectionAllowedDataStorageLocations_SharePoint),
	}
}

func (s *AndroidManagedAppProtectionAllowedDataStorageLocations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidManagedAppProtectionAllowedDataStorageLocations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidManagedAppProtectionAllowedDataStorageLocations(input string) (*AndroidManagedAppProtectionAllowedDataStorageLocations, error) {
	vals := map[string]AndroidManagedAppProtectionAllowedDataStorageLocations{
		"box":                 AndroidManagedAppProtectionAllowedDataStorageLocations_Box,
		"localstorage":        AndroidManagedAppProtectionAllowedDataStorageLocations_LocalStorage,
		"onedriveforbusiness": AndroidManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness,
		"photolibrary":        AndroidManagedAppProtectionAllowedDataStorageLocations_PhotoLibrary,
		"sharepoint":          AndroidManagedAppProtectionAllowedDataStorageLocations_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidManagedAppProtectionAllowedDataStorageLocations(input)
	return &out, nil
}
