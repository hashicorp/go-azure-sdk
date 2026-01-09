package intentcategorysettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntentCategorySettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewIntentCategorySettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*IntentCategorySettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intentcategorysettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntentCategorySettingDefinitionClient: %+v", err)
	}

	return &IntentCategorySettingDefinitionClient{
		Client: client,
	}, nil
}
