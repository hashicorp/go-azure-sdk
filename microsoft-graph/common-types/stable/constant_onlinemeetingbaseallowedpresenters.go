package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnlineMeetingBaseAllowedPresenters string

const (
	OnlineMeetingBaseAllowedPresenters_Everyone        OnlineMeetingBaseAllowedPresenters = "everyone"
	OnlineMeetingBaseAllowedPresenters_Organization    OnlineMeetingBaseAllowedPresenters = "organization"
	OnlineMeetingBaseAllowedPresenters_Organizer       OnlineMeetingBaseAllowedPresenters = "organizer"
	OnlineMeetingBaseAllowedPresenters_RoleIsPresenter OnlineMeetingBaseAllowedPresenters = "roleIsPresenter"
)

func PossibleValuesForOnlineMeetingBaseAllowedPresenters() []string {
	return []string{
		string(OnlineMeetingBaseAllowedPresenters_Everyone),
		string(OnlineMeetingBaseAllowedPresenters_Organization),
		string(OnlineMeetingBaseAllowedPresenters_Organizer),
		string(OnlineMeetingBaseAllowedPresenters_RoleIsPresenter),
	}
}

func (s *OnlineMeetingBaseAllowedPresenters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseOnlineMeetingBaseAllowedPresenters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseOnlineMeetingBaseAllowedPresenters(input string) (*OnlineMeetingBaseAllowedPresenters, error) {
	vals := map[string]OnlineMeetingBaseAllowedPresenters{
		"everyone":        OnlineMeetingBaseAllowedPresenters_Everyone,
		"organization":    OnlineMeetingBaseAllowedPresenters_Organization,
		"organizer":       OnlineMeetingBaseAllowedPresenters_Organizer,
		"roleispresenter": OnlineMeetingBaseAllowedPresenters_RoleIsPresenter,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnlineMeetingBaseAllowedPresenters(input)
	return &out, nil
}
