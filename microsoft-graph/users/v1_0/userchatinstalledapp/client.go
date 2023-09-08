package userchatinstalledapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatInstalledAppClient struct {
	Client *msgraph.Client
}

func NewUserChatInstalledAppClientWithBaseURI(api sdkEnv.Api) (*UserChatInstalledAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatinstalledapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatInstalledAppClient: %+v", err)
	}

	return &UserChatInstalledAppClient{
		Client: client,
	}, nil
}
