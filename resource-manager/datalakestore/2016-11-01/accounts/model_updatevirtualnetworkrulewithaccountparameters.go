package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateVirtualNetworkRuleWithAccountParameters struct {
	Name       string                              `json:"name"`
	Properties *UpdateVirtualNetworkRuleProperties `json:"properties,omitempty"`
}
