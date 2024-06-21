package managedinstancekeys

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceKeyProperties struct {
	AutoRotationEnabled *bool         `json:"autoRotationEnabled,omitempty"`
	CreationDate        *string       `json:"creationDate,omitempty"`
	ServerKeyType       ServerKeyType `json:"serverKeyType"`
	Thumbprint          *string       `json:"thumbprint,omitempty"`
	Uri                 *string       `json:"uri,omitempty"`
}
