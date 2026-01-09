package crosstenantaccesspolicytemplatemultitenantorganizationidentitysynchronization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient struct {
	Client *msgraph.Client
}

func NewCrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClientWithBaseURI(sdkApi sdkEnv.Api) (*CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "crosstenantaccesspolicytemplatemultitenantorganizationidentitysynchronization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient: %+v", err)
	}

	return &CrossTenantAccessPolicyTemplateMultiTenantOrganizationIdentitySynchronizationClient{
		Client: client,
	}, nil
}
