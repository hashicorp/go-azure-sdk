package manualtrigger

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityManualTriggerRequestBody struct {
	IncidentArmId       *string `json:"incidentArmId,omitempty"`
	LogicAppsResourceId string  `json:"logicAppsResourceId"`
	TenantId            *string `json:"tenantId,omitempty"`
}
