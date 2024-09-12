package settingwindow

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingWindowClient struct {
	Client *msgraph.Client
}

func NewSettingWindowClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingWindowClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingwindow", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingWindowClient: %+v", err)
	}

	return &SettingWindowClient{
		Client: client,
	}, nil
}
