package ueinformationlist

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RatType string

const (
	RatTypeFiveG RatType = "5G"
	RatTypeFourG RatType = "4G"
)

func PossibleValuesForRatType() []string {
	return []string{
		string(RatTypeFiveG),
		string(RatTypeFourG),
	}
}

func (s *RatType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRatType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRatType(input string) (*RatType, error) {
	vals := map[string]RatType{
		"5g": RatTypeFiveG,
		"4g": RatTypeFourG,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RatType(input)
	return &out, nil
}

type UeState string

const (
	UeStateConnected    UeState = "Connected"
	UeStateDeregistered UeState = "Deregistered"
	UeStateDetached     UeState = "Detached"
	UeStateIdle         UeState = "Idle"
	UeStateUnknown      UeState = "Unknown"
)

func PossibleValuesForUeState() []string {
	return []string{
		string(UeStateConnected),
		string(UeStateDeregistered),
		string(UeStateDetached),
		string(UeStateIdle),
		string(UeStateUnknown),
	}
}

func (s *UeState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUeState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUeState(input string) (*UeState, error) {
	vals := map[string]UeState{
		"connected":    UeStateConnected,
		"deregistered": UeStateDeregistered,
		"detached":     UeStateDetached,
		"idle":         UeStateIdle,
		"unknown":      UeStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UeState(input)
	return &out, nil
}
