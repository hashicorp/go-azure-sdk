package templatecategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateCategoryClient struct {
	Client *msgraph.Client
}

func NewTemplateCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatecategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateCategoryClient: %+v", err)
	}

	return &TemplateCategoryClient{
		Client: client,
	}, nil
}
