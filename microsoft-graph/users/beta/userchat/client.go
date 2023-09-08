package userchat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatClient struct {
	Client *msgraph.Client
}

func NewUserChatClientWithBaseURI(api sdkEnv.Api) (*UserChatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatClient: %+v", err)
	}

	return &UserChatClient{
		Client: client,
	}, nil
}
