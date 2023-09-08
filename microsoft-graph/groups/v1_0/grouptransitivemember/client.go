package grouptransitivemember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTransitiveMemberClient struct {
	Client *msgraph.Client
}

func NewGroupTransitiveMemberClientWithBaseURI(api sdkEnv.Api) (*GroupTransitiveMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouptransitivemember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTransitiveMemberClient: %+v", err)
	}

	return &GroupTransitiveMemberClient{
		Client: client,
	}, nil
}
