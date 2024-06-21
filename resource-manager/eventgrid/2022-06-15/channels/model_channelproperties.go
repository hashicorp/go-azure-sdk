package channels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelProperties struct {
	ChannelType                     *ChannelType              `json:"channelType,omitempty"`
	ExpirationTimeIfNotActivatedUtc *string                   `json:"expirationTimeIfNotActivatedUtc,omitempty"`
	MessageForActivation            *string                   `json:"messageForActivation,omitempty"`
	PartnerTopicInfo                *PartnerTopicInfo         `json:"partnerTopicInfo,omitempty"`
	ProvisioningState               *ChannelProvisioningState `json:"provisioningState,omitempty"`
	ReadinessState                  *ReadinessState           `json:"readinessState,omitempty"`
}
