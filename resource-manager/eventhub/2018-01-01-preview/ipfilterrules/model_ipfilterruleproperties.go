package ipfilterrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpFilterRuleProperties struct {
	Action     *IPAction `json:"action,omitempty"`
	FilterName *string   `json:"filterName,omitempty"`
	IpMask     *string   `json:"ipMask,omitempty"`
}
