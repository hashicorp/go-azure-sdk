package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOrganizationSettingsUserAccountType string

const (
	CloudPCOrganizationSettingsUserAccountType_Administrator CloudPCOrganizationSettingsUserAccountType = "administrator"
	CloudPCOrganizationSettingsUserAccountType_StandardUser  CloudPCOrganizationSettingsUserAccountType = "standardUser"
)

func PossibleValuesForCloudPCOrganizationSettingsUserAccountType() []string {
	return []string{
		string(CloudPCOrganizationSettingsUserAccountType_Administrator),
		string(CloudPCOrganizationSettingsUserAccountType_StandardUser),
	}
}

func (s *CloudPCOrganizationSettingsUserAccountType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCloudPCOrganizationSettingsUserAccountType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCloudPCOrganizationSettingsUserAccountType(input string) (*CloudPCOrganizationSettingsUserAccountType, error) {
	vals := map[string]CloudPCOrganizationSettingsUserAccountType{
		"administrator": CloudPCOrganizationSettingsUserAccountType_Administrator,
		"standarduser":  CloudPCOrganizationSettingsUserAccountType_StandardUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CloudPCOrganizationSettingsUserAccountType(input)
	return &out, nil
}
