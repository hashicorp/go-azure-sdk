package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustFrameworkKey struct {
	D         *string   `json:"d,omitempty"`
	Dp        *string   `json:"dp,omitempty"`
	Dq        *string   `json:"dq,omitempty"`
	E         *string   `json:"e,omitempty"`
	Exp       *int64    `json:"exp,omitempty"`
	K         *string   `json:"k,omitempty"`
	Kid       *string   `json:"kid,omitempty"`
	Kty       *string   `json:"kty,omitempty"`
	N         *string   `json:"n,omitempty"`
	Nbf       *int64    `json:"nbf,omitempty"`
	ODataType *string   `json:"@odata.type,omitempty"`
	P         *string   `json:"p,omitempty"`
	Q         *string   `json:"q,omitempty"`
	Qi        *string   `json:"qi,omitempty"`
	Use       *string   `json:"use,omitempty"`
	X5c       *[]string `json:"x5c,omitempty"`
	X5t       *string   `json:"x5t,omitempty"`
}
