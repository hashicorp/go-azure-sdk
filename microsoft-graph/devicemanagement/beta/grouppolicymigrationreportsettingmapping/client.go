package grouppolicymigrationreportsettingmapping

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyMigrationReportSettingMappingClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyMigrationReportSettingMappingClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyMigrationReportSettingMappingClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicymigrationreportsettingmapping", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyMigrationReportSettingMappingClient: %+v", err)
	}

	return &GroupPolicyMigrationReportSettingMappingClient{
		Client: client,
	}, nil
}
