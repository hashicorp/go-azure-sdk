package joinedteam

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateJoinedTeamCloneRequestPartsToClone string

const (
	CreateJoinedTeamCloneRequestPartsToCloneApps     CreateJoinedTeamCloneRequestPartsToClone = "apps"
	CreateJoinedTeamCloneRequestPartsToCloneChannels CreateJoinedTeamCloneRequestPartsToClone = "channels"
	CreateJoinedTeamCloneRequestPartsToCloneMembers  CreateJoinedTeamCloneRequestPartsToClone = "members"
	CreateJoinedTeamCloneRequestPartsToCloneSettings CreateJoinedTeamCloneRequestPartsToClone = "settings"
	CreateJoinedTeamCloneRequestPartsToCloneTabs     CreateJoinedTeamCloneRequestPartsToClone = "tabs"
)

func PossibleValuesForCreateJoinedTeamCloneRequestPartsToClone() []string {
	return []string{
		string(CreateJoinedTeamCloneRequestPartsToCloneApps),
		string(CreateJoinedTeamCloneRequestPartsToCloneChannels),
		string(CreateJoinedTeamCloneRequestPartsToCloneMembers),
		string(CreateJoinedTeamCloneRequestPartsToCloneSettings),
		string(CreateJoinedTeamCloneRequestPartsToCloneTabs),
	}
}

func (s *CreateJoinedTeamCloneRequestPartsToClone) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateJoinedTeamCloneRequestPartsToClone(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateJoinedTeamCloneRequestPartsToClone(input string) (*CreateJoinedTeamCloneRequestPartsToClone, error) {
	vals := map[string]CreateJoinedTeamCloneRequestPartsToClone{
		"apps":     CreateJoinedTeamCloneRequestPartsToCloneApps,
		"channels": CreateJoinedTeamCloneRequestPartsToCloneChannels,
		"members":  CreateJoinedTeamCloneRequestPartsToCloneMembers,
		"settings": CreateJoinedTeamCloneRequestPartsToCloneSettings,
		"tabs":     CreateJoinedTeamCloneRequestPartsToCloneTabs,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateJoinedTeamCloneRequestPartsToClone(input)
	return &out, nil
}

type CreateJoinedTeamCloneRequestVisibility string

const (
	CreateJoinedTeamCloneRequestVisibilityHiddenMembership CreateJoinedTeamCloneRequestVisibility = "hiddenMembership"
	CreateJoinedTeamCloneRequestVisibilityPrivate          CreateJoinedTeamCloneRequestVisibility = "private"
	CreateJoinedTeamCloneRequestVisibilityPublic           CreateJoinedTeamCloneRequestVisibility = "public"
)

func PossibleValuesForCreateJoinedTeamCloneRequestVisibility() []string {
	return []string{
		string(CreateJoinedTeamCloneRequestVisibilityHiddenMembership),
		string(CreateJoinedTeamCloneRequestVisibilityPrivate),
		string(CreateJoinedTeamCloneRequestVisibilityPublic),
	}
}

func (s *CreateJoinedTeamCloneRequestVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateJoinedTeamCloneRequestVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateJoinedTeamCloneRequestVisibility(input string) (*CreateJoinedTeamCloneRequestVisibility, error) {
	vals := map[string]CreateJoinedTeamCloneRequestVisibility{
		"hiddenmembership": CreateJoinedTeamCloneRequestVisibilityHiddenMembership,
		"private":          CreateJoinedTeamCloneRequestVisibilityPrivate,
		"public":           CreateJoinedTeamCloneRequestVisibilityPublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateJoinedTeamCloneRequestVisibility(input)
	return &out, nil
}

type ItemBodyContentType string

const (
	ItemBodyContentTypeHtml ItemBodyContentType = "html"
	ItemBodyContentTypeText ItemBodyContentType = "text"
)

func PossibleValuesForItemBodyContentType() []string {
	return []string{
		string(ItemBodyContentTypeHtml),
		string(ItemBodyContentTypeText),
	}
}

func (s *ItemBodyContentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseItemBodyContentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseItemBodyContentType(input string) (*ItemBodyContentType, error) {
	vals := map[string]ItemBodyContentType{
		"html": ItemBodyContentTypeHtml,
		"text": ItemBodyContentTypeText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ItemBodyContentType(input)
	return &out, nil
}

type TeamworkActivityTopicSource string

const (
	TeamworkActivityTopicSourceEntityUrl TeamworkActivityTopicSource = "entityUrl"
	TeamworkActivityTopicSourceText      TeamworkActivityTopicSource = "text"
)

func PossibleValuesForTeamworkActivityTopicSource() []string {
	return []string{
		string(TeamworkActivityTopicSourceEntityUrl),
		string(TeamworkActivityTopicSourceText),
	}
}

func (s *TeamworkActivityTopicSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkActivityTopicSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkActivityTopicSource(input string) (*TeamworkActivityTopicSource, error) {
	vals := map[string]TeamworkActivityTopicSource{
		"entityurl": TeamworkActivityTopicSourceEntityUrl,
		"text":      TeamworkActivityTopicSourceText,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkActivityTopicSource(input)
	return &out, nil
}
