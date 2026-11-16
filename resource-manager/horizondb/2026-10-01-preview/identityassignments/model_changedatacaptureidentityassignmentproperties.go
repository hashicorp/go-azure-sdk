package identityassignments

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChangeDataCaptureIdentityAssignmentProperties struct {
	ActivationState    ChangeDataCaptureIdentityActivationState `json:"activationState"`
	EffectiveState     ChangeDataCaptureIdentityAssignmentState `json:"effectiveState"`
	IdentityResourceId string                                   `json:"identityResourceId"`
	PrincipalId        string                                   `json:"principalId"`
	ProvisioningState  *ProvisioningState                       `json:"provisioningState,omitempty"`
	Status             ChangeDataCaptureStatus                  `json:"status"`
	TenantId           string                                   `json:"tenantId"`
	Validation         IdentityValidationStatus                 `json:"validation"`
}
