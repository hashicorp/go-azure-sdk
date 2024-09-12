package enterpriseapptransitiveroleassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EnterpriseAppTransitiveRoleAssignmentClient struct {
	Client *msgraph.Client
}

func NewEnterpriseAppTransitiveRoleAssignmentClientWithBaseURI(sdkApi sdkEnv.Api) (*EnterpriseAppTransitiveRoleAssignmentClient, error) {
	client, err := msgraph.NewClient(sdkApi, "enterpriseapptransitiveroleassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EnterpriseAppTransitiveRoleAssignmentClient: %+v", err)
	}

	return &EnterpriseAppTransitiveRoleAssignmentClient{
		Client: client,
	}, nil
}
