package userchatmember

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatMemberClient struct {
	Client *msgraph.Client
}

func NewUserChatMemberClientWithBaseURI(api sdkEnv.Api) (*UserChatMemberClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatmember", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatMemberClient: %+v", err)
	}

	return &UserChatMemberClient{
		Client: client,
	}, nil
}
