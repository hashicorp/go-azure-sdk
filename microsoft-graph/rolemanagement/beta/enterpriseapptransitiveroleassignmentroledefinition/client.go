package enterpriseapptransitiveroleassignmentroledefinition

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppTransitiveRoleAssignmentRoleDefinitionClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapptransitiveroleassignmentroledefinition", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient: %+v", err)
	}

	return &EnterpriseAppTransitiveRoleAssignmentRoleDefinitionClient{
		Client: client,
	}, nil
}
