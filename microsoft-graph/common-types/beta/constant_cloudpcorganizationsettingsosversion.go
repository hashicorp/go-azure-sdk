package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOrganizationSettingsOsVersion string

const (
	CloudPCOrganizationSettingsOsVersion_Windows10 CloudPCOrganizationSettingsOsVersion = "windows10"
	CloudPCOrganizationSettingsOsVersion_Windows11 CloudPCOrganizationSettingsOsVersion = "windows11"
)

func PossibleValuesForCloudPCOrganizationSettingsOsVersion() []string {
	return []string{
		string(CloudPCOrganizationSettingsOsVersion_Windows10),
		string(CloudPCOrganizationSettingsOsVersion_Windows11),
	}
}

func (s *CloudPCOrganizationSettingsOsVersion) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCOrganizationSettingsOsVersion(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCOrganizationSettingsOsVersion(input string) (*CloudPCOrganizationSettingsOsVersion, error) {
	vals := map[string]CloudPCOrganizationSettingsOsVersion{
		"windows10": CloudPCOrganizationSettingsOsVersion_Windows10,
		"windows11": CloudPCOrganizationSettingsOsVersion_Windows11,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOrganizationSettingsOsVersion(input)
	return &out, nil
}
