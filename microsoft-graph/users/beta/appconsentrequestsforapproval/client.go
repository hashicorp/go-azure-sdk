package appconsentrequestsforapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequestsForApprovalClient struct {
	Client *msgraph.Client
}

func NewAppConsentRequestsForApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentRequestsForApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentrequestsforapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentRequestsForApprovalClient: %+v", err)
	}

	return &AppConsentRequestsForApprovalClient{
		Client: client,
	}, nil
}
