package informationprotectionpolicies

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InformationProtectionPoliciesClient struct {
	Client *resourcemanager.Client
}

func NewInformationProtectionPoliciesClientWithBaseURI(sdkApi sdkEnv.Api) (*InformationProtectionPoliciesClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "informationprotectionpolicies", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating InformationProtectionPoliciesClient: %+v", err)
	}

	return &InformationProtectionPoliciesClient{
		Client: client,
	}, nil
}
