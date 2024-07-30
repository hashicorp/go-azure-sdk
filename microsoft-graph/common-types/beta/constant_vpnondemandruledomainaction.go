package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnOnDemandRuleDomainAction string

const (
	VpnOnDemandRuleDomainAction_ConnectIfNeeded VpnOnDemandRuleDomainAction = "connectIfNeeded"
	VpnOnDemandRuleDomainAction_NeverConnect    VpnOnDemandRuleDomainAction = "neverConnect"
)

func PossibleValuesForVpnOnDemandRuleDomainAction() []string {
	return []string{
		string(VpnOnDemandRuleDomainAction_ConnectIfNeeded),
		string(VpnOnDemandRuleDomainAction_NeverConnect),
	}
}

func (s *VpnOnDemandRuleDomainAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVpnOnDemandRuleDomainAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVpnOnDemandRuleDomainAction(input string) (*VpnOnDemandRuleDomainAction, error) {
	vals := map[string]VpnOnDemandRuleDomainAction{
		"connectifneeded": VpnOnDemandRuleDomainAction_ConnectIfNeeded,
		"neverconnect":    VpnOnDemandRuleDomainAction_NeverConnect,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnOnDemandRuleDomainAction(input)
	return &out, nil
}
