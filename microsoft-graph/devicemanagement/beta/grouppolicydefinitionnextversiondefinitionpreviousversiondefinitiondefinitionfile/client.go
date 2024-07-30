package grouppolicydefinitionnextversiondefinitionpreviousversiondefinitiondefinitionfile

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicydefinitionnextversiondefinitionpreviousversiondefinitiondefinitionfile", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClient: %+v", err)
	}

	return &GroupPolicyDefinitionNextVersionDefinitionPreviousVersionDefinitionDefinitionFileClient{
		Client: client,
	}, nil
}
