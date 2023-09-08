package meappconsentrequestsforapprovaluserconsentrequestapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient struct {
	Client *msgraph.Client
}

func NewMeAppConsentRequestsForApprovalUserConsentRequestApprovalStepClientWithBaseURI(api sdkEnv.Api) (*MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meappconsentrequestsforapprovaluserconsentrequestapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient: %+v", err)
	}

	return &MeAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient{
		Client: client,
	}, nil
}
