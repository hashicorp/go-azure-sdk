package userchatlastmessagepreview

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatLastMessagePreviewClient struct {
	Client *msgraph.Client
}

func NewUserChatLastMessagePreviewClientWithBaseURI(api sdkEnv.Api) (*UserChatLastMessagePreviewClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatlastmessagepreview", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatLastMessagePreviewClient: %+v", err)
	}

	return &UserChatLastMessagePreviewClient{
		Client: client,
	}, nil
}
