package grouppolicydefinitionpreviousversiondefinitionnextversiondefinitioncategory

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategoryClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategoryClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategoryClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionpreviousversiondefinitionnextversiondefinitioncategory", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategoryClient: %+v", err)
	}

	return &GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionCategoryClient{
		Client: client,
	}, nil
}
