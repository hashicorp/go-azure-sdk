package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TypedEmailAddressType string

const (
	TypedEmailAddressType_Main     TypedEmailAddressType = "main"
	TypedEmailAddressType_Other    TypedEmailAddressType = "other"
	TypedEmailAddressType_Personal TypedEmailAddressType = "personal"
	TypedEmailAddressType_Unknown  TypedEmailAddressType = "unknown"
	TypedEmailAddressType_Work     TypedEmailAddressType = "work"
)

func PossibleValuesForTypedEmailAddressType() []string {
	return []string{
		string(TypedEmailAddressType_Main),
		string(TypedEmailAddressType_Other),
		string(TypedEmailAddressType_Personal),
		string(TypedEmailAddressType_Unknown),
		string(TypedEmailAddressType_Work),
	}
}

func (s *TypedEmailAddressType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTypedEmailAddressType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTypedEmailAddressType(input string) (*TypedEmailAddressType, error) {
	vals := map[string]TypedEmailAddressType{
		"main":     TypedEmailAddressType_Main,
		"other":    TypedEmailAddressType_Other,
		"personal": TypedEmailAddressType_Personal,
		"unknown":  TypedEmailAddressType_Unknown,
		"work":     TypedEmailAddressType_Work,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TypedEmailAddressType(input)
	return &out, nil
}
