package deponboardingsettingimportedappledeviceidentity

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSettingImportedAppleDeviceIdentityClient struct {
	Client *msgraph.Client
}

func NewDepOnboardingSettingImportedAppleDeviceIdentityClientWithBaseURI(sdkApi sdkEnv.Api) (*DepOnboardingSettingImportedAppleDeviceIdentityClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deponboardingsettingimportedappledeviceidentity", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DepOnboardingSettingImportedAppleDeviceIdentityClient: %+v", err)
	}

	return &DepOnboardingSettingImportedAppleDeviceIdentityClient{
		Client: client,
	}, nil
}
