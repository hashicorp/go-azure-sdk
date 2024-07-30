package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerTeamsPublicationInfoCreationSourceKind string

const (
	PlannerTeamsPublicationInfoCreationSourceKind_External    PlannerTeamsPublicationInfoCreationSourceKind = "external"
	PlannerTeamsPublicationInfoCreationSourceKind_None        PlannerTeamsPublicationInfoCreationSourceKind = "none"
	PlannerTeamsPublicationInfoCreationSourceKind_Publication PlannerTeamsPublicationInfoCreationSourceKind = "publication"
)

func PossibleValuesForPlannerTeamsPublicationInfoCreationSourceKind() []string {
	return []string{
		string(PlannerTeamsPublicationInfoCreationSourceKind_External),
		string(PlannerTeamsPublicationInfoCreationSourceKind_None),
		string(PlannerTeamsPublicationInfoCreationSourceKind_Publication),
	}
}

func (s *PlannerTeamsPublicationInfoCreationSourceKind) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlannerTeamsPublicationInfoCreationSourceKind(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlannerTeamsPublicationInfoCreationSourceKind(input string) (*PlannerTeamsPublicationInfoCreationSourceKind, error) {
	vals := map[string]PlannerTeamsPublicationInfoCreationSourceKind{
		"external":    PlannerTeamsPublicationInfoCreationSourceKind_External,
		"none":        PlannerTeamsPublicationInfoCreationSourceKind_None,
		"publication": PlannerTeamsPublicationInfoCreationSourceKind_Publication,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlannerTeamsPublicationInfoCreationSourceKind(input)
	return &out, nil
}
