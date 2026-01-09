package templatecategorysettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateCategorySettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewTemplateCategorySettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateCategorySettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatecategorysettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateCategorySettingDefinitionClient: %+v", err)
	}

	return &TemplateCategorySettingDefinitionClient{
		Client: client,
	}, nil
}
