package meteamworkinstalledappchat

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeTeamworkInstalledAppChatClient struct {
	Client *msgraph.Client
}

func NewMeTeamworkInstalledAppChatClientWithBaseURI(api sdkEnv.Api) (*MeTeamworkInstalledAppChatClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meteamworkinstalledappchat", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeTeamworkInstalledAppChatClient: %+v", err)
	}

	return &MeTeamworkInstalledAppChatClient{
		Client: client,
	}, nil
}
