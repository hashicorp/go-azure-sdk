package channels

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChannelUpdateParametersProperties struct {
	ExpirationTimeIfNotActivatedUtc *string                 `json:"expirationTimeIfNotActivatedUtc,omitempty"`
	PartnerTopicInfo                *PartnerUpdateTopicInfo `json:"partnerTopicInfo,omitempty"`
}
