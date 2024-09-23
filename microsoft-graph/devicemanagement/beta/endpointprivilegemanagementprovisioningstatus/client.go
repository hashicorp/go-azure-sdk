package endpointprivilegemanagementprovisioningstatus

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndpointPrivilegeManagementProvisioningStatusClient struct {
	Client *msgraph.Client
}

func NewEndpointPrivilegeManagementProvisioningStatusClientWithBaseURI(sdkApi sdkEnv.Api) (*EndpointPrivilegeManagementProvisioningStatusClient, error) {
	client, err := msgraph.NewClient(sdkApi, "endpointprivilegemanagementprovisioningstatus", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EndpointPrivilegeManagementProvisioningStatusClient: %+v", err)
	}

	return &EndpointPrivilegeManagementProvisioningStatusClient{
		Client: client,
	}, nil
}
