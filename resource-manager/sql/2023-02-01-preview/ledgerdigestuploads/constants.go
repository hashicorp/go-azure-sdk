package ledgerdigestuploads

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LedgerDigestUploadsState string

const (
	LedgerDigestUploadsStateDisabled LedgerDigestUploadsState = "Disabled"
	LedgerDigestUploadsStateEnabled  LedgerDigestUploadsState = "Enabled"
)

func PossibleValuesForLedgerDigestUploadsState() []string {
	return []string{
		string(LedgerDigestUploadsStateDisabled),
		string(LedgerDigestUploadsStateEnabled),
	}
}

func (s *LedgerDigestUploadsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLedgerDigestUploadsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLedgerDigestUploadsState(input string) (*LedgerDigestUploadsState, error) {
	vals := map[string]LedgerDigestUploadsState{
		"disabled": LedgerDigestUploadsStateDisabled,
		"enabled":  LedgerDigestUploadsStateEnabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LedgerDigestUploadsState(input)
	return &out, nil
}
