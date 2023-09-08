package directoryoutboundshareduserprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectoryOutboundSharedUserProfileClient struct {
	Client *msgraph.Client
}

func NewDirectoryOutboundSharedUserProfileClientWithBaseURI(api sdkEnv.Api) (*DirectoryOutboundSharedUserProfileClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directoryoutboundshareduserprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectoryOutboundSharedUserProfileClient: %+v", err)
	}

	return &DirectoryOutboundSharedUserProfileClient{
		Client: client,
	}, nil
}
