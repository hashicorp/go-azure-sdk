package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInCounts struct {
	NoSignIn   *int64  `json:"noSignIn,omitempty"`
	ODataType  *string `json:"@odata.type,omitempty"`
	WithSignIn *int64  `json:"withSignIn,omitempty"`
}
