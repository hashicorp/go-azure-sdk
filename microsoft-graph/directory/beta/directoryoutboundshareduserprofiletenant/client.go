package directoryoutboundshareduserprofiletenant

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryOutboundSharedUserProfileTenantClient struct {
	Client *msgraph.Client
}

func NewDirectoryOutboundSharedUserProfileTenantClientWithBaseURI(api sdkEnv.Api) (*DirectoryOutboundSharedUserProfileTenantClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryoutboundshareduserprofiletenant", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryOutboundSharedUserProfileTenantClient: %+v", err)
	}

	return &DirectoryOutboundSharedUserProfileTenantClient{
		Client: client,
	}, nil
}
