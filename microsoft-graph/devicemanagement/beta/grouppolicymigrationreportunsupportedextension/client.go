package grouppolicymigrationreportunsupportedextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyMigrationReportUnsupportedExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyMigrationReportUnsupportedExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyMigrationReportUnsupportedExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicymigrationreportunsupportedextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyMigrationReportUnsupportedExtensionClient: %+v", err)
	}

	return &GroupPolicyMigrationReportUnsupportedExtensionClient{
		Client: client,
	}, nil
}
