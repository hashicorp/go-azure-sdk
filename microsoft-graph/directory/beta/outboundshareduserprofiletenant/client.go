package outboundshareduserprofiletenant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutboundSharedUserProfileTenantClient struct {
	Client *msgraph.Client
}

func NewOutboundSharedUserProfileTenantClientWithBaseURI(sdkApi sdkEnv.Api) (*OutboundSharedUserProfileTenantClient, error) {
	client, err := msgraph.NewClient(sdkApi, "outboundshareduserprofiletenant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OutboundSharedUserProfileTenantClient: %+v", err)
	}

	return &OutboundSharedUserProfileTenantClient{
		Client: client,
	}, nil
}
