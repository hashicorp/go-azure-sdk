package category

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CategoryClient struct {
	Client *msgraph.Client
}

func NewCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*CategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "category", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CategoryClient: %+v", err)
	}

	return &CategoryClient{
		Client: client,
	}, nil
}
