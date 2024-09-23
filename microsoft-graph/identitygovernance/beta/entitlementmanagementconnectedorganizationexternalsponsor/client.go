package entitlementmanagementconnectedorganizationexternalsponsor

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementConnectedOrganizationExternalSponsorClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementConnectedOrganizationExternalSponsorClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementConnectedOrganizationExternalSponsorClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementconnectedorganizationexternalsponsor", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementConnectedOrganizationExternalSponsorClient: %+v", err)
	}

	return &EntitlementManagementConnectedOrganizationExternalSponsorClient{
		Client: client,
	}, nil
}
