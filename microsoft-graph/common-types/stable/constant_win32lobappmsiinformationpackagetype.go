package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Win32LobAppMsiInformationPackageType string

const (
	Win32LobAppMsiInformationPackageType_DualPurpose Win32LobAppMsiInformationPackageType = "dualPurpose"
	Win32LobAppMsiInformationPackageType_PerMachine  Win32LobAppMsiInformationPackageType = "perMachine"
	Win32LobAppMsiInformationPackageType_PerUser     Win32LobAppMsiInformationPackageType = "perUser"
)

func PossibleValuesForWin32LobAppMsiInformationPackageType() []string {
	return []string{
		string(Win32LobAppMsiInformationPackageType_DualPurpose),
		string(Win32LobAppMsiInformationPackageType_PerMachine),
		string(Win32LobAppMsiInformationPackageType_PerUser),
	}
}

func (s *Win32LobAppMsiInformationPackageType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWin32LobAppMsiInformationPackageType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWin32LobAppMsiInformationPackageType(input string) (*Win32LobAppMsiInformationPackageType, error) {
	vals := map[string]Win32LobAppMsiInformationPackageType{
		"dualpurpose": Win32LobAppMsiInformationPackageType_DualPurpose,
		"permachine":  Win32LobAppMsiInformationPackageType_PerMachine,
		"peruser":     Win32LobAppMsiInformationPackageType_PerUser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Win32LobAppMsiInformationPackageType(input)
	return &out, nil
}
