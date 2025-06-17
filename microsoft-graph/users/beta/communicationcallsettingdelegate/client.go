package communicationcallsettingdelegate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommunicationCallSettingDelegateClient struct {
	Client *msgraph.Client
}

func NewCommunicationCallSettingDelegateClientWithBaseURI(sdkApi sdkEnv.Api) (*CommunicationCallSettingDelegateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "communicationcallsettingdelegate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CommunicationCallSettingDelegateClient: %+v", err)
	}

	return &CommunicationCallSettingDelegateClient{
		Client: client,
	}, nil
}
