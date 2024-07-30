package lifecycleworkflowdeleteditemworkflow

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ObjectIdentity struct {
	Issuer           *string `json:"issuer,omitempty"`
	IssuerAssignedId *string `json:"issuerAssignedId,omitempty"`
	ODataType        *string `json:"@odata.type,omitempty"`
	SignInType       *string `json:"signInType,omitempty"`
}
