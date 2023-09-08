package policymobileappmanagementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyMobileAppManagementPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyMobileAppManagementPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyMobileAppManagementPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policymobileappmanagementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyMobileAppManagementPolicyClient: %+v", err)
	}

	return &PolicyMobileAppManagementPolicyClient{
		Client: client,
	}, nil
}
