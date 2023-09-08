package mechatinstalledapp

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeChatInstalledAppClient struct {
	Client *msgraph.Client
}

func NewMeChatInstalledAppClientWithBaseURI(api sdkEnv.Api) (*MeChatInstalledAppClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mechatinstalledapp", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeChatInstalledAppClient: %+v", err)
	}

	return &MeChatInstalledAppClient{
		Client: client,
	}, nil
}
