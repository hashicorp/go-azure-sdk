package configurationsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConfigurationSettingClient struct {
	Client *msgraph.Client
}

func NewConfigurationSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*ConfigurationSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "configurationsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ConfigurationSettingClient: %+v", err)
	}

	return &ConfigurationSettingClient{
		Client: client,
	}, nil
}
