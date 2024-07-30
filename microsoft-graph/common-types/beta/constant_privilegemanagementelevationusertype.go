package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementElevationUserType string

const (
	PrivilegeManagementElevationUserType_AzureAd      PrivilegeManagementElevationUserType = "azureAd"
	PrivilegeManagementElevationUserType_Hybrid       PrivilegeManagementElevationUserType = "hybrid"
	PrivilegeManagementElevationUserType_Local        PrivilegeManagementElevationUserType = "local"
	PrivilegeManagementElevationUserType_Undetermined PrivilegeManagementElevationUserType = "undetermined"
)

func PossibleValuesForPrivilegeManagementElevationUserType() []string {
	return []string{
		string(PrivilegeManagementElevationUserType_AzureAd),
		string(PrivilegeManagementElevationUserType_Hybrid),
		string(PrivilegeManagementElevationUserType_Local),
		string(PrivilegeManagementElevationUserType_Undetermined),
	}
}

func (s *PrivilegeManagementElevationUserType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivilegeManagementElevationUserType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivilegeManagementElevationUserType(input string) (*PrivilegeManagementElevationUserType, error) {
	vals := map[string]PrivilegeManagementElevationUserType{
		"azuread":      PrivilegeManagementElevationUserType_AzureAd,
		"hybrid":       PrivilegeManagementElevationUserType_Hybrid,
		"local":        PrivilegeManagementElevationUserType_Local,
		"undetermined": PrivilegeManagementElevationUserType_Undetermined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivilegeManagementElevationUserType(input)
	return &out, nil
}
