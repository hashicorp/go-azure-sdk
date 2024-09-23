package hardwareconfigurationassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareConfigurationAssignmentClient struct {
	Client *msgraph.Client
}

func NewHardwareConfigurationAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*HardwareConfigurationAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "hardwareconfigurationassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HardwareConfigurationAssignmentClient: %+v", err)
	}

	return &HardwareConfigurationAssignmentClient{
		Client: client,
	}, nil
}
