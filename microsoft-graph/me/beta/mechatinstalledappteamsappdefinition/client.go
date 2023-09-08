package mechatinstalledappteamsappdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatInstalledAppTeamsAppDefinitionClient struct {
	Client *msgraph.Client
}

func NewMeChatInstalledAppTeamsAppDefinitionClientWithBaseURI(api sdkEnv.Api) (*MeChatInstalledAppTeamsAppDefinitionClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatinstalledappteamsappdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatInstalledAppTeamsAppDefinitionClient: %+v", err)
	}

	return &MeChatInstalledAppTeamsAppDefinitionClient{
		Client: client,
	}, nil
}
