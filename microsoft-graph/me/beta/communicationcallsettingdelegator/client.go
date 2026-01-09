package communicationcallsettingdelegator

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunicationCallSettingDelegatorClient struct {
	Client *msgraph.Client
}

func NewCommunicationCallSettingDelegatorClientWithBaseURI(sdkApi sdkEnv.Api) (*CommunicationCallSettingDelegatorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "communicationcallsettingdelegator", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CommunicationCallSettingDelegatorClient: %+v", err)
	}

	return &CommunicationCallSettingDelegatorClient{
		Client: client,
	}, nil
}
