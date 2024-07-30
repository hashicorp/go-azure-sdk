package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingAllowedPresenters string

const (
	OnlineMeetingAllowedPresenters_Everyone        OnlineMeetingAllowedPresenters = "everyone"
	OnlineMeetingAllowedPresenters_Organization    OnlineMeetingAllowedPresenters = "organization"
	OnlineMeetingAllowedPresenters_Organizer       OnlineMeetingAllowedPresenters = "organizer"
	OnlineMeetingAllowedPresenters_RoleIsPresenter OnlineMeetingAllowedPresenters = "roleIsPresenter"
)

func PossibleValuesForOnlineMeetingAllowedPresenters() []string {
	return []string{
		string(OnlineMeetingAllowedPresenters_Everyone),
		string(OnlineMeetingAllowedPresenters_Organization),
		string(OnlineMeetingAllowedPresenters_Organizer),
		string(OnlineMeetingAllowedPresenters_RoleIsPresenter),
	}
}

func (s *OnlineMeetingAllowedPresenters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingAllowedPresenters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingAllowedPresenters(input string) (*OnlineMeetingAllowedPresenters, error) {
	vals := map[string]OnlineMeetingAllowedPresenters{
		"everyone":        OnlineMeetingAllowedPresenters_Everyone,
		"organization":    OnlineMeetingAllowedPresenters_Organization,
		"organizer":       OnlineMeetingAllowedPresenters_Organizer,
		"roleispresenter": OnlineMeetingAllowedPresenters_RoleIsPresenter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingAllowedPresenters(input)
	return &out, nil
}
