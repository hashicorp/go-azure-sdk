package emailregistration

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmailRegistration struct {
	ActivationCode           *string             `json:"activationCode,omitempty"`
	ActivationExpirationDate *string             `json:"activationExpirationDate,omitempty"`
	Email                    *string             `json:"email,omitempty"`
	RegistrationStatus       *RegistrationStatus `json:"registrationStatus,omitempty"`
	TenantId                 *string             `json:"tenantId,omitempty"`
}
