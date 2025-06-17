package communicationcallsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunicationCallSettingClient struct {
	Client *msgraph.Client
}

func NewCommunicationCallSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*CommunicationCallSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "communicationcallsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CommunicationCallSettingClient: %+v", err)
	}

	return &CommunicationCallSettingClient{
		Client: client,
	}, nil
}
