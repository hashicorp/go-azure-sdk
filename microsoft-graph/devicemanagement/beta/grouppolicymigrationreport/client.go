package grouppolicymigrationreport

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyMigrationReportClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyMigrationReportClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyMigrationReportClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicymigrationreport", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyMigrationReportClient: %+v", err)
	}

	return &GroupPolicyMigrationReportClient{
		Client: client,
	}, nil
}
