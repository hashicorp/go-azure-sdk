package policymobileappmanagementpolicyincludedgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyMobileAppManagementPolicyIncludedGroupClient struct {
	Client *msgraph.Client
}

func NewPolicyMobileAppManagementPolicyIncludedGroupClientWithBaseURI(api sdkEnv.Api) (*PolicyMobileAppManagementPolicyIncludedGroupClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policymobileappmanagementpolicyincludedgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyMobileAppManagementPolicyIncludedGroupClient: %+v", err)
	}

	return &PolicyMobileAppManagementPolicyIncludedGroupClient{
		Client: client,
	}, nil
}
