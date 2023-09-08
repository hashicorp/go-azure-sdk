package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaDeploymentStatusState string

const (
	ZebraFotaDeploymentStatusStatecanceled        ZebraFotaDeploymentStatusState = "Canceled"
	ZebraFotaDeploymentStatusStatecompleted       ZebraFotaDeploymentStatusState = "Completed"
	ZebraFotaDeploymentStatusStatecreateFailed    ZebraFotaDeploymentStatusState = "CreateFailed"
	ZebraFotaDeploymentStatusStatecreated         ZebraFotaDeploymentStatusState = "Created"
	ZebraFotaDeploymentStatusStateinProgress      ZebraFotaDeploymentStatusState = "InProgress"
	ZebraFotaDeploymentStatusStatependingCancel   ZebraFotaDeploymentStatusState = "PendingCancel"
	ZebraFotaDeploymentStatusStatependingCreation ZebraFotaDeploymentStatusState = "PendingCreation"
)

func PossibleValuesForZebraFotaDeploymentStatusState() []string {
	return []string{
		string(ZebraFotaDeploymentStatusStatecanceled),
		string(ZebraFotaDeploymentStatusStatecompleted),
		string(ZebraFotaDeploymentStatusStatecreateFailed),
		string(ZebraFotaDeploymentStatusStatecreated),
		string(ZebraFotaDeploymentStatusStateinProgress),
		string(ZebraFotaDeploymentStatusStatependingCancel),
		string(ZebraFotaDeploymentStatusStatependingCreation),
	}
}

func parseZebraFotaDeploymentStatusState(input string) (*ZebraFotaDeploymentStatusState, error) {
	vals := map[string]ZebraFotaDeploymentStatusState{
		"canceled":        ZebraFotaDeploymentStatusStatecanceled,
		"completed":       ZebraFotaDeploymentStatusStatecompleted,
		"createfailed":    ZebraFotaDeploymentStatusStatecreateFailed,
		"created":         ZebraFotaDeploymentStatusStatecreated,
		"inprogress":      ZebraFotaDeploymentStatusStateinProgress,
		"pendingcancel":   ZebraFotaDeploymentStatusStatependingCancel,
		"pendingcreation": ZebraFotaDeploymentStatusStatependingCreation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ZebraFotaDeploymentStatusState(input)
	return &out, nil
}
