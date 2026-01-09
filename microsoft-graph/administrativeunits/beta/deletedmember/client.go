package deletedmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeletedMemberClient struct {
	Client *msgraph.Client
}

func NewDeletedMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*DeletedMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deletedmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeletedMemberClient: %+v", err)
	}

	return &DeletedMemberClient{
		Client: client,
	}, nil
}
