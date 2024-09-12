package categorysettingdefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CategorySettingDefinitionClient struct {
	Client *msgraph.Client
}

func NewCategorySettingDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*CategorySettingDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "categorysettingdefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CategorySettingDefinitionClient: %+v", err)
	}

	return &CategorySettingDefinitionClient{
		Client: client,
	}, nil
}
