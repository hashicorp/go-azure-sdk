package partnerconfigurations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Partner struct {
	AuthorizationExpirationTimeInUtc *string `json:"authorizationExpirationTimeInUtc,omitempty"`
	PartnerName                      *string `json:"partnerName,omitempty"`
	PartnerRegistrationImmutableId   *string `json:"partnerRegistrationImmutableId,omitempty"`
}
