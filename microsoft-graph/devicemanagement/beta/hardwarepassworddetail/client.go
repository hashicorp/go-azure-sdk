package hardwarepassworddetail

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwarePasswordDetailClient struct {
	Client *msgraph.Client
}

func NewHardwarePasswordDetailClientWithBaseURI(sdkApi sdkEnv.Api) (*HardwarePasswordDetailClient, error) {
	client, err := msgraph.NewClient(sdkApi, "hardwarepassworddetail", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HardwarePasswordDetailClient: %+v", err)
	}

	return &HardwarePasswordDetailClient{
		Client: client,
	}, nil
}
