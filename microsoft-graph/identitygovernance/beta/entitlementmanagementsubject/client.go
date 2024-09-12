package entitlementmanagementsubject

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntitlementManagementSubjectClient struct {
	Client *msgraph.Client
}

func NewEntitlementManagementSubjectClientWithBaseURI(sdkApi sdkEnv.Api) (*EntitlementManagementSubjectClient, error) {
	client, err := msgraph.NewClient(sdkApi, "entitlementmanagementsubject", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating EntitlementManagementSubjectClient: %+v", err)
	}

	return &EntitlementManagementSubjectClient{
		Client: client,
	}, nil
}
