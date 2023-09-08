package policyclaimsmappingpolicy

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyClaimsMappingPolicyClient struct {
	Client *msgraph.Client
}

func NewPolicyClaimsMappingPolicyClientWithBaseURI(api sdkEnv.Api) (*PolicyClaimsMappingPolicyClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyclaimsmappingpolicy", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyClaimsMappingPolicyClient: %+v", err)
	}

	return &PolicyClaimsMappingPolicyClient{
		Client: client,
	}, nil
}
