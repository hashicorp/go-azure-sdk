package userchatpinnedmessage

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatPinnedMessageClient struct {
	Client *msgraph.Client
}

func NewUserChatPinnedMessageClientWithBaseURI(api sdkEnv.Api) (*UserChatPinnedMessageClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatpinnedmessage", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatPinnedMessageClient: %+v", err)
	}

	return &UserChatPinnedMessageClient{
		Client: client,
	}, nil
}
