package settingwindowinstance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingWindowInstanceClient struct {
	Client *msgraph.Client
}

func NewSettingWindowInstanceClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingWindowInstanceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingwindowinstance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingWindowInstanceClient: %+v", err)
	}

	return &SettingWindowInstanceClient{
		Client: client,
	}, nil
}
