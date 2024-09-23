package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitioncategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategoryClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionnextversiondefinitionpreviousversiondefinitioncategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategoryClient: %+v", err)
	}

	return &GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionCategoryClient{
		Client: client,
	}, nil
}
