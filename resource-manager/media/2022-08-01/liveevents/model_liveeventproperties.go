package liveevents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LiveEventProperties struct {
	Created                 *string                   `json:"created,omitempty"`
	CrossSiteAccessPolicies *CrossSiteAccessPolicies  `json:"crossSiteAccessPolicies,omitempty"`
	Description             *string                   `json:"description,omitempty"`
	Encoding                *LiveEventEncoding        `json:"encoding,omitempty"`
	HostnamePrefix          *string                   `json:"hostnamePrefix,omitempty"`
	Input                   LiveEventInput            `json:"input"`
	LastModified            *string                   `json:"lastModified,omitempty"`
	Preview                 *LiveEventPreview         `json:"preview,omitempty"`
	ProvisioningState       *string                   `json:"provisioningState,omitempty"`
	ResourceState           *LiveEventResourceState   `json:"resourceState,omitempty"`
	StreamOptions           *[]StreamOptionsFlag      `json:"streamOptions,omitempty"`
	Transcriptions          *[]LiveEventTranscription `json:"transcriptions,omitempty"`
	UseStaticHostname       *bool                     `json:"useStaticHostname,omitempty"`
}
