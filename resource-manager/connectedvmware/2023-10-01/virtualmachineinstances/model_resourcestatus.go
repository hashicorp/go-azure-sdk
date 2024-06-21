package virtualmachineinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceStatus struct {
	LastUpdatedAt *string `json:"lastUpdatedAt,omitempty"`
	Message       *string `json:"message,omitempty"`
	Reason        *string `json:"reason,omitempty"`
	Severity      *string `json:"severity,omitempty"`
	Status        *string `json:"status,omitempty"`
	Type          *string `json:"type,omitempty"`
}
