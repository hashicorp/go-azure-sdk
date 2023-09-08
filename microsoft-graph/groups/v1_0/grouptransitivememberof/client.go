package grouptransitivememberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTransitiveMemberOfClient struct {
	Client *msgraph.Client
}

func NewGroupTransitiveMemberOfClientWithBaseURI(api sdkEnv.Api) (*GroupTransitiveMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "grouptransitivememberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupTransitiveMemberOfClient: %+v", err)
	}

	return &GroupTransitiveMemberOfClient{
		Client: client,
	}, nil
}
