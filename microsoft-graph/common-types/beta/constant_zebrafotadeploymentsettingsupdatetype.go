package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentSettingsUpdateType string

const (
	ZebraFotaDeploymentSettingsUpdateType_Auto   ZebraFotaDeploymentSettingsUpdateType = "auto"
	ZebraFotaDeploymentSettingsUpdateType_Custom ZebraFotaDeploymentSettingsUpdateType = "custom"
	ZebraFotaDeploymentSettingsUpdateType_Latest ZebraFotaDeploymentSettingsUpdateType = "latest"
)

func PossibleValuesForZebraFotaDeploymentSettingsUpdateType() []string {
	return []string{
		string(ZebraFotaDeploymentSettingsUpdateType_Auto),
		string(ZebraFotaDeploymentSettingsUpdateType_Custom),
		string(ZebraFotaDeploymentSettingsUpdateType_Latest),
	}
}

func (s *ZebraFotaDeploymentSettingsUpdateType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaDeploymentSettingsUpdateType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaDeploymentSettingsUpdateType(input string) (*ZebraFotaDeploymentSettingsUpdateType, error) {
	vals := map[string]ZebraFotaDeploymentSettingsUpdateType{
		"auto":   ZebraFotaDeploymentSettingsUpdateType_Auto,
		"custom": ZebraFotaDeploymentSettingsUpdateType_Custom,
		"latest": ZebraFotaDeploymentSettingsUpdateType_Latest,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaDeploymentSettingsUpdateType(input)
	return &out, nil
}
