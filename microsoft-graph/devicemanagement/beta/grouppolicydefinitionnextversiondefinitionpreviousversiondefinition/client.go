package grouppolicydefinitionnextversiondefinitionpreviousversiondefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionnextversiondefinitionpreviousversiondefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClient: %+v", err)
	}

	return &GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionClient{
		Client: client,
	}, nil
}
