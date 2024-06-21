package partnertopics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartnerTopicProperties struct {
	ActivationState                 *PartnerTopicActivationState   `json:"activationState,omitempty"`
	EventTypeInfo                   *EventTypeInfo                 `json:"eventTypeInfo,omitempty"`
	ExpirationTimeIfNotActivatedUtc *string                        `json:"expirationTimeIfNotActivatedUtc,omitempty"`
	MessageForActivation            *string                        `json:"messageForActivation,omitempty"`
	PartnerRegistrationImmutableId  *string                        `json:"partnerRegistrationImmutableId,omitempty"`
	PartnerTopicFriendlyDescription *string                        `json:"partnerTopicFriendlyDescription,omitempty"`
	ProvisioningState               *PartnerTopicProvisioningState `json:"provisioningState,omitempty"`
	Source                          *string                        `json:"source,omitempty"`
}
