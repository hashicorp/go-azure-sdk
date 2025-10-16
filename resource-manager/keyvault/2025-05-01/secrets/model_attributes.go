package secrets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Attributes struct {
	Created *int64 `json:"created,omitempty"`
	Enabled *bool  `json:"enabled,omitempty"`
	Exp     *int64 `json:"exp,omitempty"`
	Nbf     *int64 `json:"nbf,omitempty"`
	Updated *int64 `json:"updated,omitempty"`
}
