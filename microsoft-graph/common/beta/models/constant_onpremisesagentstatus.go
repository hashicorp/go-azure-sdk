package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnPremisesAgentStatus string

const (
	OnPremisesAgentStatusactive   OnPremisesAgentStatus = "Active"
	OnPremisesAgentStatusinactive OnPremisesAgentStatus = "Inactive"
)

func PossibleValuesForOnPremisesAgentStatus() []string {
	return []string{
		string(OnPremisesAgentStatusactive),
		string(OnPremisesAgentStatusinactive),
	}
}

func parseOnPremisesAgentStatus(input string) (*OnPremisesAgentStatus, error) {
	vals := map[string]OnPremisesAgentStatus{
		"active":   OnPremisesAgentStatusactive,
		"inactive": OnPremisesAgentStatusinactive,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnPremisesAgentStatus(input)
	return &out, nil
}
