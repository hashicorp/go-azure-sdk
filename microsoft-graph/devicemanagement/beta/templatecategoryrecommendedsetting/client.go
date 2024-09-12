package templatecategoryrecommendedsetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateCategoryRecommendedSettingClient struct {
	Client *msgraph.Client
}

func NewTemplateCategoryRecommendedSettingClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateCategoryRecommendedSettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatecategoryrecommendedsetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateCategoryRecommendedSettingClient: %+v", err)
	}

	return &TemplateCategoryRecommendedSettingClient{
		Client: client,
	}, nil
}
