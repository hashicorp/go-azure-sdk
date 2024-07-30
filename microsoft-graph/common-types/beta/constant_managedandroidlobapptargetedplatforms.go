package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedAndroidLobAppTargetedPlatforms string

const (
	ManagedAndroidLobAppTargetedPlatforms_AndroidDeviceAdministrator ManagedAndroidLobAppTargetedPlatforms = "androidDeviceAdministrator"
	ManagedAndroidLobAppTargetedPlatforms_AndroidOpenSourceProject   ManagedAndroidLobAppTargetedPlatforms = "androidOpenSourceProject"
)

func PossibleValuesForManagedAndroidLobAppTargetedPlatforms() []string {
	return []string{
		string(ManagedAndroidLobAppTargetedPlatforms_AndroidDeviceAdministrator),
		string(ManagedAndroidLobAppTargetedPlatforms_AndroidOpenSourceProject),
	}
}

func (s *ManagedAndroidLobAppTargetedPlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedAndroidLobAppTargetedPlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedAndroidLobAppTargetedPlatforms(input string) (*ManagedAndroidLobAppTargetedPlatforms, error) {
	vals := map[string]ManagedAndroidLobAppTargetedPlatforms{
		"androiddeviceadministrator": ManagedAndroidLobAppTargetedPlatforms_AndroidDeviceAdministrator,
		"androidopensourceproject":   ManagedAndroidLobAppTargetedPlatforms_AndroidOpenSourceProject,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedAndroidLobAppTargetedPlatforms(input)
	return &out, nil
}
