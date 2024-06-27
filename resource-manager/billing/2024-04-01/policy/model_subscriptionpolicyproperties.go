package policy

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionPolicyProperties struct {
	Policies          *[]PolicySummary   `json:"policies,omitempty"`
	ProvisioningState *ProvisioningState `json:"provisioningState,omitempty"`
}
