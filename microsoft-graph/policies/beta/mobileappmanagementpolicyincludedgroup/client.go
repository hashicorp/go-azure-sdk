package mobileappmanagementpolicyincludedgroup

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppManagementPolicyIncludedGroupClient struct {
	Client *msgraph.Client
}

func NewMobileAppManagementPolicyIncludedGroupClientWithBaseURI(sdkApi sdkEnv.Api) (*MobileAppManagementPolicyIncludedGroupClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mobileappmanagementpolicyincludedgroup", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MobileAppManagementPolicyIncludedGroupClient: %+v", err)
	}

	return &MobileAppManagementPolicyIncludedGroupClient{
		Client: client,
	}, nil
}
