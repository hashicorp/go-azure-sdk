package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidLobAppTargetedPlatforms string

const (
	AndroidLobAppTargetedPlatformsandroidDeviceAdministrator AndroidLobAppTargetedPlatforms = "AndroidDeviceAdministrator"
	AndroidLobAppTargetedPlatformsandroidOpenSourceProject   AndroidLobAppTargetedPlatforms = "AndroidOpenSourceProject"
)

func PossibleValuesForAndroidLobAppTargetedPlatforms() []string {
	return []string{
		string(AndroidLobAppTargetedPlatformsandroidDeviceAdministrator),
		string(AndroidLobAppTargetedPlatformsandroidOpenSourceProject),
	}
}

func parseAndroidLobAppTargetedPlatforms(input string) (*AndroidLobAppTargetedPlatforms, error) {
	vals := map[string]AndroidLobAppTargetedPlatforms{
		"androiddeviceadministrator": AndroidLobAppTargetedPlatformsandroidDeviceAdministrator,
		"androidopensourceproject":   AndroidLobAppTargetedPlatformsandroidOpenSourceProject,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidLobAppTargetedPlatforms(input)
	return &out, nil
}
