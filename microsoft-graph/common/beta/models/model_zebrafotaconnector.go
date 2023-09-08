package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ZebraFotaConnector struct {
	EnrollmentAuthorizationUrl *string                  `json:"enrollmentAuthorizationUrl,omitempty"`
	EnrollmentToken            *string                  `json:"enrollmentToken,omitempty"`
	FotaAppsApproved           *bool                    `json:"fotaAppsApproved,omitempty"`
	Id                         *string                  `json:"id,omitempty"`
	LastSyncDateTime           *string                  `json:"lastSyncDateTime,omitempty"`
	ODataType                  *string                  `json:"@odata.type,omitempty"`
	State                      *ZebraFotaConnectorState `json:"state,omitempty"`
}
