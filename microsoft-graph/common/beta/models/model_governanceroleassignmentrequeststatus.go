package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GovernanceRoleAssignmentRequestStatus struct {
	ODataType     *string     `json:"@odata.type,omitempty"`
	Status        *string     `json:"status,omitempty"`
	StatusDetails *[]KeyValue `json:"statusDetails,omitempty"`
	SubStatus     *string     `json:"subStatus,omitempty"`
}
