package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesRoleStorageClassInfo struct {
	Name           *string                `json:"name,omitempty"`
	PosixCompliant *PosixComplianceStatus `json:"posixCompliant,omitempty"`
	Type           *string                `json:"type,omitempty"`
}
