package virtualeventwebinar

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualEventWebinarClient struct {
	Client *msgraph.Client
}

func NewVirtualEventWebinarClientWithBaseURI(sdkApi sdkEnv.Api) (*VirtualEventWebinarClient, error) {
	client, err := msgraph.NewClient(sdkApi, "virtualeventwebinar", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating VirtualEventWebinarClient: %+v", err)
	}

	return &VirtualEventWebinarClient{
		Client: client,
	}, nil
}
