package entitlementmanagementconnectedorganizationinternalsponsor

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementConnectedOrganizationInternalSponsorClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementConnectedOrganizationInternalSponsorClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementConnectedOrganizationInternalSponsorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementconnectedorganizationinternalsponsor", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementConnectedOrganizationInternalSponsorClient: %+v", err)
	}

	return &EntitlementManagementConnectedOrganizationInternalSponsorClient{
		Client: client,
	}, nil
}
