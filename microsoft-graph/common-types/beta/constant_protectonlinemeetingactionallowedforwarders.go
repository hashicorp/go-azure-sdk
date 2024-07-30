package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProtectOnlineMeetingActionAllowedForwarders string

const (
	ProtectOnlineMeetingActionAllowedForwarders_Everyone  ProtectOnlineMeetingActionAllowedForwarders = "everyone"
	ProtectOnlineMeetingActionAllowedForwarders_Organizer ProtectOnlineMeetingActionAllowedForwarders = "organizer"
)

func PossibleValuesForProtectOnlineMeetingActionAllowedForwarders() []string {
	return []string{
		string(ProtectOnlineMeetingActionAllowedForwarders_Everyone),
		string(ProtectOnlineMeetingActionAllowedForwarders_Organizer),
	}
}

func (s *ProtectOnlineMeetingActionAllowedForwarders) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProtectOnlineMeetingActionAllowedForwarders(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProtectOnlineMeetingActionAllowedForwarders(input string) (*ProtectOnlineMeetingActionAllowedForwarders, error) {
	vals := map[string]ProtectOnlineMeetingActionAllowedForwarders{
		"everyone":  ProtectOnlineMeetingActionAllowedForwarders_Everyone,
		"organizer": ProtectOnlineMeetingActionAllowedForwarders_Organizer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProtectOnlineMeetingActionAllowedForwarders(input)
	return &out, nil
}
