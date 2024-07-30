package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegeManagementElevationRequest struct {
	ApplicationDetail                  *ApplicationDetail                         `json:"applicationDetail,omitempty"`
	DeviceName                         *string                                    `json:"deviceName,omitempty"`
	Id                                 *string                                    `json:"id,omitempty"`
	ODataType                          *string                                    `json:"@odata.type,omitempty"`
	RequestCreatedDateTime             *string                                    `json:"requestCreatedDateTime,omitempty"`
	RequestExpiryDateTime              *string                                    `json:"requestExpiryDateTime,omitempty"`
	RequestJustification               *string                                    `json:"requestJustification,omitempty"`
	RequestLastModifiedDateTime        *string                                    `json:"requestLastModifiedDateTime,omitempty"`
	RequestedByUserId                  *string                                    `json:"requestedByUserId,omitempty"`
	RequestedByUserPrincipalName       *string                                    `json:"requestedByUserPrincipalName,omitempty"`
	RequestedOnDeviceId                *string                                    `json:"requestedOnDeviceId,omitempty"`
	ReviewCompletedByUserId            *string                                    `json:"reviewCompletedByUserId,omitempty"`
	ReviewCompletedByUserPrincipalName *string                                    `json:"reviewCompletedByUserPrincipalName,omitempty"`
	ReviewCompletedDateTime            *string                                    `json:"reviewCompletedDateTime,omitempty"`
	ReviewerJustification              *string                                    `json:"reviewerJustification,omitempty"`
	Status                             *PrivilegeManagementElevationRequestStatus `json:"status,omitempty"`
}
