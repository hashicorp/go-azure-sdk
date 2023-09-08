package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsPropertyRuleValuesJoinedBy string

const (
	ExternalConnectorsPropertyRuleValuesJoinedByand ExternalConnectorsPropertyRuleValuesJoinedBy = "And"
	ExternalConnectorsPropertyRuleValuesJoinedByor  ExternalConnectorsPropertyRuleValuesJoinedBy = "Or"
)

func PossibleValuesForExternalConnectorsPropertyRuleValuesJoinedBy() []string {
	return []string{
		string(ExternalConnectorsPropertyRuleValuesJoinedByand),
		string(ExternalConnectorsPropertyRuleValuesJoinedByor),
	}
}

func parseExternalConnectorsPropertyRuleValuesJoinedBy(input string) (*ExternalConnectorsPropertyRuleValuesJoinedBy, error) {
	vals := map[string]ExternalConnectorsPropertyRuleValuesJoinedBy{
		"and": ExternalConnectorsPropertyRuleValuesJoinedByand,
		"or":  ExternalConnectorsPropertyRuleValuesJoinedByor,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsPropertyRuleValuesJoinedBy(input)
	return &out, nil
}
