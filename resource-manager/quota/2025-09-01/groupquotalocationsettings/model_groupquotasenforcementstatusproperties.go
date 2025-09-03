package groupquotalocationsettings

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupQuotasEnforcementStatusProperties struct {
	EnforcedGroupName  *string           `json:"enforcedGroupName,omitempty"`
	EnforcementEnabled *EnforcementState `json:"enforcementEnabled,omitempty"`
	FaultCode          *string           `json:"faultCode,omitempty"`
	ProvisioningState  *RequestState     `json:"provisioningState,omitempty"`
}
