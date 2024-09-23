package hardwareconfigurationuserrunstate

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareConfigurationUserRunStateClient struct {
	Client *msgraph.Client
}

func NewHardwareConfigurationUserRunStateClientWithBaseURI(sdkApi sdkEnv.Api) (*HardwareConfigurationUserRunStateClient, error) {
	client, err := msgraph.NewClient(sdkApi, "hardwareconfigurationuserrunstate", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HardwareConfigurationUserRunStateClient: %+v", err)
	}

	return &HardwareConfigurationUserRunStateClient{
		Client: client,
	}, nil
}
