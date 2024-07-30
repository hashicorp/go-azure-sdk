package crosstenantaccesspolicypartner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyPartnerClient struct {
	Client *msgraph.Client
}

func NewCrossTenantAccessPolicyPartnerClientWithBaseURI(sdkApi sdkEnv.Api) (*CrossTenantAccessPolicyPartnerClient, error) {
	client, err := msgraph.NewClient(sdkApi, "crosstenantaccesspolicypartner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CrossTenantAccessPolicyPartnerClient: %+v", err)
	}

	return &CrossTenantAccessPolicyPartnerClient{
		Client: client,
	}, nil
}
