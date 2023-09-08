package userappconsentrequestsforapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAppConsentRequestsForApprovalClient struct {
	Client *msgraph.Client
}

func NewUserAppConsentRequestsForApprovalClientWithBaseURI(api sdkEnv.Api) (*UserAppConsentRequestsForApprovalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userappconsentrequestsforapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAppConsentRequestsForApprovalClient: %+v", err)
	}

	return &UserAppConsentRequestsForApprovalClient{
		Client: client,
	}, nil
}
