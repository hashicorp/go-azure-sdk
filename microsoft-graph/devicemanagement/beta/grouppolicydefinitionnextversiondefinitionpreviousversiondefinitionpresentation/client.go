package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionnextversiondefinitionpreviousversiondefinitionpresentation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient: %+v", err)
	}

	return &GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionPresentationClient{
		Client: client,
	}, nil
}
