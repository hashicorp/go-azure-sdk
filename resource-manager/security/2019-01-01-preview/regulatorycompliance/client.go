package regulatorycompliance

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegulatoryComplianceClient struct {
	Client *resourcemanager.Client
}

func NewRegulatoryComplianceClientWithBaseURI(sdkApi sdkEnv.Api) (*RegulatoryComplianceClient, error) {
	client, err := resourcemanager.NewResourceManagerClient(sdkApi, "regulatorycompliance", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating RegulatoryComplianceClient: %+v", err)
	}

	return &RegulatoryComplianceClient{
		Client: client,
	}, nil
}
