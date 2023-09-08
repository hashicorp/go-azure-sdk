package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTroubleshootingErrorDetails struct {
	Context        *string                                         `json:"context,omitempty"`
	Failure        *string                                         `json:"failure,omitempty"`
	FailureDetails *string                                         `json:"failureDetails,omitempty"`
	ODataType      *string                                         `json:"@odata.type,omitempty"`
	Remediation    *string                                         `json:"remediation,omitempty"`
	Resources      *[]DeviceManagementTroubleshootingErrorResource `json:"resources,omitempty"`
}
