package lots

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LotSummary struct {
	ETag       *string        `json:"eTag,omitempty"`
	Id         *string        `json:"id,omitempty"`
	Name       *string        `json:"name,omitempty"`
	Properties *LotProperties `json:"properties,omitempty"`
	Type       *string        `json:"type,omitempty"`
}
