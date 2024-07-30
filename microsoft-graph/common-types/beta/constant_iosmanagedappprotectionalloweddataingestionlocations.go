package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppProtectionAllowedDataIngestionLocations string

const (
	IosManagedAppProtectionAllowedDataIngestionLocations_Camera              IosManagedAppProtectionAllowedDataIngestionLocations = "camera"
	IosManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness IosManagedAppProtectionAllowedDataIngestionLocations = "oneDriveForBusiness"
	IosManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary        IosManagedAppProtectionAllowedDataIngestionLocations = "photoLibrary"
	IosManagedAppProtectionAllowedDataIngestionLocations_SharePoint          IosManagedAppProtectionAllowedDataIngestionLocations = "sharePoint"
)

func PossibleValuesForIosManagedAppProtectionAllowedDataIngestionLocations() []string {
	return []string{
		string(IosManagedAppProtectionAllowedDataIngestionLocations_Camera),
		string(IosManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness),
		string(IosManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary),
		string(IosManagedAppProtectionAllowedDataIngestionLocations_SharePoint),
	}
}

func (s *IosManagedAppProtectionAllowedDataIngestionLocations) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosManagedAppProtectionAllowedDataIngestionLocations(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosManagedAppProtectionAllowedDataIngestionLocations(input string) (*IosManagedAppProtectionAllowedDataIngestionLocations, error) {
	vals := map[string]IosManagedAppProtectionAllowedDataIngestionLocations{
		"camera":              IosManagedAppProtectionAllowedDataIngestionLocations_Camera,
		"onedriveforbusiness": IosManagedAppProtectionAllowedDataIngestionLocations_OneDriveForBusiness,
		"photolibrary":        IosManagedAppProtectionAllowedDataIngestionLocations_PhotoLibrary,
		"sharepoint":          IosManagedAppProtectionAllowedDataIngestionLocations_SharePoint,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosManagedAppProtectionAllowedDataIngestionLocations(input)
	return &out, nil
}
