package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidLobAppTargetedPlatforms string

const (
	AndroidLobAppTargetedPlatforms_AndroidDeviceAdministrator AndroidLobAppTargetedPlatforms = "androidDeviceAdministrator"
	AndroidLobAppTargetedPlatforms_AndroidOpenSourceProject   AndroidLobAppTargetedPlatforms = "androidOpenSourceProject"
)

func PossibleValuesForAndroidLobAppTargetedPlatforms() []string {
	return []string{
		string(AndroidLobAppTargetedPlatforms_AndroidDeviceAdministrator),
		string(AndroidLobAppTargetedPlatforms_AndroidOpenSourceProject),
	}
}

func (s *AndroidLobAppTargetedPlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAndroidLobAppTargetedPlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAndroidLobAppTargetedPlatforms(input string) (*AndroidLobAppTargetedPlatforms, error) {
	vals := map[string]AndroidLobAppTargetedPlatforms{
		"androiddeviceadministrator": AndroidLobAppTargetedPlatforms_AndroidDeviceAdministrator,
		"androidopensourceproject":   AndroidLobAppTargetedPlatforms_AndroidOpenSourceProject,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidLobAppTargetedPlatforms(input)
	return &out, nil
}
