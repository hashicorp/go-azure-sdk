package hardwareconfiguration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareConfigurationClient struct {
	Client *msgraph.Client
}

func NewHardwareConfigurationClientWithBaseURI(sdkApi sdkEnv.Api) (*HardwareConfigurationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "hardwareconfiguration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HardwareConfigurationClient: %+v", err)
	}

	return &HardwareConfigurationClient{
		Client: client,
	}, nil
}
