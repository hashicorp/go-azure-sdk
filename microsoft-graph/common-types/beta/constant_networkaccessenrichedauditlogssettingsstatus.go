package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessEnrichedAuditLogsSettingsStatus string

const (
	NetworkaccessEnrichedAuditLogsSettingsStatus_Disabled NetworkaccessEnrichedAuditLogsSettingsStatus = "disabled"
	NetworkaccessEnrichedAuditLogsSettingsStatus_Enabled  NetworkaccessEnrichedAuditLogsSettingsStatus = "enabled"
)

func PossibleValuesForNetworkaccessEnrichedAuditLogsSettingsStatus() []string {
	return []string{
		string(NetworkaccessEnrichedAuditLogsSettingsStatus_Disabled),
		string(NetworkaccessEnrichedAuditLogsSettingsStatus_Enabled),
	}
}

func (s *NetworkaccessEnrichedAuditLogsSettingsStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkaccessEnrichedAuditLogsSettingsStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkaccessEnrichedAuditLogsSettingsStatus(input string) (*NetworkaccessEnrichedAuditLogsSettingsStatus, error) {
	vals := map[string]NetworkaccessEnrichedAuditLogsSettingsStatus{
		"disabled": NetworkaccessEnrichedAuditLogsSettingsStatus_Disabled,
		"enabled":  NetworkaccessEnrichedAuditLogsSettingsStatus_Enabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkaccessEnrichedAuditLogsSettingsStatus(input)
	return &out, nil
}
