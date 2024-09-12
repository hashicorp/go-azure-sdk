package intentcategorysetting

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntentCategorySettingClient struct {
	Client *msgraph.Client
}

func NewIntentCategorySettingClientWithBaseURI(sdkApi sdkEnv.Api) (*IntentCategorySettingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intentcategorysetting", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntentCategorySettingClient: %+v", err)
	}

	return &IntentCategorySettingClient{
		Client: client,
	}, nil
}
