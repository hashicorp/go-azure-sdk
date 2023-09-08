package usercloudpc

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserCloudPCClient struct {
	Client *msgraph.Client
}

func NewUserCloudPCClientWithBaseURI(api sdkEnv.Api) (*UserCloudPCClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "usercloudpc", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserCloudPCClient: %+v", err)
	}

	return &UserCloudPCClient{
		Client: client,
	}, nil
}
