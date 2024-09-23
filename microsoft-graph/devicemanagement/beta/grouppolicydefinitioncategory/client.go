package grouppolicydefinitioncategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionCategoryClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitioncategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionCategoryClient: %+v", err)
	}

	return &GroupPolicyDefinitionCategoryClient{
		Client: client,
	}, nil
}
