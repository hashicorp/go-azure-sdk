package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagedDeviceJoinType string

const (
	WindowsManagedDeviceJoinType_AzureADJoined       WindowsManagedDeviceJoinType = "azureADJoined"
	WindowsManagedDeviceJoinType_AzureADRegistered   WindowsManagedDeviceJoinType = "azureADRegistered"
	WindowsManagedDeviceJoinType_HybridAzureADJoined WindowsManagedDeviceJoinType = "hybridAzureADJoined"
	WindowsManagedDeviceJoinType_Unknown             WindowsManagedDeviceJoinType = "unknown"
)

func PossibleValuesForWindowsManagedDeviceJoinType() []string {
	return []string{
		string(WindowsManagedDeviceJoinType_AzureADJoined),
		string(WindowsManagedDeviceJoinType_AzureADRegistered),
		string(WindowsManagedDeviceJoinType_HybridAzureADJoined),
		string(WindowsManagedDeviceJoinType_Unknown),
	}
}

func (s *WindowsManagedDeviceJoinType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindowsManagedDeviceJoinType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindowsManagedDeviceJoinType(input string) (*WindowsManagedDeviceJoinType, error) {
	vals := map[string]WindowsManagedDeviceJoinType{
		"azureadjoined":       WindowsManagedDeviceJoinType_AzureADJoined,
		"azureadregistered":   WindowsManagedDeviceJoinType_AzureADRegistered,
		"hybridazureadjoined": WindowsManagedDeviceJoinType_HybridAzureADJoined,
		"unknown":             WindowsManagedDeviceJoinType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsManagedDeviceJoinType(input)
	return &out, nil
}
