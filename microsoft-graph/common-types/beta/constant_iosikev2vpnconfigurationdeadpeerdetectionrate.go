package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosikEv2VpnConfigurationDeadPeerDetectionRate string

const (
	IosikEv2VpnConfigurationDeadPeerDetectionRate_High   IosikEv2VpnConfigurationDeadPeerDetectionRate = "high"
	IosikEv2VpnConfigurationDeadPeerDetectionRate_Low    IosikEv2VpnConfigurationDeadPeerDetectionRate = "low"
	IosikEv2VpnConfigurationDeadPeerDetectionRate_Medium IosikEv2VpnConfigurationDeadPeerDetectionRate = "medium"
	IosikEv2VpnConfigurationDeadPeerDetectionRate_None   IosikEv2VpnConfigurationDeadPeerDetectionRate = "none"
)

func PossibleValuesForIosikEv2VpnConfigurationDeadPeerDetectionRate() []string {
	return []string{
		string(IosikEv2VpnConfigurationDeadPeerDetectionRate_High),
		string(IosikEv2VpnConfigurationDeadPeerDetectionRate_Low),
		string(IosikEv2VpnConfigurationDeadPeerDetectionRate_Medium),
		string(IosikEv2VpnConfigurationDeadPeerDetectionRate_None),
	}
}

func (s *IosikEv2VpnConfigurationDeadPeerDetectionRate) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseIosikEv2VpnConfigurationDeadPeerDetectionRate(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseIosikEv2VpnConfigurationDeadPeerDetectionRate(input string) (*IosikEv2VpnConfigurationDeadPeerDetectionRate, error) {
	vals := map[string]IosikEv2VpnConfigurationDeadPeerDetectionRate{
		"high":   IosikEv2VpnConfigurationDeadPeerDetectionRate_High,
		"low":    IosikEv2VpnConfigurationDeadPeerDetectionRate_Low,
		"medium": IosikEv2VpnConfigurationDeadPeerDetectionRate_Medium,
		"none":   IosikEv2VpnConfigurationDeadPeerDetectionRate_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosikEv2VpnConfigurationDeadPeerDetectionRate(input)
	return &out, nil
}
