package meappconsentrequestsforapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAppConsentRequestsForApprovalClient struct {
	Client *msgraph.Client
}

func NewMeAppConsentRequestsForApprovalClientWithBaseURI(api sdkEnv.Api) (*MeAppConsentRequestsForApprovalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meappconsentrequestsforapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAppConsentRequestsForApprovalClient: %+v", err)
	}

	return &MeAppConsentRequestsForApprovalClient{
		Client: client,
	}, nil
}
