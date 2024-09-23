package crosstenantaccesspolicydefault

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyDefaultClient struct {
	Client *msgraph.Client
}

func NewCrossTenantAccessPolicyDefaultClientWithBaseURI(sdkApi sdkEnv.Api) (*CrossTenantAccessPolicyDefaultClient, error) {
	client, err := msgraph.NewClient(sdkApi, "crosstenantaccesspolicydefault", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CrossTenantAccessPolicyDefaultClient: %+v", err)
	}

	return &CrossTenantAccessPolicyDefaultClient{
		Client: client,
	}, nil
}
