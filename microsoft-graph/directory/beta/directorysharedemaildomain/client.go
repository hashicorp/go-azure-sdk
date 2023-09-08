package directorysharedemaildomain

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectorySharedEmailDomainClient struct {
	Client *msgraph.Client
}

func NewDirectorySharedEmailDomainClientWithBaseURI(api sdkEnv.Api) (*DirectorySharedEmailDomainClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "directorysharedemaildomain", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DirectorySharedEmailDomainClient: %+v", err)
	}

	return &DirectorySharedEmailDomainClient{
		Client: client,
	}, nil
}
