package userchattabteamsapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatTabTeamsAppClient struct {
	Client *msgraph.Client
}

func NewUserChatTabTeamsAppClientWithBaseURI(api sdkEnv.Api) (*UserChatTabTeamsAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchattabteamsapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatTabTeamsAppClient: %+v", err)
	}

	return &UserChatTabTeamsAppClient{
		Client: client,
	}, nil
}
