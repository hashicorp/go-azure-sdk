package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateVirtualEndpointProvisioningPolicyApplyConfigRequest struct {
	CloudPCIds     *[]string                                                                `json:"cloudPcIds,omitempty"`
	PolicySettings *CreateVirtualEndpointProvisioningPolicyApplyConfigRequestPolicySettings `json:"policySettings,omitempty"`
}
