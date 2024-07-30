package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharepointSettingsSharingDomainRestrictionMode string

const (
	SharepointSettingsSharingDomainRestrictionMode_AllowList SharepointSettingsSharingDomainRestrictionMode = "allowList"
	SharepointSettingsSharingDomainRestrictionMode_BlockList SharepointSettingsSharingDomainRestrictionMode = "blockList"
	SharepointSettingsSharingDomainRestrictionMode_None      SharepointSettingsSharingDomainRestrictionMode = "none"
)

func PossibleValuesForSharepointSettingsSharingDomainRestrictionMode() []string {
	return []string{
		string(SharepointSettingsSharingDomainRestrictionMode_AllowList),
		string(SharepointSettingsSharingDomainRestrictionMode_BlockList),
		string(SharepointSettingsSharingDomainRestrictionMode_None),
	}
}

func (s *SharepointSettingsSharingDomainRestrictionMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharepointSettingsSharingDomainRestrictionMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharepointSettingsSharingDomainRestrictionMode(input string) (*SharepointSettingsSharingDomainRestrictionMode, error) {
	vals := map[string]SharepointSettingsSharingDomainRestrictionMode{
		"allowlist": SharepointSettingsSharingDomainRestrictionMode_AllowList,
		"blocklist": SharepointSettingsSharingDomainRestrictionMode_BlockList,
		"none":      SharepointSettingsSharingDomainRestrictionMode_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharepointSettingsSharingDomainRestrictionMode(input)
	return &out, nil
}
