package settingexchange

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingExchangeClient struct {
	Client *msgraph.Client
}

func NewSettingExchangeClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingExchangeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingexchange", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingExchangeClient: %+v", err)
	}

	return &SettingExchangeClient{
		Client: client,
	}, nil
}
