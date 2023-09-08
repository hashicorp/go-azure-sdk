package groupowner

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupOwnerClient struct {
	Client *msgraph.Client
}

func NewGroupOwnerClientWithBaseURI(api sdkEnv.Api) (*GroupOwnerClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupowner", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupOwnerClient: %+v", err)
	}

	return &GroupOwnerClient{
		Client: client,
	}, nil
}
