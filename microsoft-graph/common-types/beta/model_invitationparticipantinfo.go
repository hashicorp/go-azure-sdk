package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationParticipantInfo struct {
	EndpointType                       *InvitationParticipantInfoEndpointType `json:"endpointType,omitempty"`
	Hidden                             *bool                                  `json:"hidden,omitempty"`
	Identity                           *IdentitySet                           `json:"identity,omitempty"`
	ODataType                          *string                                `json:"@odata.type,omitempty"`
	ParticipantId                      *string                                `json:"participantId,omitempty"`
	RemoveFromDefaultAudioRoutingGroup *bool                                  `json:"removeFromDefaultAudioRoutingGroup,omitempty"`
	ReplacesCallId                     *string                                `json:"replacesCallId,omitempty"`
}
