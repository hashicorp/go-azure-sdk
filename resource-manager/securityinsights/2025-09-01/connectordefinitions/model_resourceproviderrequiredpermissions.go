package connectordefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourceProviderRequiredPermissions struct {
	Action *bool `json:"action,omitempty"`
	Delete *bool `json:"delete,omitempty"`
	Read   *bool `json:"read,omitempty"`
	Write  *bool `json:"write,omitempty"`
}
