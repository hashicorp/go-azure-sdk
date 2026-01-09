package crosstenantaccesspolicypartneridentitysynchronization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyPartnerIdentitySynchronizationClient struct {
	Client *msgraph.Client
}

func NewCrossTenantAccessPolicyPartnerIdentitySynchronizationClientWithBaseURI(sdkApi sdkEnv.Api) (*CrossTenantAccessPolicyPartnerIdentitySynchronizationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "crosstenantaccesspolicypartneridentitysynchronization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CrossTenantAccessPolicyPartnerIdentitySynchronizationClient: %+v", err)
	}

	return &CrossTenantAccessPolicyPartnerIdentitySynchronizationClient{
		Client: client,
	}, nil
}
