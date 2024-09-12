package mobileappmanagementpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppManagementPolicyClient struct {
	Client *msgraph.Client
}

func NewMobileAppManagementPolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*MobileAppManagementPolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "mobileappmanagementpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MobileAppManagementPolicyClient: %+v", err)
	}

	return &MobileAppManagementPolicyClient{
		Client: client,
	}, nil
}
