package templatemigratabletocategoryrecommendedsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateMigratableToCategoryRecommendedSettingClient struct {
	Client *msgraph.Client
}

func NewTemplateMigratableToCategoryRecommendedSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateMigratableToCategoryRecommendedSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatemigratabletocategoryrecommendedsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateMigratableToCategoryRecommendedSettingClient: %+v", err)
	}

	return &TemplateMigratableToCategoryRecommendedSettingClient{
		Client: client,
	}, nil
}
