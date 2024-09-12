package templatemigratabletocategorysettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateMigratableToCategorySettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewTemplateMigratableToCategorySettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateMigratableToCategorySettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatemigratabletocategorysettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateMigratableToCategorySettingDefinitionClient: %+v", err)
	}

	return &TemplateMigratableToCategorySettingDefinitionClient{
		Client: client,
	}, nil
}
