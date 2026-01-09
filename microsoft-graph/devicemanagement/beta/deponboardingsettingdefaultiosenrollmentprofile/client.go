package deponboardingsettingdefaultiosenrollmentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSettingDefaultIosEnrollmentProfileClient struct {
	Client *msgraph.Client
}

func NewDepOnboardingSettingDefaultIosEnrollmentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*DepOnboardingSettingDefaultIosEnrollmentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deponboardingsettingdefaultiosenrollmentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DepOnboardingSettingDefaultIosEnrollmentProfileClient: %+v", err)
	}

	return &DepOnboardingSettingDefaultIosEnrollmentProfileClient{
		Client: client,
	}, nil
}
