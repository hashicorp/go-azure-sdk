package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharepointSettingsSharingCapability string

const (
	SharepointSettingsSharingCapabilitydisabled                        SharepointSettingsSharingCapability = "Disabled"
	SharepointSettingsSharingCapabilityexistingExternalUserSharingOnly SharepointSettingsSharingCapability = "ExistingExternalUserSharingOnly"
	SharepointSettingsSharingCapabilityexternalUserAndGuestSharing     SharepointSettingsSharingCapability = "ExternalUserAndGuestSharing"
	SharepointSettingsSharingCapabilityexternalUserSharingOnly         SharepointSettingsSharingCapability = "ExternalUserSharingOnly"
)

func PossibleValuesForSharepointSettingsSharingCapability() []string {
	return []string{
		string(SharepointSettingsSharingCapabilitydisabled),
		string(SharepointSettingsSharingCapabilityexistingExternalUserSharingOnly),
		string(SharepointSettingsSharingCapabilityexternalUserAndGuestSharing),
		string(SharepointSettingsSharingCapabilityexternalUserSharingOnly),
	}
}

func parseSharepointSettingsSharingCapability(input string) (*SharepointSettingsSharingCapability, error) {
	vals := map[string]SharepointSettingsSharingCapability{
		"disabled":                        SharepointSettingsSharingCapabilitydisabled,
		"existingexternalusersharingonly": SharepointSettingsSharingCapabilityexistingExternalUserSharingOnly,
		"externaluserandguestsharing":     SharepointSettingsSharingCapabilityexternalUserAndGuestSharing,
		"externalusersharingonly":         SharepointSettingsSharingCapabilityexternalUserSharingOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SharepointSettingsSharingCapability(input)
	return &out, nil
}
