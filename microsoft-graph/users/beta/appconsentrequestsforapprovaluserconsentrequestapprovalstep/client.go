package appconsentrequestsforapprovaluserconsentrequestapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequestsForApprovalUserConsentRequestApprovalStepClient struct {
	Client *msgraph.Client
}

func NewAppConsentRequestsForApprovalUserConsentRequestApprovalStepClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentRequestsForApprovalUserConsentRequestApprovalStepClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentrequestsforapprovaluserconsentrequestapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentRequestsForApprovalUserConsentRequestApprovalStepClient: %+v", err)
	}

	return &AppConsentRequestsForApprovalUserConsentRequestApprovalStepClient{
		Client: client,
	}, nil
}
