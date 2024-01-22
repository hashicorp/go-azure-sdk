package customers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Reseller struct {
	Description *string `json:"description,omitempty"`
	ResellerId  *string `json:"resellerId,omitempty"`
}
