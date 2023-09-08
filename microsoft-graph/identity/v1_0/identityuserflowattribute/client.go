package identityuserflowattribute

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityUserFlowAttributeClient struct {
	Client *msgraph.Client
}

func NewIdentityUserFlowAttributeClientWithBaseURI(api sdkEnv.Api) (*IdentityUserFlowAttributeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityuserflowattribute", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityUserFlowAttributeClient: %+v", err)
	}

	return &IdentityUserFlowAttributeClient{
		Client: client,
	}, nil
}
