package windowsinformationprotectiondeviceregistration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsInformationProtectionDeviceRegistrationClient struct {
	Client *msgraph.Client
}

func NewWindowsInformationProtectionDeviceRegistrationClientWithBaseURI(sdkApi sdkEnv.Api) (*WindowsInformationProtectionDeviceRegistrationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "windowsinformationprotectiondeviceregistration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WindowsInformationProtectionDeviceRegistrationClient: %+v", err)
	}

	return &WindowsInformationProtectionDeviceRegistrationClient{
		Client: client,
	}, nil
}
