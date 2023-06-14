package longtermretentionbackups

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LongTermRetentionDatabaseState string

const (
	LongTermRetentionDatabaseStateAll     LongTermRetentionDatabaseState = "All"
	LongTermRetentionDatabaseStateDeleted LongTermRetentionDatabaseState = "Deleted"
	LongTermRetentionDatabaseStateLive    LongTermRetentionDatabaseState = "Live"
)

func PossibleValuesForLongTermRetentionDatabaseState() []string {
	return []string{
		string(LongTermRetentionDatabaseStateAll),
		string(LongTermRetentionDatabaseStateDeleted),
		string(LongTermRetentionDatabaseStateLive),
	}
}

func (s *LongTermRetentionDatabaseState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLongTermRetentionDatabaseState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLongTermRetentionDatabaseState(input string) (*LongTermRetentionDatabaseState, error) {
	vals := map[string]LongTermRetentionDatabaseState{
		"all":     LongTermRetentionDatabaseStateAll,
		"deleted": LongTermRetentionDatabaseStateDeleted,
		"live":    LongTermRetentionDatabaseStateLive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LongTermRetentionDatabaseState(input)
	return &out, nil
}
