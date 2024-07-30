package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAllowedDataStorageLocations string

const (
	IosManagedAppProtectionAllowedDataStorageLocations_Box                 IosManagedAppProtectionAllowedDataStorageLocations = "box"
	IosManagedAppProtectionAllowedDataStorageLocations_LocalStorage        IosManagedAppProtectionAllowedDataStorageLocations = "localStorage"
	IosManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness IosManagedAppProtectionAllowedDataStorageLocations = "oneDriveForBusiness"
	IosManagedAppProtectionAllowedDataStorageLocations_SharePoint          IosManagedAppProtectionAllowedDataStorageLocations = "sharePoint"
)

func PossibleValuesForIosManagedAppProtectionAllowedDataStorageLocations() []string {
	return []string{
		string(IosManagedAppProtectionAllowedDataStorageLocations_Box),
		string(IosManagedAppProtectionAllowedDataStorageLocations_LocalStorage),
		string(IosManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness),
		string(IosManagedAppProtectionAllowedDataStorageLocations_SharePoint),
	}
}

func (s *IosManagedAppProtectionAllowedDataStorageLocations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAllowedDataStorageLocations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAllowedDataStorageLocations(input string) (*IosManagedAppProtectionAllowedDataStorageLocations, error) {
	vals := map[string]IosManagedAppProtectionAllowedDataStorageLocations{
		"box":                 IosManagedAppProtectionAllowedDataStorageLocations_Box,
		"localstorage":        IosManagedAppProtectionAllowedDataStorageLocations_LocalStorage,
		"onedriveforbusiness": IosManagedAppProtectionAllowedDataStorageLocations_OneDriveForBusiness,
		"sharepoint":          IosManagedAppProtectionAllowedDataStorageLocations_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAllowedDataStorageLocations(input)
	return &out, nil
}
