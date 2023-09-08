package userchatmessagehostedcontent

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatMessageHostedContentClient struct {
	Client *msgraph.Client
}

func NewUserChatMessageHostedContentClientWithBaseURI(api sdkEnv.Api) (*UserChatMessageHostedContentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatmessagehostedcontent", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatMessageHostedContentClient: %+v", err)
	}

	return &UserChatMessageHostedContentClient{
		Client: client,
	}, nil
}
