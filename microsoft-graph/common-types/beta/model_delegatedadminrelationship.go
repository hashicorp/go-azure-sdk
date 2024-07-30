package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DelegatedAdminRelationship struct {
	AccessAssignments    *[]DelegatedAdminAccessAssignment              `json:"accessAssignments,omitempty"`
	AccessDetails        *DelegatedAdminAccessDetails                   `json:"accessDetails,omitempty"`
	ActivatedDateTime    *string                                        `json:"activatedDateTime,omitempty"`
	AutoExtendDuration   *string                                        `json:"autoExtendDuration,omitempty"`
	CreatedDateTime      *string                                        `json:"createdDateTime,omitempty"`
	Customer             *DelegatedAdminRelationshipCustomerParticipant `json:"customer,omitempty"`
	DisplayName          *string                                        `json:"displayName,omitempty"`
	Duration             *string                                        `json:"duration,omitempty"`
	EndDateTime          *string                                        `json:"endDateTime,omitempty"`
	Id                   *string                                        `json:"id,omitempty"`
	LastModifiedDateTime *string                                        `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                        `json:"@odata.type,omitempty"`
	Operations           *[]DelegatedAdminRelationshipOperation         `json:"operations,omitempty"`
	Requests             *[]DelegatedAdminRelationshipRequest           `json:"requests,omitempty"`
	Status               *DelegatedAdminRelationshipStatus              `json:"status,omitempty"`
}
