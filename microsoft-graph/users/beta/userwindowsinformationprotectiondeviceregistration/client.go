package userwindowsinformationprotectiondeviceregistration

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserWindowsInformationProtectionDeviceRegistrationClient struct {
	Client *msgraph.Client
}

func NewUserWindowsInformationProtectionDeviceRegistrationClientWithBaseURI(api sdkEnv.Api) (*UserWindowsInformationProtectionDeviceRegistrationClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userwindowsinformationprotectiondeviceregistration", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserWindowsInformationProtectionDeviceRegistrationClient: %+v", err)
	}

	return &UserWindowsInformationProtectionDeviceRegistrationClient{
		Client: client,
	}, nil
}
