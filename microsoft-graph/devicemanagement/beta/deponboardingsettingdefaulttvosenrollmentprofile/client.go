package deponboardingsettingdefaulttvosenrollmentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSettingDefaultTvOSEnrollmentProfileClient struct {
	Client *msgraph.Client
}

func NewDepOnboardingSettingDefaultTvOSEnrollmentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*DepOnboardingSettingDefaultTvOSEnrollmentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deponboardingsettingdefaulttvosenrollmentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DepOnboardingSettingDefaultTvOSEnrollmentProfileClient: %+v", err)
	}

	return &DepOnboardingSettingDefaultTvOSEnrollmentProfileClient{
		Client: client,
	}, nil
}
