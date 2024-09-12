package deponboardingsettingenrollmentprofile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSettingEnrollmentProfileClient struct {
	Client *msgraph.Client
}

func NewDepOnboardingSettingEnrollmentProfileClientWithBaseURI(sdkApi sdkEnv.Api) (*DepOnboardingSettingEnrollmentProfileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deponboardingsettingenrollmentprofile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DepOnboardingSettingEnrollmentProfileClient: %+v", err)
	}

	return &DepOnboardingSettingEnrollmentProfileClient{
		Client: client,
	}, nil
}
