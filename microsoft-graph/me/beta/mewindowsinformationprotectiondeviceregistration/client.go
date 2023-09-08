package mewindowsinformationprotectiondeviceregistration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeWindowsInformationProtectionDeviceRegistrationClient struct {
	Client *msgraph.Client
}

func NewMeWindowsInformationProtectionDeviceRegistrationClientWithBaseURI(api sdkEnv.Api) (*MeWindowsInformationProtectionDeviceRegistrationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "mewindowsinformationprotectiondeviceregistration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeWindowsInformationProtectionDeviceRegistrationClient: %+v", err)
	}

	return &MeWindowsInformationProtectionDeviceRegistrationClient{
		Client: client,
	}, nil
}
