package entitlementmanagementsubjectconnectedorganization

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementSubjectConnectedOrganizationClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementSubjectConnectedOrganizationClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementSubjectConnectedOrganizationClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementsubjectconnectedorganization", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementSubjectConnectedOrganizationClient: %+v", err)
	}

	return &EntitlementManagementSubjectConnectedOrganizationClient{
		Client: client,
	}, nil
}
