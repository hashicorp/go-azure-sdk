package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RejectJoinResponseReason string

const (
	RejectJoinResponseReason_Busy      RejectJoinResponseReason = "busy"
	RejectJoinResponseReason_Forbidden RejectJoinResponseReason = "forbidden"
	RejectJoinResponseReason_None      RejectJoinResponseReason = "none"
)

func PossibleValuesForRejectJoinResponseReason() []string {
	return []string{
		string(RejectJoinResponseReason_Busy),
		string(RejectJoinResponseReason_Forbidden),
		string(RejectJoinResponseReason_None),
	}
}

func (s *RejectJoinResponseReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRejectJoinResponseReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRejectJoinResponseReason(input string) (*RejectJoinResponseReason, error) {
	vals := map[string]RejectJoinResponseReason{
		"busy":      RejectJoinResponseReason_Busy,
		"forbidden": RejectJoinResponseReason_Forbidden,
		"none":      RejectJoinResponseReason_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RejectJoinResponseReason(input)
	return &out, nil
}
