package nameavailability

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Reason string

const (
	ReasonAlreadyExists Reason = "AlreadyExists"
)

func PossibleValuesForReason() []string {
	return []string{
		string(ReasonAlreadyExists),
	}
}

func (s *Reason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseReason(input string) (*Reason, error) {
	vals := map[string]Reason{
		"alreadyexists": ReasonAlreadyExists,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Reason(input)
	return &out, nil
}

type Type string

const (
	TypeMicrosoftPointVideoIndexerAccounts Type = "Microsoft.VideoIndexer/accounts"
)

func PossibleValuesForType() []string {
	return []string{
		string(TypeMicrosoftPointVideoIndexerAccounts),
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
		"microsoft.videoindexer/accounts": TypeMicrosoftPointVideoIndexerAccounts,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Type(input)
	return &out, nil
}
