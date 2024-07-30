package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncomingContext struct {
	ODataType             *string      `json:"@odata.type,omitempty"`
	ObservedParticipantId *string      `json:"observedParticipantId,omitempty"`
	OnBehalfOf            *IdentitySet `json:"onBehalfOf,omitempty"`
	SourceParticipantId   *string      `json:"sourceParticipantId,omitempty"`
	Transferor            *IdentitySet `json:"transferor,omitempty"`
}
