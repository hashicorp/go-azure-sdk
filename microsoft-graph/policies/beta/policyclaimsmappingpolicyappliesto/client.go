package policyclaimsmappingpolicyappliesto

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyClaimsMappingPolicyAppliesToClient struct {
	Client *msgraph.Client
}

func NewPolicyClaimsMappingPolicyAppliesToClientWithBaseURI(api sdkEnv.Api) (*PolicyClaimsMappingPolicyAppliesToClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "policyclaimsmappingpolicyappliesto", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PolicyClaimsMappingPolicyAppliesToClient: %+v", err)
	}

	return &PolicyClaimsMappingPolicyAppliesToClient{
		Client: client,
	}, nil
}
