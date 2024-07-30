package b2xuserflowuserattributeassignmentuserattribute

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xUserFlowUserAttributeAssignmentUserAttributeClient struct {
	Client *msgraph.Client
}

func NewB2xUserFlowUserAttributeAssignmentUserAttributeClientWithBaseURI(sdkApi sdkEnv.Api) (*B2xUserFlowUserAttributeAssignmentUserAttributeClient, error) {
	client, err := msgraph.NewClient(sdkApi, "b2xuserflowuserattributeassignmentuserattribute", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating B2xUserFlowUserAttributeAssignmentUserAttributeClient: %+v", err)
	}

	return &B2xUserFlowUserAttributeAssignmentUserAttributeClient{
		Client: client,
	}, nil
}
