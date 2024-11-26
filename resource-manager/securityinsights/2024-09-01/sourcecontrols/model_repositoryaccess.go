package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RepositoryAccess struct {
	ClientId       *string              `json:"clientId,omitempty"`
	Code           *string              `json:"code,omitempty"`
	InstallationId *string              `json:"installationId,omitempty"`
	Kind           RepositoryAccessKind `json:"kind"`
	State          *string              `json:"state,omitempty"`
	Token          *string              `json:"token,omitempty"`
}
