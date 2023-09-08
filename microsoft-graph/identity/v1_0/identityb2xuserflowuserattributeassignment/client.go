package identityb2xuserflowuserattributeassignment

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityB2xUserFlowUserAttributeAssignmentClient struct {
	Client *msgraph.Client
}

func NewIdentityB2xUserFlowUserAttributeAssignmentClientWithBaseURI(api sdkEnv.Api) (*IdentityB2xUserFlowUserAttributeAssignmentClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "identityb2xuserflowuserattributeassignment", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IdentityB2xUserFlowUserAttributeAssignmentClient: %+v", err)
	}

	return &IdentityB2xUserFlowUserAttributeAssignmentClient{
		Client: client,
	}, nil
}
