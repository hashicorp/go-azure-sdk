package user

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserContractProperties struct {
	Email            *string                    `json:"email,omitempty"`
	FirstName        *string                    `json:"firstName,omitempty"`
	Groups           *[]GroupContractProperties `json:"groups,omitempty"`
	Identities       *[]UserIdentityContract    `json:"identities,omitempty"`
	LastName         *string                    `json:"lastName,omitempty"`
	Note             *string                    `json:"note,omitempty"`
	RegistrationDate *string                    `json:"registrationDate,omitempty"`
	State            *UserState                 `json:"state,omitempty"`
}
