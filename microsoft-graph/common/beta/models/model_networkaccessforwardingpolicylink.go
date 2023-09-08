package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessForwardingPolicyLink struct {
	Id        *string                                 `json:"id,omitempty"`
	ODataType *string                                 `json:"@odata.type,omitempty"`
	Policy    *NetworkaccessPolicy                    `json:"policy,omitempty"`
	State     *NetworkaccessForwardingPolicyLinkState `json:"state,omitempty"`
	Version   *string                                 `json:"version,omitempty"`
}
