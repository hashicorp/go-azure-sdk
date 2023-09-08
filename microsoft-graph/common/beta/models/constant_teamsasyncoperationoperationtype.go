package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAsyncOperationOperationType string

const (
	TeamsAsyncOperationOperationTypearchiveTeam   TeamsAsyncOperationOperationType = "ArchiveTeam"
	TeamsAsyncOperationOperationTypecloneTeam     TeamsAsyncOperationOperationType = "CloneTeam"
	TeamsAsyncOperationOperationTypecreateChannel TeamsAsyncOperationOperationType = "CreateChannel"
	TeamsAsyncOperationOperationTypecreateChat    TeamsAsyncOperationOperationType = "CreateChat"
	TeamsAsyncOperationOperationTypecreateTeam    TeamsAsyncOperationOperationType = "CreateTeam"
	TeamsAsyncOperationOperationTypeinvalid       TeamsAsyncOperationOperationType = "Invalid"
	TeamsAsyncOperationOperationTypeteamifyGroup  TeamsAsyncOperationOperationType = "TeamifyGroup"
	TeamsAsyncOperationOperationTypeunarchiveTeam TeamsAsyncOperationOperationType = "UnarchiveTeam"
)

func PossibleValuesForTeamsAsyncOperationOperationType() []string {
	return []string{
		string(TeamsAsyncOperationOperationTypearchiveTeam),
		string(TeamsAsyncOperationOperationTypecloneTeam),
		string(TeamsAsyncOperationOperationTypecreateChannel),
		string(TeamsAsyncOperationOperationTypecreateChat),
		string(TeamsAsyncOperationOperationTypecreateTeam),
		string(TeamsAsyncOperationOperationTypeinvalid),
		string(TeamsAsyncOperationOperationTypeteamifyGroup),
		string(TeamsAsyncOperationOperationTypeunarchiveTeam),
	}
}

func parseTeamsAsyncOperationOperationType(input string) (*TeamsAsyncOperationOperationType, error) {
	vals := map[string]TeamsAsyncOperationOperationType{
		"archiveteam":   TeamsAsyncOperationOperationTypearchiveTeam,
		"cloneteam":     TeamsAsyncOperationOperationTypecloneTeam,
		"createchannel": TeamsAsyncOperationOperationTypecreateChannel,
		"createchat":    TeamsAsyncOperationOperationTypecreateChat,
		"createteam":    TeamsAsyncOperationOperationTypecreateTeam,
		"invalid":       TeamsAsyncOperationOperationTypeinvalid,
		"teamifygroup":  TeamsAsyncOperationOperationTypeteamifyGroup,
		"unarchiveteam": TeamsAsyncOperationOperationTypeunarchiveTeam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAsyncOperationOperationType(input)
	return &out, nil
}
