package hardwarepasswordinfo

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwarePasswordInfoClient struct {
	Client *msgraph.Client
}

func NewHardwarePasswordInfoClientWithBaseURI(sdkApi sdkEnv.Api) (*HardwarePasswordInfoClient, error) {
	client, err := msgraph.NewClient(sdkApi, "hardwarepasswordinfo", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating HardwarePasswordInfoClient: %+v", err)
	}

	return &HardwarePasswordInfoClient{
		Client: client,
	}, nil
}
