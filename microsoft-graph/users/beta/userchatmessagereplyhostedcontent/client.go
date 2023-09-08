package userchatmessagereplyhostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatMessageReplyHostedContentClient struct {
	Client *msgraph.Client
}

func NewUserChatMessageReplyHostedContentClientWithBaseURI(api sdkEnv.Api) (*UserChatMessageReplyHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatmessagereplyhostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatMessageReplyHostedContentClient: %+v", err)
	}

	return &UserChatMessageReplyHostedContentClient{
		Client: client,
	}, nil
}
