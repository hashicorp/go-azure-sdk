package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommandOutputSettings struct {
	AssociatedIdentity *IdentitySelector        `json:"associatedIdentity,omitempty"`
	ContainerURL       *string                  `json:"containerUrl,omitempty"`
	Overrides          *[]CommandOutputOverride `json:"overrides,omitempty"`
}
