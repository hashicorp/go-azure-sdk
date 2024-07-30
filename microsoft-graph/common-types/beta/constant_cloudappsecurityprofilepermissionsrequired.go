package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudAppSecurityProfilePermissionsRequired string

const (
	CloudAppSecurityProfilePermissionsRequired_Administrator CloudAppSecurityProfilePermissionsRequired = "administrator"
	CloudAppSecurityProfilePermissionsRequired_Anonymous     CloudAppSecurityProfilePermissionsRequired = "anonymous"
	CloudAppSecurityProfilePermissionsRequired_Guest         CloudAppSecurityProfilePermissionsRequired = "guest"
	CloudAppSecurityProfilePermissionsRequired_System        CloudAppSecurityProfilePermissionsRequired = "system"
	CloudAppSecurityProfilePermissionsRequired_Unknown       CloudAppSecurityProfilePermissionsRequired = "unknown"
	CloudAppSecurityProfilePermissionsRequired_User          CloudAppSecurityProfilePermissionsRequired = "user"
)

func PossibleValuesForCloudAppSecurityProfilePermissionsRequired() []string {
	return []string{
		string(CloudAppSecurityProfilePermissionsRequired_Administrator),
		string(CloudAppSecurityProfilePermissionsRequired_Anonymous),
		string(CloudAppSecurityProfilePermissionsRequired_Guest),
		string(CloudAppSecurityProfilePermissionsRequired_System),
		string(CloudAppSecurityProfilePermissionsRequired_Unknown),
		string(CloudAppSecurityProfilePermissionsRequired_User),
	}
}

func (s *CloudAppSecurityProfilePermissionsRequired) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudAppSecurityProfilePermissionsRequired(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudAppSecurityProfilePermissionsRequired(input string) (*CloudAppSecurityProfilePermissionsRequired, error) {
	vals := map[string]CloudAppSecurityProfilePermissionsRequired{
		"administrator": CloudAppSecurityProfilePermissionsRequired_Administrator,
		"anonymous":     CloudAppSecurityProfilePermissionsRequired_Anonymous,
		"guest":         CloudAppSecurityProfilePermissionsRequired_Guest,
		"system":        CloudAppSecurityProfilePermissionsRequired_System,
		"unknown":       CloudAppSecurityProfilePermissionsRequired_Unknown,
		"user":          CloudAppSecurityProfilePermissionsRequired_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudAppSecurityProfilePermissionsRequired(input)
	return &out, nil
}
