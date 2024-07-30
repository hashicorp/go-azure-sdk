package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSAppleEventReceiver struct {
	Allowed         *bool                                  `json:"allowed,omitempty"`
	CodeRequirement *string                                `json:"codeRequirement,omitempty"`
	Identifier      *string                                `json:"identifier,omitempty"`
	IdentifierType  *MacOSAppleEventReceiverIdentifierType `json:"identifierType,omitempty"`
	ODataType       *string                                `json:"@odata.type,omitempty"`
}
