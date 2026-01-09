package deponboardingsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DepOnboardingSettingClient struct {
	Client *msgraph.Client
}

func NewDepOnboardingSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*DepOnboardingSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "deponboardingsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DepOnboardingSettingClient: %+v", err)
	}

	return &DepOnboardingSettingClient{
		Client: client,
	}, nil
}
