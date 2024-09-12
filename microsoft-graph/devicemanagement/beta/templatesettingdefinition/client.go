package templatesettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateSettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewTemplateSettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateSettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatesettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateSettingDefinitionClient: %+v", err)
	}

	return &TemplateSettingDefinitionClient{
		Client: client,
	}, nil
}
