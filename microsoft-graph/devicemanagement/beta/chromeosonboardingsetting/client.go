package chromeosonboardingsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChromeOSOnboardingSettingClient struct {
	Client *msgraph.Client
}

func NewChromeOSOnboardingSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*ChromeOSOnboardingSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "chromeosonboardingsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ChromeOSOnboardingSettingClient: %+v", err)
	}

	return &ChromeOSOnboardingSettingClient{
		Client: client,
	}, nil
}
