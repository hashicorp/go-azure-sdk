package meapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeApprovalStepClient struct {
	Client *msgraph.Client
}

func NewMeApprovalStepClientWithBaseURI(api sdkEnv.Api) (*MeApprovalStepClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeApprovalStepClient: %+v", err)
	}

	return &MeApprovalStepClient{
		Client: client,
	}, nil
}
