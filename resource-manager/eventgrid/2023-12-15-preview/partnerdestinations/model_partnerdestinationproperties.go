package partnerdestinations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerDestinationProperties struct {
	ActivationState                 *PartnerDestinationActivationState   `json:"activationState,omitempty"`
	EndpointBaseUrl                 *string                              `json:"endpointBaseUrl,omitempty"`
	EndpointServiceContext          *string                              `json:"endpointServiceContext,omitempty"`
	ExpirationTimeIfNotActivatedUtc *string                              `json:"expirationTimeIfNotActivatedUtc,omitempty"`
	MessageForActivation            *string                              `json:"messageForActivation,omitempty"`
	PartnerRegistrationImmutableId  *string                              `json:"partnerRegistrationImmutableId,omitempty"`
	ProvisioningState               *PartnerDestinationProvisioningState `json:"provisioningState,omitempty"`
}
