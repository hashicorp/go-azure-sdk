package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCOnPremisesConnectionStatusDetails struct {
	EndDateTime   *string                                   `json:"endDateTime,omitempty"`
	HealthChecks  *[]CloudPCOnPremisesConnectionHealthCheck `json:"healthChecks,omitempty"`
	ODataType     *string                                   `json:"@odata.type,omitempty"`
	StartDateTime *string                                   `json:"startDateTime,omitempty"`
}
