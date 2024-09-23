package devicecompliancescript

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptClient struct {
	Client *msgraph.Client
}

func NewDeviceComplianceScriptClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceComplianceScriptClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancescript", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceComplianceScriptClient: %+v", err)
	}

	return &DeviceComplianceScriptClient{
		Client: client,
	}, nil
}
