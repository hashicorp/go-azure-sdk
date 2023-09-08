package userchattab

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatTabClient struct {
	Client *msgraph.Client
}

func NewUserChatTabClientWithBaseURI(api sdkEnv.Api) (*UserChatTabClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchattab", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatTabClient: %+v", err)
	}

	return &UserChatTabClient{
		Client: client,
	}, nil
}
