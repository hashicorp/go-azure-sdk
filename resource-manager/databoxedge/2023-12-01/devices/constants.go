package devices

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkAdapterDHCPStatus string

const (
	NetworkAdapterDHCPStatusDisabled NetworkAdapterDHCPStatus = "Disabled"
	NetworkAdapterDHCPStatusEnabled  NetworkAdapterDHCPStatus = "Enabled"
)

func PossibleValuesForNetworkAdapterDHCPStatus() []string {
	return []string{
		string(NetworkAdapterDHCPStatusDisabled),
		string(NetworkAdapterDHCPStatusEnabled),
	}
}

func (s *NetworkAdapterDHCPStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkAdapterDHCPStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkAdapterDHCPStatus(input string) (*NetworkAdapterDHCPStatus, error) {
	vals := map[string]NetworkAdapterDHCPStatus{
		"disabled": NetworkAdapterDHCPStatusDisabled,
		"enabled":  NetworkAdapterDHCPStatusEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkAdapterDHCPStatus(input)
	return &out, nil
}

type NetworkAdapterRDMAStatus string

const (
	NetworkAdapterRDMAStatusCapable   NetworkAdapterRDMAStatus = "Capable"
	NetworkAdapterRDMAStatusIncapable NetworkAdapterRDMAStatus = "Incapable"
)

func PossibleValuesForNetworkAdapterRDMAStatus() []string {
	return []string{
		string(NetworkAdapterRDMAStatusCapable),
		string(NetworkAdapterRDMAStatusIncapable),
	}
}

func (s *NetworkAdapterRDMAStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkAdapterRDMAStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkAdapterRDMAStatus(input string) (*NetworkAdapterRDMAStatus, error) {
	vals := map[string]NetworkAdapterRDMAStatus{
		"capable":   NetworkAdapterRDMAStatusCapable,
		"incapable": NetworkAdapterRDMAStatusIncapable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkAdapterRDMAStatus(input)
	return &out, nil
}

type NetworkAdapterStatus string

const (
	NetworkAdapterStatusActive   NetworkAdapterStatus = "Active"
	NetworkAdapterStatusInactive NetworkAdapterStatus = "Inactive"
)

func PossibleValuesForNetworkAdapterStatus() []string {
	return []string{
		string(NetworkAdapterStatusActive),
		string(NetworkAdapterStatusInactive),
	}
}

func (s *NetworkAdapterStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkAdapterStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkAdapterStatus(input string) (*NetworkAdapterStatus, error) {
	vals := map[string]NetworkAdapterStatus{
		"active":   NetworkAdapterStatusActive,
		"inactive": NetworkAdapterStatusInactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkAdapterStatus(input)
	return &out, nil
}

type NetworkGroup string

const (
	NetworkGroupNonRDMA NetworkGroup = "NonRDMA"
	NetworkGroupNone    NetworkGroup = "None"
	NetworkGroupRDMA    NetworkGroup = "RDMA"
)

func PossibleValuesForNetworkGroup() []string {
	return []string{
		string(NetworkGroupNonRDMA),
		string(NetworkGroupNone),
		string(NetworkGroupRDMA),
	}
}

func (s *NetworkGroup) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseNetworkGroup(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseNetworkGroup(input string) (*NetworkGroup, error) {
	vals := map[string]NetworkGroup{
		"nonrdma": NetworkGroupNonRDMA,
		"none":    NetworkGroupNone,
		"rdma":    NetworkGroupRDMA,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NetworkGroup(input)
	return &out, nil
}
