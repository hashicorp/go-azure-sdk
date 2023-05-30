package changedatacapture

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionType string

const (
	ConnectionTypeLinkedservicetype ConnectionType = "linkedservicetype"
)

func PossibleValuesForConnectionType() []string {
	return []string{
		string(ConnectionTypeLinkedservicetype),
	}
}

func (s *ConnectionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseConnectionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseConnectionType(input string) (*ConnectionType, error) {
	vals := map[string]ConnectionType{
		"linkedservicetype": ConnectionTypeLinkedservicetype,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConnectionType(input)
	return &out, nil
}

type FrequencyType string

const (
	FrequencyTypeHour   FrequencyType = "Hour"
	FrequencyTypeMinute FrequencyType = "Minute"
	FrequencyTypeSecond FrequencyType = "Second"
)

func PossibleValuesForFrequencyType() []string {
	return []string{
		string(FrequencyTypeHour),
		string(FrequencyTypeMinute),
		string(FrequencyTypeSecond),
	}
}

func (s *FrequencyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFrequencyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFrequencyType(input string) (*FrequencyType, error) {
	vals := map[string]FrequencyType{
		"hour":   FrequencyTypeHour,
		"minute": FrequencyTypeMinute,
		"second": FrequencyTypeSecond,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FrequencyType(input)
	return &out, nil
}

type MappingType string

const (
	MappingTypeAggregate MappingType = "Aggregate"
	MappingTypeDerived   MappingType = "Derived"
	MappingTypeDirect    MappingType = "Direct"
)

func PossibleValuesForMappingType() []string {
	return []string{
		string(MappingTypeAggregate),
		string(MappingTypeDerived),
		string(MappingTypeDirect),
	}
}

func (s *MappingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMappingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMappingType(input string) (*MappingType, error) {
	vals := map[string]MappingType{
		"aggregate": MappingTypeAggregate,
		"derived":   MappingTypeDerived,
		"direct":    MappingTypeDirect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MappingType(input)
	return &out, nil
}

type Type string

const (
	TypeLinkedServiceReference Type = "LinkedServiceReference"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeLinkedServiceReference),
	}
}

func (s *Type) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseType(input string) (*Type, error) {
	vals := map[string]Type{
		"linkedservicereference": TypeLinkedServiceReference,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
