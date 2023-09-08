package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidLobAppTargetedPlatforms string

const (
	ManagedAndroidLobAppTargetedPlatformsandroidDeviceAdministrator ManagedAndroidLobAppTargetedPlatforms = "AndroidDeviceAdministrator"
	ManagedAndroidLobAppTargetedPlatformsandroidOpenSourceProject   ManagedAndroidLobAppTargetedPlatforms = "AndroidOpenSourceProject"
)

func PossibleValuesForManagedAndroidLobAppTargetedPlatforms() []string {
	return []string{
		string(ManagedAndroidLobAppTargetedPlatformsandroidDeviceAdministrator),
		string(ManagedAndroidLobAppTargetedPlatformsandroidOpenSourceProject),
	}
}

func parseManagedAndroidLobAppTargetedPlatforms(input string) (*ManagedAndroidLobAppTargetedPlatforms, error) {
	vals := map[string]ManagedAndroidLobAppTargetedPlatforms{
		"androiddeviceadministrator": ManagedAndroidLobAppTargetedPlatformsandroidDeviceAdministrator,
		"androidopensourceproject":   ManagedAndroidLobAppTargetedPlatformsandroidOpenSourceProject,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAndroidLobAppTargetedPlatforms(input)
	return &out, nil
}
