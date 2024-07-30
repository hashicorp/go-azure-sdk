package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynchronizationJobRestartCriteriaResetScope string

const (
	SynchronizationJobRestartCriteriaResetScope_ConnectorDataStore SynchronizationJobRestartCriteriaResetScope = "ConnectorDataStore"
	SynchronizationJobRestartCriteriaResetScope_Escrows            SynchronizationJobRestartCriteriaResetScope = "Escrows"
	SynchronizationJobRestartCriteriaResetScope_ForceDeletes       SynchronizationJobRestartCriteriaResetScope = "ForceDeletes"
	SynchronizationJobRestartCriteriaResetScope_Full               SynchronizationJobRestartCriteriaResetScope = "Full"
	SynchronizationJobRestartCriteriaResetScope_None               SynchronizationJobRestartCriteriaResetScope = "None"
	SynchronizationJobRestartCriteriaResetScope_QuarantineState    SynchronizationJobRestartCriteriaResetScope = "QuarantineState"
	SynchronizationJobRestartCriteriaResetScope_Watermark          SynchronizationJobRestartCriteriaResetScope = "Watermark"
)

func PossibleValuesForSynchronizationJobRestartCriteriaResetScope() []string {
	return []string{
		string(SynchronizationJobRestartCriteriaResetScope_ConnectorDataStore),
		string(SynchronizationJobRestartCriteriaResetScope_Escrows),
		string(SynchronizationJobRestartCriteriaResetScope_ForceDeletes),
		string(SynchronizationJobRestartCriteriaResetScope_Full),
		string(SynchronizationJobRestartCriteriaResetScope_None),
		string(SynchronizationJobRestartCriteriaResetScope_QuarantineState),
		string(SynchronizationJobRestartCriteriaResetScope_Watermark),
	}
}

func (s *SynchronizationJobRestartCriteriaResetScope) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSynchronizationJobRestartCriteriaResetScope(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSynchronizationJobRestartCriteriaResetScope(input string) (*SynchronizationJobRestartCriteriaResetScope, error) {
	vals := map[string]SynchronizationJobRestartCriteriaResetScope{
		"connectordatastore": SynchronizationJobRestartCriteriaResetScope_ConnectorDataStore,
		"escrows":            SynchronizationJobRestartCriteriaResetScope_Escrows,
		"forcedeletes":       SynchronizationJobRestartCriteriaResetScope_ForceDeletes,
		"full":               SynchronizationJobRestartCriteriaResetScope_Full,
		"none":               SynchronizationJobRestartCriteriaResetScope_None,
		"quarantinestate":    SynchronizationJobRestartCriteriaResetScope_QuarantineState,
		"watermark":          SynchronizationJobRestartCriteriaResetScope_Watermark,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationJobRestartCriteriaResetScope(input)
	return &out, nil
}
