package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentStatusState string

const (
	ZebraFotaDeploymentStatusState_Canceled        ZebraFotaDeploymentStatusState = "canceled"
	ZebraFotaDeploymentStatusState_Completed       ZebraFotaDeploymentStatusState = "completed"
	ZebraFotaDeploymentStatusState_CreateFailed    ZebraFotaDeploymentStatusState = "createFailed"
	ZebraFotaDeploymentStatusState_Created         ZebraFotaDeploymentStatusState = "created"
	ZebraFotaDeploymentStatusState_InProgress      ZebraFotaDeploymentStatusState = "inProgress"
	ZebraFotaDeploymentStatusState_PendingCancel   ZebraFotaDeploymentStatusState = "pendingCancel"
	ZebraFotaDeploymentStatusState_PendingCreation ZebraFotaDeploymentStatusState = "pendingCreation"
)

func PossibleValuesForZebraFotaDeploymentStatusState() []string {
	return []string{
		string(ZebraFotaDeploymentStatusState_Canceled),
		string(ZebraFotaDeploymentStatusState_Completed),
		string(ZebraFotaDeploymentStatusState_CreateFailed),
		string(ZebraFotaDeploymentStatusState_Created),
		string(ZebraFotaDeploymentStatusState_InProgress),
		string(ZebraFotaDeploymentStatusState_PendingCancel),
		string(ZebraFotaDeploymentStatusState_PendingCreation),
	}
}

func (s *ZebraFotaDeploymentStatusState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseZebraFotaDeploymentStatusState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseZebraFotaDeploymentStatusState(input string) (*ZebraFotaDeploymentStatusState, error) {
	vals := map[string]ZebraFotaDeploymentStatusState{
		"canceled":        ZebraFotaDeploymentStatusState_Canceled,
		"completed":       ZebraFotaDeploymentStatusState_Completed,
		"createfailed":    ZebraFotaDeploymentStatusState_CreateFailed,
		"created":         ZebraFotaDeploymentStatusState_Created,
		"inprogress":      ZebraFotaDeploymentStatusState_InProgress,
		"pendingcancel":   ZebraFotaDeploymentStatusState_PendingCancel,
		"pendingcreation": ZebraFotaDeploymentStatusState_PendingCreation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaDeploymentStatusState(input)
	return &out, nil
}
