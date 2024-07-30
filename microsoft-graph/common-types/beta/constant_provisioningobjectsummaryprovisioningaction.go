package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningObjectSummaryProvisioningAction string

const (
	ProvisioningObjectSummaryProvisioningAction_Create       ProvisioningObjectSummaryProvisioningAction = "create"
	ProvisioningObjectSummaryProvisioningAction_Delete       ProvisioningObjectSummaryProvisioningAction = "delete"
	ProvisioningObjectSummaryProvisioningAction_Disable      ProvisioningObjectSummaryProvisioningAction = "disable"
	ProvisioningObjectSummaryProvisioningAction_Other        ProvisioningObjectSummaryProvisioningAction = "other"
	ProvisioningObjectSummaryProvisioningAction_StagedDelete ProvisioningObjectSummaryProvisioningAction = "stagedDelete"
	ProvisioningObjectSummaryProvisioningAction_Update       ProvisioningObjectSummaryProvisioningAction = "update"
)

func PossibleValuesForProvisioningObjectSummaryProvisioningAction() []string {
	return []string{
		string(ProvisioningObjectSummaryProvisioningAction_Create),
		string(ProvisioningObjectSummaryProvisioningAction_Delete),
		string(ProvisioningObjectSummaryProvisioningAction_Disable),
		string(ProvisioningObjectSummaryProvisioningAction_Other),
		string(ProvisioningObjectSummaryProvisioningAction_StagedDelete),
		string(ProvisioningObjectSummaryProvisioningAction_Update),
	}
}

func (s *ProvisioningObjectSummaryProvisioningAction) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningObjectSummaryProvisioningAction(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningObjectSummaryProvisioningAction(input string) (*ProvisioningObjectSummaryProvisioningAction, error) {
	vals := map[string]ProvisioningObjectSummaryProvisioningAction{
		"create":       ProvisioningObjectSummaryProvisioningAction_Create,
		"delete":       ProvisioningObjectSummaryProvisioningAction_Delete,
		"disable":      ProvisioningObjectSummaryProvisioningAction_Disable,
		"other":        ProvisioningObjectSummaryProvisioningAction_Other,
		"stageddelete": ProvisioningObjectSummaryProvisioningAction_StagedDelete,
		"update":       ProvisioningObjectSummaryProvisioningAction_Update,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningObjectSummaryProvisioningAction(input)
	return &out, nil
}
