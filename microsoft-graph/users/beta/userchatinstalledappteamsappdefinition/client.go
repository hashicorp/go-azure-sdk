package userchatinstalledappteamsappdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserChatInstalledAppTeamsAppDefinitionClient struct {
	Client *msgraph.Client
}

func NewUserChatInstalledAppTeamsAppDefinitionClientWithBaseURI(api sdkEnv.Api) (*UserChatInstalledAppTeamsAppDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userchatinstalledappteamsappdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserChatInstalledAppTeamsAppDefinitionClient: %+v", err)
	}

	return &UserChatInstalledAppTeamsAppDefinitionClient{
		Client: client,
	}, nil
}
