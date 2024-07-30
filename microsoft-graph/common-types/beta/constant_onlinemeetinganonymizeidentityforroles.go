package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAnonymizeIdentityForRoles string

const (
	OnlineMeetingAnonymizeIdentityForRoles_Attendee    OnlineMeetingAnonymizeIdentityForRoles = "attendee"
	OnlineMeetingAnonymizeIdentityForRoles_Coorganizer OnlineMeetingAnonymizeIdentityForRoles = "coorganizer"
	OnlineMeetingAnonymizeIdentityForRoles_Presenter   OnlineMeetingAnonymizeIdentityForRoles = "presenter"
	OnlineMeetingAnonymizeIdentityForRoles_Producer    OnlineMeetingAnonymizeIdentityForRoles = "producer"
)

func PossibleValuesForOnlineMeetingAnonymizeIdentityForRoles() []string {
	return []string{
		string(OnlineMeetingAnonymizeIdentityForRoles_Attendee),
		string(OnlineMeetingAnonymizeIdentityForRoles_Coorganizer),
		string(OnlineMeetingAnonymizeIdentityForRoles_Presenter),
		string(OnlineMeetingAnonymizeIdentityForRoles_Producer),
	}
}

func (s *OnlineMeetingAnonymizeIdentityForRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingAnonymizeIdentityForRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingAnonymizeIdentityForRoles(input string) (*OnlineMeetingAnonymizeIdentityForRoles, error) {
	vals := map[string]OnlineMeetingAnonymizeIdentityForRoles{
		"attendee":    OnlineMeetingAnonymizeIdentityForRoles_Attendee,
		"coorganizer": OnlineMeetingAnonymizeIdentityForRoles_Coorganizer,
		"presenter":   OnlineMeetingAnonymizeIdentityForRoles_Presenter,
		"producer":    OnlineMeetingAnonymizeIdentityForRoles_Producer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingAnonymizeIdentityForRoles(input)
	return &out, nil
}
