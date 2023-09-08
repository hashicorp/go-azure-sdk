package groupteam

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateGroupByIdTeamCloneRequestPartsToClone string

const (
	CreateGroupByIdTeamCloneRequestPartsToCloneapps     CreateGroupByIdTeamCloneRequestPartsToClone = "Apps"
	CreateGroupByIdTeamCloneRequestPartsToClonechannels CreateGroupByIdTeamCloneRequestPartsToClone = "Channels"
	CreateGroupByIdTeamCloneRequestPartsToClonemembers  CreateGroupByIdTeamCloneRequestPartsToClone = "Members"
	CreateGroupByIdTeamCloneRequestPartsToClonesettings CreateGroupByIdTeamCloneRequestPartsToClone = "Settings"
	CreateGroupByIdTeamCloneRequestPartsToClonetabs     CreateGroupByIdTeamCloneRequestPartsToClone = "Tabs"
)

func PossibleValuesForCreateGroupByIdTeamCloneRequestPartsToClone() []string {
	return []string{
		string(CreateGroupByIdTeamCloneRequestPartsToCloneapps),
		string(CreateGroupByIdTeamCloneRequestPartsToClonechannels),
		string(CreateGroupByIdTeamCloneRequestPartsToClonemembers),
		string(CreateGroupByIdTeamCloneRequestPartsToClonesettings),
		string(CreateGroupByIdTeamCloneRequestPartsToClonetabs),
	}
}

func (s *CreateGroupByIdTeamCloneRequestPartsToClone) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateGroupByIdTeamCloneRequestPartsToClone(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateGroupByIdTeamCloneRequestPartsToClone(input string) (*CreateGroupByIdTeamCloneRequestPartsToClone, error) {
	vals := map[string]CreateGroupByIdTeamCloneRequestPartsToClone{
		"apps":     CreateGroupByIdTeamCloneRequestPartsToCloneapps,
		"channels": CreateGroupByIdTeamCloneRequestPartsToClonechannels,
		"members":  CreateGroupByIdTeamCloneRequestPartsToClonemembers,
		"settings": CreateGroupByIdTeamCloneRequestPartsToClonesettings,
		"tabs":     CreateGroupByIdTeamCloneRequestPartsToClonetabs,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateGroupByIdTeamCloneRequestPartsToClone(input)
	return &out, nil
}

type CreateGroupByIdTeamCloneRequestVisibility string

const (
	CreateGroupByIdTeamCloneRequestVisibilityhiddenMembership CreateGroupByIdTeamCloneRequestVisibility = "HiddenMembership"
	CreateGroupByIdTeamCloneRequestVisibilityprivate          CreateGroupByIdTeamCloneRequestVisibility = "Private"
	CreateGroupByIdTeamCloneRequestVisibilitypublic           CreateGroupByIdTeamCloneRequestVisibility = "Public"
)

func PossibleValuesForCreateGroupByIdTeamCloneRequestVisibility() []string {
	return []string{
		string(CreateGroupByIdTeamCloneRequestVisibilityhiddenMembership),
		string(CreateGroupByIdTeamCloneRequestVisibilityprivate),
		string(CreateGroupByIdTeamCloneRequestVisibilitypublic),
	}
}

func (s *CreateGroupByIdTeamCloneRequestVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateGroupByIdTeamCloneRequestVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateGroupByIdTeamCloneRequestVisibility(input string) (*CreateGroupByIdTeamCloneRequestVisibility, error) {
	vals := map[string]CreateGroupByIdTeamCloneRequestVisibility{
		"hiddenmembership": CreateGroupByIdTeamCloneRequestVisibilityhiddenMembership,
		"private":          CreateGroupByIdTeamCloneRequestVisibilityprivate,
		"public":           CreateGroupByIdTeamCloneRequestVisibilitypublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateGroupByIdTeamCloneRequestVisibility(input)
	return &out, nil
}
