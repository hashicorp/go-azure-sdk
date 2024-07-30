package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceJoinType string

const (
	ManagedDeviceJoinType_AzureADJoined       ManagedDeviceJoinType = "azureADJoined"
	ManagedDeviceJoinType_AzureADRegistered   ManagedDeviceJoinType = "azureADRegistered"
	ManagedDeviceJoinType_HybridAzureADJoined ManagedDeviceJoinType = "hybridAzureADJoined"
	ManagedDeviceJoinType_Unknown             ManagedDeviceJoinType = "unknown"
)

func PossibleValuesForManagedDeviceJoinType() []string {
	return []string{
		string(ManagedDeviceJoinType_AzureADJoined),
		string(ManagedDeviceJoinType_AzureADRegistered),
		string(ManagedDeviceJoinType_HybridAzureADJoined),
		string(ManagedDeviceJoinType_Unknown),
	}
}

func (s *ManagedDeviceJoinType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedDeviceJoinType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedDeviceJoinType(input string) (*ManagedDeviceJoinType, error) {
	vals := map[string]ManagedDeviceJoinType{
		"azureadjoined":       ManagedDeviceJoinType_AzureADJoined,
		"azureadregistered":   ManagedDeviceJoinType_AzureADRegistered,
		"hybridazureadjoined": ManagedDeviceJoinType_HybridAzureADJoined,
		"unknown":             ManagedDeviceJoinType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceJoinType(input)
	return &out, nil
}
