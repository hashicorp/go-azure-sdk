package settingiteminsight

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingItemInsightClient struct {
	Client *msgraph.Client
}

func NewSettingItemInsightClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingItemInsightClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingiteminsight", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingItemInsightClient: %+v", err)
	}

	return &SettingItemInsightClient{
		Client: client,
	}, nil
}
