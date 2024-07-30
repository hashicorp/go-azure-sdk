package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftTunnelServerTunnelServerHealthStatus string

const (
	MicrosoftTunnelServerTunnelServerHealthStatus_Healthy           MicrosoftTunnelServerTunnelServerHealthStatus = "healthy"
	MicrosoftTunnelServerTunnelServerHealthStatus_Offline           MicrosoftTunnelServerTunnelServerHealthStatus = "offline"
	MicrosoftTunnelServerTunnelServerHealthStatus_Unhealthy         MicrosoftTunnelServerTunnelServerHealthStatus = "unhealthy"
	MicrosoftTunnelServerTunnelServerHealthStatus_Unknown           MicrosoftTunnelServerTunnelServerHealthStatus = "unknown"
	MicrosoftTunnelServerTunnelServerHealthStatus_UpgradeFailed     MicrosoftTunnelServerTunnelServerHealthStatus = "upgradeFailed"
	MicrosoftTunnelServerTunnelServerHealthStatus_UpgradeInProgress MicrosoftTunnelServerTunnelServerHealthStatus = "upgradeInProgress"
	MicrosoftTunnelServerTunnelServerHealthStatus_Warning           MicrosoftTunnelServerTunnelServerHealthStatus = "warning"
)

func PossibleValuesForMicrosoftTunnelServerTunnelServerHealthStatus() []string {
	return []string{
		string(MicrosoftTunnelServerTunnelServerHealthStatus_Healthy),
		string(MicrosoftTunnelServerTunnelServerHealthStatus_Offline),
		string(MicrosoftTunnelServerTunnelServerHealthStatus_Unhealthy),
		string(MicrosoftTunnelServerTunnelServerHealthStatus_Unknown),
		string(MicrosoftTunnelServerTunnelServerHealthStatus_UpgradeFailed),
		string(MicrosoftTunnelServerTunnelServerHealthStatus_UpgradeInProgress),
		string(MicrosoftTunnelServerTunnelServerHealthStatus_Warning),
	}
}

func (s *MicrosoftTunnelServerTunnelServerHealthStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMicrosoftTunnelServerTunnelServerHealthStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMicrosoftTunnelServerTunnelServerHealthStatus(input string) (*MicrosoftTunnelServerTunnelServerHealthStatus, error) {
	vals := map[string]MicrosoftTunnelServerTunnelServerHealthStatus{
		"healthy":           MicrosoftTunnelServerTunnelServerHealthStatus_Healthy,
		"offline":           MicrosoftTunnelServerTunnelServerHealthStatus_Offline,
		"unhealthy":         MicrosoftTunnelServerTunnelServerHealthStatus_Unhealthy,
		"unknown":           MicrosoftTunnelServerTunnelServerHealthStatus_Unknown,
		"upgradefailed":     MicrosoftTunnelServerTunnelServerHealthStatus_UpgradeFailed,
		"upgradeinprogress": MicrosoftTunnelServerTunnelServerHealthStatus_UpgradeInProgress,
		"warning":           MicrosoftTunnelServerTunnelServerHealthStatus_Warning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftTunnelServerTunnelServerHealthStatus(input)
	return &out, nil
}
