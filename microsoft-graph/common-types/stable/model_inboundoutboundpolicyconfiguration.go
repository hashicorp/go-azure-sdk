package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InboundOutboundPolicyConfiguration struct {
	InboundAllowed  *bool   `json:"inboundAllowed,omitempty"`
	ODataType       *string `json:"@odata.type,omitempty"`
	OutboundAllowed *bool   `json:"outboundAllowed,omitempty"`
}
