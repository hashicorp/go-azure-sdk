package userchatmessagereply

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatMessageReplyClient struct {
	Client *msgraph.Client
}

func NewUserChatMessageReplyClientWithBaseURI(api sdkEnv.Api) (*UserChatMessageReplyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatmessagereply", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatMessageReplyClient: %+v", err)
	}

	return &UserChatMessageReplyClient{
		Client: client,
	}, nil
}
