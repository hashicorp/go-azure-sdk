package sqlpools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolPatchInfo struct {
	Location   *string                    `json:"location,omitempty"`
	Properties *SqlPoolResourceProperties `json:"properties,omitempty"`
	Sku        *Sku                       `json:"sku,omitempty"`
	Tags       *map[string]string         `json:"tags,omitempty"`
}
