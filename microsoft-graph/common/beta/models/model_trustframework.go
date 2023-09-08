package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrustFramework struct {
	KeySets   *[]TrustFrameworkKeySet `json:"keySets,omitempty"`
	ODataType *string                 `json:"@odata.type,omitempty"`
	Policies  *[]TrustFrameworkPolicy `json:"policies,omitempty"`
}
