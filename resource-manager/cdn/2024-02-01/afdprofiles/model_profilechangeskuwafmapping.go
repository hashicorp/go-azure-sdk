package afdprofiles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProfileChangeSkuWafMapping struct {
	ChangeToWafPolicy  ResourceReference `json:"changeToWafPolicy"`
	SecurityPolicyName string            `json:"securityPolicyName"`
}
