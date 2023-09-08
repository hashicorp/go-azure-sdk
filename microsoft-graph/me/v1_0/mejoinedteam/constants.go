package mejoinedteam

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeJoinedTeamByIdCloneRequestPartsToClone string

const (
	CreateMeJoinedTeamByIdCloneRequestPartsToCloneapps     CreateMeJoinedTeamByIdCloneRequestPartsToClone = "Apps"
	CreateMeJoinedTeamByIdCloneRequestPartsToClonechannels CreateMeJoinedTeamByIdCloneRequestPartsToClone = "Channels"
	CreateMeJoinedTeamByIdCloneRequestPartsToClonemembers  CreateMeJoinedTeamByIdCloneRequestPartsToClone = "Members"
	CreateMeJoinedTeamByIdCloneRequestPartsToClonesettings CreateMeJoinedTeamByIdCloneRequestPartsToClone = "Settings"
	CreateMeJoinedTeamByIdCloneRequestPartsToClonetabs     CreateMeJoinedTeamByIdCloneRequestPartsToClone = "Tabs"
)

func PossibleValuesForCreateMeJoinedTeamByIdCloneRequestPartsToClone() []string {
	return []string{
		string(CreateMeJoinedTeamByIdCloneRequestPartsToCloneapps),
		string(CreateMeJoinedTeamByIdCloneRequestPartsToClonechannels),
		string(CreateMeJoinedTeamByIdCloneRequestPartsToClonemembers),
		string(CreateMeJoinedTeamByIdCloneRequestPartsToClonesettings),
		string(CreateMeJoinedTeamByIdCloneRequestPartsToClonetabs),
	}
}

func (s *CreateMeJoinedTeamByIdCloneRequestPartsToClone) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeJoinedTeamByIdCloneRequestPartsToClone(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeJoinedTeamByIdCloneRequestPartsToClone(input string) (*CreateMeJoinedTeamByIdCloneRequestPartsToClone, error) {
	vals := map[string]CreateMeJoinedTeamByIdCloneRequestPartsToClone{
		"apps":     CreateMeJoinedTeamByIdCloneRequestPartsToCloneapps,
		"channels": CreateMeJoinedTeamByIdCloneRequestPartsToClonechannels,
		"members":  CreateMeJoinedTeamByIdCloneRequestPartsToClonemembers,
		"settings": CreateMeJoinedTeamByIdCloneRequestPartsToClonesettings,
		"tabs":     CreateMeJoinedTeamByIdCloneRequestPartsToClonetabs,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeJoinedTeamByIdCloneRequestPartsToClone(input)
	return &out, nil
}

type CreateMeJoinedTeamByIdCloneRequestVisibility string

const (
	CreateMeJoinedTeamByIdCloneRequestVisibilityhiddenMembership CreateMeJoinedTeamByIdCloneRequestVisibility = "HiddenMembership"
	CreateMeJoinedTeamByIdCloneRequestVisibilityprivate          CreateMeJoinedTeamByIdCloneRequestVisibility = "Private"
	CreateMeJoinedTeamByIdCloneRequestVisibilitypublic           CreateMeJoinedTeamByIdCloneRequestVisibility = "Public"
)

func PossibleValuesForCreateMeJoinedTeamByIdCloneRequestVisibility() []string {
	return []string{
		string(CreateMeJoinedTeamByIdCloneRequestVisibilityhiddenMembership),
		string(CreateMeJoinedTeamByIdCloneRequestVisibilityprivate),
		string(CreateMeJoinedTeamByIdCloneRequestVisibilitypublic),
	}
}

func (s *CreateMeJoinedTeamByIdCloneRequestVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateMeJoinedTeamByIdCloneRequestVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateMeJoinedTeamByIdCloneRequestVisibility(input string) (*CreateMeJoinedTeamByIdCloneRequestVisibility, error) {
	vals := map[string]CreateMeJoinedTeamByIdCloneRequestVisibility{
		"hiddenmembership": CreateMeJoinedTeamByIdCloneRequestVisibilityhiddenMembership,
		"private":          CreateMeJoinedTeamByIdCloneRequestVisibilityprivate,
		"public":           CreateMeJoinedTeamByIdCloneRequestVisibilitypublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateMeJoinedTeamByIdCloneRequestVisibility(input)
	return &out, nil
}
