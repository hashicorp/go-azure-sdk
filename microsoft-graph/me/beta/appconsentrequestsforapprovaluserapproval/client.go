package appconsentrequestsforapprovaluserapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequestsForApprovalUserApprovalClient struct {
	Client *msgraph.Client
}

func NewAppConsentRequestsForApprovalUserApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentRequestsForApprovalUserApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentrequestsforapprovaluserapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentRequestsForApprovalUserApprovalClient: %+v", err)
	}

	return &AppConsentRequestsForApprovalUserApprovalClient{
		Client: client,
	}, nil
}
