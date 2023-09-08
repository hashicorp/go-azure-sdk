package directoryinboundshareduserprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryInboundSharedUserProfileClient struct {
	Client *msgraph.Client
}

func NewDirectoryInboundSharedUserProfileClientWithBaseURI(api sdkEnv.Api) (*DirectoryInboundSharedUserProfileClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryinboundshareduserprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryInboundSharedUserProfileClient: %+v", err)
	}

	return &DirectoryInboundSharedUserProfileClient{
		Client: client,
	}, nil
}
