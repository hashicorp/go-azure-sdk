package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkApplicationIdentityApplicationIdentityType string

const (
	TeamworkApplicationIdentityApplicationIdentityType_AadApplication     TeamworkApplicationIdentityApplicationIdentityType = "aadApplication"
	TeamworkApplicationIdentityApplicationIdentityType_Bot                TeamworkApplicationIdentityApplicationIdentityType = "bot"
	TeamworkApplicationIdentityApplicationIdentityType_Office365Connector TeamworkApplicationIdentityApplicationIdentityType = "office365Connector"
	TeamworkApplicationIdentityApplicationIdentityType_OutgoingWebhook    TeamworkApplicationIdentityApplicationIdentityType = "outgoingWebhook"
	TeamworkApplicationIdentityApplicationIdentityType_TenantBot          TeamworkApplicationIdentityApplicationIdentityType = "tenantBot"
)

func PossibleValuesForTeamworkApplicationIdentityApplicationIdentityType() []string {
	return []string{
		string(TeamworkApplicationIdentityApplicationIdentityType_AadApplication),
		string(TeamworkApplicationIdentityApplicationIdentityType_Bot),
		string(TeamworkApplicationIdentityApplicationIdentityType_Office365Connector),
		string(TeamworkApplicationIdentityApplicationIdentityType_OutgoingWebhook),
		string(TeamworkApplicationIdentityApplicationIdentityType_TenantBot),
	}
}

func (s *TeamworkApplicationIdentityApplicationIdentityType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseTeamworkApplicationIdentityApplicationIdentityType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseTeamworkApplicationIdentityApplicationIdentityType(input string) (*TeamworkApplicationIdentityApplicationIdentityType, error) {
	vals := map[string]TeamworkApplicationIdentityApplicationIdentityType{
		"aadapplication":     TeamworkApplicationIdentityApplicationIdentityType_AadApplication,
		"bot":                TeamworkApplicationIdentityApplicationIdentityType_Bot,
		"office365connector": TeamworkApplicationIdentityApplicationIdentityType_Office365Connector,
		"outgoingwebhook":    TeamworkApplicationIdentityApplicationIdentityType_OutgoingWebhook,
		"tenantbot":          TeamworkApplicationIdentityApplicationIdentityType_TenantBot,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TeamworkApplicationIdentityApplicationIdentityType(input)
	return &out, nil
}
