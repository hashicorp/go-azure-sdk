package hardwareconfigurationrunsummary

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareConfigurationRunSummaryClient struct {
	Client *msgraph.Client
}

func NewHardwareConfigurationRunSummaryClientWithBaseURI(sdkApi sdkEnv.Api) (*HardwareConfigurationRunSummaryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "hardwareconfigurationrunsummary", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HardwareConfigurationRunSummaryClient: %+v", err)
	}

	return &HardwareConfigurationRunSummaryClient{
		Client: client,
	}, nil
}
