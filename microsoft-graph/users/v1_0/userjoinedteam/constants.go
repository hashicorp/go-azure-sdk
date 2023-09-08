package userjoinedteam

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone string

const (
	CreateUserByIdJoinedTeamByIdCloneRequestPartsToCloneapps     CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone = "Apps"
	CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonechannels CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone = "Channels"
	CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonemembers  CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone = "Members"
	CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonesettings CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone = "Settings"
	CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonetabs     CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone = "Tabs"
)

func PossibleValuesForCreateUserByIdJoinedTeamByIdCloneRequestPartsToClone() []string {
	return []string{
		string(CreateUserByIdJoinedTeamByIdCloneRequestPartsToCloneapps),
		string(CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonechannels),
		string(CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonemembers),
		string(CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonesettings),
		string(CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonetabs),
	}
}

func (s *CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdJoinedTeamByIdCloneRequestPartsToClone(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdJoinedTeamByIdCloneRequestPartsToClone(input string) (*CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone, error) {
	vals := map[string]CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone{
		"apps":     CreateUserByIdJoinedTeamByIdCloneRequestPartsToCloneapps,
		"channels": CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonechannels,
		"members":  CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonemembers,
		"settings": CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonesettings,
		"tabs":     CreateUserByIdJoinedTeamByIdCloneRequestPartsToClonetabs,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdJoinedTeamByIdCloneRequestPartsToClone(input)
	return &out, nil
}

type CreateUserByIdJoinedTeamByIdCloneRequestVisibility string

const (
	CreateUserByIdJoinedTeamByIdCloneRequestVisibilityhiddenMembership CreateUserByIdJoinedTeamByIdCloneRequestVisibility = "HiddenMembership"
	CreateUserByIdJoinedTeamByIdCloneRequestVisibilityprivate          CreateUserByIdJoinedTeamByIdCloneRequestVisibility = "Private"
	CreateUserByIdJoinedTeamByIdCloneRequestVisibilitypublic           CreateUserByIdJoinedTeamByIdCloneRequestVisibility = "Public"
)

func PossibleValuesForCreateUserByIdJoinedTeamByIdCloneRequestVisibility() []string {
	return []string{
		string(CreateUserByIdJoinedTeamByIdCloneRequestVisibilityhiddenMembership),
		string(CreateUserByIdJoinedTeamByIdCloneRequestVisibilityprivate),
		string(CreateUserByIdJoinedTeamByIdCloneRequestVisibilitypublic),
	}
}

func (s *CreateUserByIdJoinedTeamByIdCloneRequestVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateUserByIdJoinedTeamByIdCloneRequestVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateUserByIdJoinedTeamByIdCloneRequestVisibility(input string) (*CreateUserByIdJoinedTeamByIdCloneRequestVisibility, error) {
	vals := map[string]CreateUserByIdJoinedTeamByIdCloneRequestVisibility{
		"hiddenmembership": CreateUserByIdJoinedTeamByIdCloneRequestVisibilityhiddenMembership,
		"private":          CreateUserByIdJoinedTeamByIdCloneRequestVisibilityprivate,
		"public":           CreateUserByIdJoinedTeamByIdCloneRequestVisibilitypublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateUserByIdJoinedTeamByIdCloneRequestVisibility(input)
	return &out, nil
}
