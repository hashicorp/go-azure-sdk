package settingshiftpreference

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingShiftPreferenceClient struct {
	Client *msgraph.Client
}

func NewSettingShiftPreferenceClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingShiftPreferenceClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingshiftpreference", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingShiftPreferenceClient: %+v", err)
	}

	return &SettingShiftPreferenceClient{
		Client: client,
	}, nil
}
