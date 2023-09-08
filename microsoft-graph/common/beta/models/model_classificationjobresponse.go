package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClassificationJobResponse struct {
	CreationDateTime *string                          `json:"creationDateTime,omitempty"`
	EndDateTime      *string                          `json:"endDateTime,omitempty"`
	Error            *ClassificationError             `json:"error,omitempty"`
	Id               *string                          `json:"id,omitempty"`
	ODataType        *string                          `json:"@odata.type,omitempty"`
	Result           *DetectedSensitiveContentWrapper `json:"result,omitempty"`
	StartDateTime    *string                          `json:"startDateTime,omitempty"`
	Status           *string                          `json:"status,omitempty"`
	TenantId         *string                          `json:"tenantId,omitempty"`
	Type             *string                          `json:"type,omitempty"`
	UserId           *string                          `json:"userId,omitempty"`
}
