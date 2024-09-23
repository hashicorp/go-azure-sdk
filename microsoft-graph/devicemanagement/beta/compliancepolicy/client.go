package compliancepolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompliancePolicyClient struct {
	Client *msgraph.Client
}

func NewCompliancePolicyClientWithBaseURI(sdkApi sdkEnv.Api) (*CompliancePolicyClient, error) {
	client, err := msgraph.NewClient(sdkApi, "compliancepolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CompliancePolicyClient: %+v", err)
	}

	return &CompliancePolicyClient{
		Client: client,
	}, nil
}
