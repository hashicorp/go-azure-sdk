package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingBaseAnonymizeIdentityForRoles string

const (
	OnlineMeetingBaseAnonymizeIdentityForRoles_Attendee    OnlineMeetingBaseAnonymizeIdentityForRoles = "attendee"
	OnlineMeetingBaseAnonymizeIdentityForRoles_Coorganizer OnlineMeetingBaseAnonymizeIdentityForRoles = "coorganizer"
	OnlineMeetingBaseAnonymizeIdentityForRoles_Presenter   OnlineMeetingBaseAnonymizeIdentityForRoles = "presenter"
	OnlineMeetingBaseAnonymizeIdentityForRoles_Producer    OnlineMeetingBaseAnonymizeIdentityForRoles = "producer"
)

func PossibleValuesForOnlineMeetingBaseAnonymizeIdentityForRoles() []string {
	return []string{
		string(OnlineMeetingBaseAnonymizeIdentityForRoles_Attendee),
		string(OnlineMeetingBaseAnonymizeIdentityForRoles_Coorganizer),
		string(OnlineMeetingBaseAnonymizeIdentityForRoles_Presenter),
		string(OnlineMeetingBaseAnonymizeIdentityForRoles_Producer),
	}
}

func (s *OnlineMeetingBaseAnonymizeIdentityForRoles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingBaseAnonymizeIdentityForRoles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingBaseAnonymizeIdentityForRoles(input string) (*OnlineMeetingBaseAnonymizeIdentityForRoles, error) {
	vals := map[string]OnlineMeetingBaseAnonymizeIdentityForRoles{
		"attendee":    OnlineMeetingBaseAnonymizeIdentityForRoles_Attendee,
		"coorganizer": OnlineMeetingBaseAnonymizeIdentityForRoles_Coorganizer,
		"presenter":   OnlineMeetingBaseAnonymizeIdentityForRoles_Presenter,
		"producer":    OnlineMeetingBaseAnonymizeIdentityForRoles_Producer,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingBaseAnonymizeIdentityForRoles(input)
	return &out, nil
}
