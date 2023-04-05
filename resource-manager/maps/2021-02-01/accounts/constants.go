package accounts

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyType string

const (
	KeyTypePrimary   KeyType = "primary"
	KeyTypeSecondary KeyType = "secondary"
)

func PossibleValuesForKeyType() []string {
	return []string{
		string(KeyTypePrimary),
		string(KeyTypeSecondary),
	}
}

func (s *KeyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForKeyType() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = KeyType(decoded)
	return nil
}

type Kind string

const (
	KindGenOne Kind = "Gen1"
	KindGenTwo Kind = "Gen2"
)

func PossibleValuesForKind() []string {
	return []string{
		string(KindGenOne),
		string(KindGenTwo),
	}
}

func (s *Kind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForKind() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = Kind(decoded)
	return nil
}

type Name string

const (
	NameGTwo  Name = "G2"
	NameSOne  Name = "S1"
	NameSZero Name = "S0"
)

func PossibleValuesForName() []string {
	return []string{
		string(NameGTwo),
		string(NameSOne),
		string(NameSZero),
	}
}

func (s *Name) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	for _, v := range PossibleValuesForName() {
		if strings.EqualFold(v, decoded) {
			decoded = v
			break
		}
	}
	*s = Name(decoded)
	return nil
}
