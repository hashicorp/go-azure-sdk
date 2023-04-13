package operations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Display struct {
	Description *string `json:"description,omitempty"`
	Operation   *string `json:"operation,omitempty"`
	Provider    *string `json:"provider,omitempty"`
	Resource    *string `json:"resource,omitempty"`
}
