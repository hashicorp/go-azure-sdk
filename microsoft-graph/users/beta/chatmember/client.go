package chatmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMemberClient struct {
	Client *msgraph.Client
}

func NewChatMemberClientWithBaseURI(sdkApi sdkEnv.Api) (*ChatMemberClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chatmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChatMemberClient: %+v", err)
	}

	return &ChatMemberClient{
		Client: client,
	}, nil
}
