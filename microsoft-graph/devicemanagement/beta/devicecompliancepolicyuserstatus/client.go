package devicecompliancepolicyuserstatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyUserStatusClient struct {
	Client *msgraph.Client
}

func NewDeviceCompliancePolicyUserStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*DeviceCompliancePolicyUserStatusClient, error) {
	client, err := msgraph.NewClient(sdkApi, "devicecompliancepolicyuserstatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DeviceCompliancePolicyUserStatusClient: %+v", err)
	}

	return &DeviceCompliancePolicyUserStatusClient{
		Client: client,
	}, nil
}
