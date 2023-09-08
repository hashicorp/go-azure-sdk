package userchatmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatMessageClient struct {
	Client *msgraph.Client
}

func NewUserChatMessageClientWithBaseURI(api sdkEnv.Api) (*UserChatMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatMessageClient: %+v", err)
	}

	return &UserChatMessageClient{
		Client: client,
	}, nil
}
