package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsCreepIndexDistribution struct {
	AuthorizationSystem *AuthorizationSystem `json:"authorizationSystem,omitempty"`
	CreatedDateTime     *string              `json:"createdDateTime,omitempty"`
	HighRiskProfile     *RiskProfile         `json:"highRiskProfile,omitempty"`
	Id                  *string              `json:"id,omitempty"`
	LowRiskProfile      *RiskProfile         `json:"lowRiskProfile,omitempty"`
	MediumRiskProfile   *RiskProfile         `json:"mediumRiskProfile,omitempty"`
	ODataType           *string              `json:"@odata.type,omitempty"`
}
