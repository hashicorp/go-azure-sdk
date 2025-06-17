package datasecurityandgovernanceprotectionscope

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSecurityAndGovernanceProtectionScopeClient struct {
	Client *msgraph.Client
}

func NewDataSecurityAndGovernanceProtectionScopeClientWithBaseURI(sdkApi sdkEnv.Api) (*DataSecurityAndGovernanceProtectionScopeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "datasecurityandgovernanceprotectionscope", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataSecurityAndGovernanceProtectionScopeClient: %+v", err)
	}

	return &DataSecurityAndGovernanceProtectionScopeClient{
		Client: client,
	}, nil
}
