package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommandOutputOverride struct {
	AssociatedIdentity *IdentitySelector  `json:"associatedIdentity,omitempty"`
	CommandOutputType  *CommandOutputType `json:"commandOutputType,omitempty"`
	ContainerURL       *string            `json:"containerUrl,omitempty"`
}
