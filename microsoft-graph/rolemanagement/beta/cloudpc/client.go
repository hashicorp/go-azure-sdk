package cloudpc

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCClient struct {
	Client *msgraph.Client
}

func NewCloudPCClientWithBaseURI(sdkApi sdkEnv.Api) (*CloudPCClient, error) {
	client, err := msgraph.NewClient(sdkApi, "cloudpc", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CloudPCClient: %+v", err)
	}

	return &CloudPCClient{
		Client: client,
	}, nil
}
