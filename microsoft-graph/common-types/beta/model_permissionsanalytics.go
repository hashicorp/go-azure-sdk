package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PermissionsAnalytics struct {
	Findings                           *[]Finding                           `json:"findings,omitempty"`
	Id                                 *string                              `json:"id,omitempty"`
	ODataType                          *string                              `json:"@odata.type,omitempty"`
	PermissionsCreepIndexDistributions *[]PermissionsCreepIndexDistribution `json:"permissionsCreepIndexDistributions,omitempty"`
}
