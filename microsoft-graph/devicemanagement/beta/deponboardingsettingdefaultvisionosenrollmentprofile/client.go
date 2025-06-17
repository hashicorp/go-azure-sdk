package deponboardingsettingdefaultvisionosenrollmentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSettingDefaultVisionOSEnrollmentProfileClient struct {
	Client *msgraph.Client
}

func NewDepOnboardingSettingDefaultVisionOSEnrollmentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*DepOnboardingSettingDefaultVisionOSEnrollmentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deponboardingsettingdefaultvisionosenrollmentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DepOnboardingSettingDefaultVisionOSEnrollmentProfileClient: %+v", err)
	}

	return &DepOnboardingSettingDefaultVisionOSEnrollmentProfileClient{
		Client: client,
	}, nil
}
