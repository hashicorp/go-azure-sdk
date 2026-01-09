package grouppolicymigrationreportunsupportedgrouppolicyextension

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient struct {
	Client *msgraph.Client
}

func NewGroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClientWithBaseURI(sdkApi sdkEnv.Api) (*GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient, error) {
	client, err := msgraph.NewClient(sdkApi, "grouppolicymigrationreportunsupportedgrouppolicyextension", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient: %+v", err)
	}

	return &GroupPolicyMigrationReportUnsupportedGroupPolicyExtensionClient{
		Client: client,
	}, nil
}
