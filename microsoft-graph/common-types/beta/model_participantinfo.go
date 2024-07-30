package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ParticipantInfo struct {
	CountryCode           *string                      `json:"countryCode,omitempty"`
	EndpointType          *ParticipantInfoEndpointType `json:"endpointType,omitempty"`
	Identity              *IdentitySet                 `json:"identity,omitempty"`
	LanguageId            *string                      `json:"languageId,omitempty"`
	NonAnonymizedIdentity *IdentitySet                 `json:"nonAnonymizedIdentity,omitempty"`
	ODataType             *string                      `json:"@odata.type,omitempty"`
	ParticipantId         *string                      `json:"participantId,omitempty"`
	PlatformId            *string                      `json:"platformId,omitempty"`
	Region                *string                      `json:"region,omitempty"`
}
