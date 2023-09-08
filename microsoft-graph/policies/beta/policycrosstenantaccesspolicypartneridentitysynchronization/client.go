package policycrosstenantaccesspolicypartneridentitysynchronization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyCrossTenantAccessPolicyPartnerIdentitySynchronizationClient struct {
	Client *msgraph.Client
}

func NewPolicyCrossTenantAccessPolicyPartnerIdentitySynchronizationClientWithBaseURI(api sdkEnv.Api) (*PolicyCrossTenantAccessPolicyPartnerIdentitySynchronizationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policycrosstenantaccesspolicypartneridentitysynchronization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyCrossTenantAccessPolicyPartnerIdentitySynchronizationClient: %+v", err)
	}

	return &PolicyCrossTenantAccessPolicyPartnerIdentitySynchronizationClient{
		Client: client,
	}, nil
}
