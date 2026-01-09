package settingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewSettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*SettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "settingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SettingDefinitionClient: %+v", err)
	}

	return &SettingDefinitionClient{
		Client: client,
	}, nil
}
