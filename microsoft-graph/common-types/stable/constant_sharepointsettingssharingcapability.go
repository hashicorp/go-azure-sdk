package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharepointSettingsSharingCapability string

const (
	SharepointSettingsSharingCapability_Disabled                        SharepointSettingsSharingCapability = "disabled"
	SharepointSettingsSharingCapability_ExistingExternalUserSharingOnly SharepointSettingsSharingCapability = "existingExternalUserSharingOnly"
	SharepointSettingsSharingCapability_ExternalUserAndGuestSharing     SharepointSettingsSharingCapability = "externalUserAndGuestSharing"
	SharepointSettingsSharingCapability_ExternalUserSharingOnly         SharepointSettingsSharingCapability = "externalUserSharingOnly"
)

func PossibleValuesForSharepointSettingsSharingCapability() []string {
	return []string{
		string(SharepointSettingsSharingCapability_Disabled),
		string(SharepointSettingsSharingCapability_ExistingExternalUserSharingOnly),
		string(SharepointSettingsSharingCapability_ExternalUserAndGuestSharing),
		string(SharepointSettingsSharingCapability_ExternalUserSharingOnly),
	}
}

func (s *SharepointSettingsSharingCapability) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSharepointSettingsSharingCapability(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSharepointSettingsSharingCapability(input string) (*SharepointSettingsSharingCapability, error) {
	vals := map[string]SharepointSettingsSharingCapability{
		"disabled":                        SharepointSettingsSharingCapability_Disabled,
		"existingexternalusersharingonly": SharepointSettingsSharingCapability_ExistingExternalUserSharingOnly,
		"externaluserandguestsharing":     SharepointSettingsSharingCapability_ExternalUserAndGuestSharing,
		"externalusersharingonly":         SharepointSettingsSharingCapability_ExternalUserSharingOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharepointSettingsSharingCapability(input)
	return &out, nil
}
