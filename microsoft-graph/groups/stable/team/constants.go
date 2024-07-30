package team

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTeamCloneRequestPartsToClone string

const (
	CreateTeamCloneRequestPartsToCloneApps     CreateTeamCloneRequestPartsToClone = "apps"
	CreateTeamCloneRequestPartsToCloneChannels CreateTeamCloneRequestPartsToClone = "channels"
	CreateTeamCloneRequestPartsToCloneMembers  CreateTeamCloneRequestPartsToClone = "members"
	CreateTeamCloneRequestPartsToCloneSettings CreateTeamCloneRequestPartsToClone = "settings"
	CreateTeamCloneRequestPartsToCloneTabs     CreateTeamCloneRequestPartsToClone = "tabs"
)

func PossibleValuesForCreateTeamCloneRequestPartsToClone() []string {
	return []string{
		string(CreateTeamCloneRequestPartsToCloneApps),
		string(CreateTeamCloneRequestPartsToCloneChannels),
		string(CreateTeamCloneRequestPartsToCloneMembers),
		string(CreateTeamCloneRequestPartsToCloneSettings),
		string(CreateTeamCloneRequestPartsToCloneTabs),
	}
}

func (s *CreateTeamCloneRequestPartsToClone) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateTeamCloneRequestPartsToClone(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateTeamCloneRequestPartsToClone(input string) (*CreateTeamCloneRequestPartsToClone, error) {
	vals := map[string]CreateTeamCloneRequestPartsToClone{
		"apps":     CreateTeamCloneRequestPartsToCloneApps,
		"channels": CreateTeamCloneRequestPartsToCloneChannels,
		"members":  CreateTeamCloneRequestPartsToCloneMembers,
		"settings": CreateTeamCloneRequestPartsToCloneSettings,
		"tabs":     CreateTeamCloneRequestPartsToCloneTabs,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateTeamCloneRequestPartsToClone(input)
	return &out, nil
}

type CreateTeamCloneRequestVisibility string

const (
	CreateTeamCloneRequestVisibilityHiddenMembership CreateTeamCloneRequestVisibility = "hiddenMembership"
	CreateTeamCloneRequestVisibilityPrivate          CreateTeamCloneRequestVisibility = "private"
	CreateTeamCloneRequestVisibilityPublic           CreateTeamCloneRequestVisibility = "public"
)

func PossibleValuesForCreateTeamCloneRequestVisibility() []string {
	return []string{
		string(CreateTeamCloneRequestVisibilityHiddenMembership),
		string(CreateTeamCloneRequestVisibilityPrivate),
		string(CreateTeamCloneRequestVisibilityPublic),
	}
}

func (s *CreateTeamCloneRequestVisibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCreateTeamCloneRequestVisibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCreateTeamCloneRequestVisibility(input string) (*CreateTeamCloneRequestVisibility, error) {
	vals := map[string]CreateTeamCloneRequestVisibility{
		"hiddenmembership": CreateTeamCloneRequestVisibilityHiddenMembership,
		"private":          CreateTeamCloneRequestVisibilityPrivate,
		"public":           CreateTeamCloneRequestVisibilityPublic,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CreateTeamCloneRequestVisibility(input)
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
