package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceNowConnectionServiceNowConnectionStatus string

const (
	ServiceNowConnectionServiceNowConnectionStatus_Disabled ServiceNowConnectionServiceNowConnectionStatus = "disabled"
	ServiceNowConnectionServiceNowConnectionStatus_Enabled  ServiceNowConnectionServiceNowConnectionStatus = "enabled"
)

func PossibleValuesForServiceNowConnectionServiceNowConnectionStatus() []string {
	return []string{
		string(ServiceNowConnectionServiceNowConnectionStatus_Disabled),
		string(ServiceNowConnectionServiceNowConnectionStatus_Enabled),
	}
}

func (s *ServiceNowConnectionServiceNowConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseServiceNowConnectionServiceNowConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseServiceNowConnectionServiceNowConnectionStatus(input string) (*ServiceNowConnectionServiceNowConnectionStatus, error) {
	vals := map[string]ServiceNowConnectionServiceNowConnectionStatus{
		"disabled": ServiceNowConnectionServiceNowConnectionStatus_Disabled,
		"enabled":  ServiceNowConnectionServiceNowConnectionStatus_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ServiceNowConnectionServiceNowConnectionStatus(input)
	return &out, nil
}
