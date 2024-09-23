package claimsmappingpolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClaimsMappingPolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewClaimsMappingPolicyAppliesToClientWithBaseURI(sdkApi sdkEnv.Api) (*ClaimsMappingPolicyAppliesToClient, error) {
	client, err := msgraph.NewClient(sdkApi, "claimsmappingpolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ClaimsMappingPolicyAppliesToClient: %+v", err)
	}

	return &ClaimsMappingPolicyAppliesToClient{
		Client: client,
	}, nil
}
