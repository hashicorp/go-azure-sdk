package managedledgerdigestuploads

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedLedgerDigestUploadsState string

const (
	ManagedLedgerDigestUploadsStateDisabled ManagedLedgerDigestUploadsState = "Disabled"
	ManagedLedgerDigestUploadsStateEnabled  ManagedLedgerDigestUploadsState = "Enabled"
)

func PossibleValuesForManagedLedgerDigestUploadsState() []string {
	return []string{
		string(ManagedLedgerDigestUploadsStateDisabled),
		string(ManagedLedgerDigestUploadsStateEnabled),
	}
}

func (s *ManagedLedgerDigestUploadsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedLedgerDigestUploadsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedLedgerDigestUploadsState(input string) (*ManagedLedgerDigestUploadsState, error) {
	vals := map[string]ManagedLedgerDigestUploadsState{
		"disabled": ManagedLedgerDigestUploadsStateDisabled,
		"enabled":  ManagedLedgerDigestUploadsStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedLedgerDigestUploadsState(input)
	return &out, nil
}
