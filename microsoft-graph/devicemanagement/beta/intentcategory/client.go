package intentcategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntentCategoryClient struct {
	Client *msgraph.Client
}

func NewIntentCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*IntentCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "intentcategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntentCategoryClient: %+v", err)
	}

	return &IntentCategoryClient{
		Client: client,
	}, nil
}
