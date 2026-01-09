package appconsentrequestsforapprovaluserconsentrequestapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequestsForApprovalUserConsentRequestApprovalClient struct {
	Client *msgraph.Client
}

func NewAppConsentRequestsForApprovalUserConsentRequestApprovalClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentRequestsForApprovalUserConsentRequestApprovalClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentrequestsforapprovaluserconsentrequestapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentRequestsForApprovalUserConsentRequestApprovalClient: %+v", err)
	}

	return &AppConsentRequestsForApprovalUserConsentRequestApprovalClient{
		Client: client,
	}, nil
}
