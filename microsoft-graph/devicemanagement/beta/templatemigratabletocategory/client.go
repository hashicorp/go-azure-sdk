package templatemigratabletocategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemplateMigratableToCategoryClient struct {
	Client *msgraph.Client
}

func NewTemplateMigratableToCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*TemplateMigratableToCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "templatemigratabletocategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TemplateMigratableToCategoryClient: %+v", err)
	}

	return &TemplateMigratableToCategoryClient{
		Client: client,
	}, nil
}
