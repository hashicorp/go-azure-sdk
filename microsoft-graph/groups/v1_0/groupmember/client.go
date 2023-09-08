package groupmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupMemberClient struct {
	Client *msgraph.Client
}

func NewGroupMemberClientWithBaseURI(api sdkEnv.Api) (*GroupMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupMemberClient: %+v", err)
	}

	return &GroupMemberClient{
		Client: client,
	}, nil
}
