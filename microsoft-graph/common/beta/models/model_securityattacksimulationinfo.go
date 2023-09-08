package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAttackSimulationInfo struct {
	AttackSimDateTime     *string `json:"attackSimDateTime,omitempty"`
	AttackSimDurationTime *string `json:"attackSimDurationTime,omitempty"`
	AttackSimId           *string `json:"attackSimId,omitempty"`
	AttackSimUserId       *string `json:"attackSimUserId,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
}
