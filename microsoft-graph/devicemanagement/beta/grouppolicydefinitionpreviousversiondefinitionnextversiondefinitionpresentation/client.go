package grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentation

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionpreviousversiondefinitionnextversiondefinitionpresentation", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient: %+v", err)
	}

	return &GroupPolicyDefinitionPreviousVersionDefinitionNextVersionDefinitionPresentationClient{
		Client: client,
	}, nil
}
