package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamsAsyncOperationOperationType string

const (
	TeamsAsyncOperationOperationType_ArchiveTeam   TeamsAsyncOperationOperationType = "archiveTeam"
	TeamsAsyncOperationOperationType_CloneTeam     TeamsAsyncOperationOperationType = "cloneTeam"
	TeamsAsyncOperationOperationType_CreateChannel TeamsAsyncOperationOperationType = "createChannel"
	TeamsAsyncOperationOperationType_CreateChat    TeamsAsyncOperationOperationType = "createChat"
	TeamsAsyncOperationOperationType_CreateTeam    TeamsAsyncOperationOperationType = "createTeam"
	TeamsAsyncOperationOperationType_Invalid       TeamsAsyncOperationOperationType = "invalid"
	TeamsAsyncOperationOperationType_TeamifyGroup  TeamsAsyncOperationOperationType = "teamifyGroup"
	TeamsAsyncOperationOperationType_UnarchiveTeam TeamsAsyncOperationOperationType = "unarchiveTeam"
)

func PossibleValuesForTeamsAsyncOperationOperationType() []string {
	return []string{
		string(TeamsAsyncOperationOperationType_ArchiveTeam),
		string(TeamsAsyncOperationOperationType_CloneTeam),
		string(TeamsAsyncOperationOperationType_CreateChannel),
		string(TeamsAsyncOperationOperationType_CreateChat),
		string(TeamsAsyncOperationOperationType_CreateTeam),
		string(TeamsAsyncOperationOperationType_Invalid),
		string(TeamsAsyncOperationOperationType_TeamifyGroup),
		string(TeamsAsyncOperationOperationType_UnarchiveTeam),
	}
}

func (s *TeamsAsyncOperationOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamsAsyncOperationOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamsAsyncOperationOperationType(input string) (*TeamsAsyncOperationOperationType, error) {
	vals := map[string]TeamsAsyncOperationOperationType{
		"archiveteam":   TeamsAsyncOperationOperationType_ArchiveTeam,
		"cloneteam":     TeamsAsyncOperationOperationType_CloneTeam,
		"createchannel": TeamsAsyncOperationOperationType_CreateChannel,
		"createchat":    TeamsAsyncOperationOperationType_CreateChat,
		"createteam":    TeamsAsyncOperationOperationType_CreateTeam,
		"invalid":       TeamsAsyncOperationOperationType_Invalid,
		"teamifygroup":  TeamsAsyncOperationOperationType_TeamifyGroup,
		"unarchiveteam": TeamsAsyncOperationOperationType_UnarchiveTeam,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamsAsyncOperationOperationType(input)
	return &out, nil
}
