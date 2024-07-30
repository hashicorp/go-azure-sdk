package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentSettingsScheduleMode string

const (
	ZebraFotaDeploymentSettingsScheduleMode_InstallNow ZebraFotaDeploymentSettingsScheduleMode = "installNow"
	ZebraFotaDeploymentSettingsScheduleMode_Scheduled  ZebraFotaDeploymentSettingsScheduleMode = "scheduled"
)

func PossibleValuesForZebraFotaDeploymentSettingsScheduleMode() []string {
	return []string{
		string(ZebraFotaDeploymentSettingsScheduleMode_InstallNow),
		string(ZebraFotaDeploymentSettingsScheduleMode_Scheduled),
	}
}

func (s *ZebraFotaDeploymentSettingsScheduleMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaDeploymentSettingsScheduleMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaDeploymentSettingsScheduleMode(input string) (*ZebraFotaDeploymentSettingsScheduleMode, error) {
	vals := map[string]ZebraFotaDeploymentSettingsScheduleMode{
		"installnow": ZebraFotaDeploymentSettingsScheduleMode_InstallNow,
		"scheduled":  ZebraFotaDeploymentSettingsScheduleMode_Scheduled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaDeploymentSettingsScheduleMode(input)
	return &out, nil
}
