package remoteassistancesetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemoteAssistanceSettingClient struct {
	Client *msgraph.Client
}

func NewRemoteAssistanceSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*RemoteAssistanceSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "remoteassistancesetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RemoteAssistanceSettingClient: %+v", err)
	}

	return &RemoteAssistanceSettingClient{
		Client: client,
	}, nil
}
