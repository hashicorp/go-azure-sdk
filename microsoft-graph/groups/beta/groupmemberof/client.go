package groupmemberof

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupMemberOfClient struct {
	Client *msgraph.Client
}

func NewGroupMemberOfClientWithBaseURI(api sdkEnv.Api) (*GroupMemberOfClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "groupmemberof", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupMemberOfClient: %+v", err)
	}

	return &GroupMemberOfClient{
		Client: client,
	}, nil
}
