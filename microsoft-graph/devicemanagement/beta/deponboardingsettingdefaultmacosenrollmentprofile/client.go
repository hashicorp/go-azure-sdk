package deponboardingsettingdefaultmacosenrollmentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSettingDefaultMacOsEnrollmentProfileClient struct {
	Client *msgraph.Client
}

func NewDepOnboardingSettingDefaultMacOsEnrollmentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*DepOnboardingSettingDefaultMacOsEnrollmentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deponboardingsettingdefaultmacosenrollmentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DepOnboardingSettingDefaultMacOsEnrollmentProfileClient: %+v", err)
	}

	return &DepOnboardingSettingDefaultMacOsEnrollmentProfileClient{
		Client: client,
	}, nil
}
