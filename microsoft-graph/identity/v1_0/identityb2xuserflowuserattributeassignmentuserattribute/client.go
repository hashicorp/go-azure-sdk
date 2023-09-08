package identityb2xuserflowuserattributeassignmentuserattribute

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityB2xUserFlowUserAttributeAssignmentUserAttributeClient struct {
	Client *msgraph.Client
}

func NewIdentityB2xUserFlowUserAttributeAssignmentUserAttributeClientWithBaseURI(api sdkEnv.Api) (*IdentityB2xUserFlowUserAttributeAssignmentUserAttributeClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityb2xuserflowuserattributeassignmentuserattribute", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityB2xUserFlowUserAttributeAssignmentUserAttributeClient: %+v", err)
	}

	return &IdentityB2xUserFlowUserAttributeAssignmentUserAttributeClient{
		Client: client,
	}, nil
}
