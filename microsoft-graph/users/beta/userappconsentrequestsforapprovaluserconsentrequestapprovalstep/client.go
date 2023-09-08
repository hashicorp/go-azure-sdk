package userappconsentrequestsforapprovaluserconsentrequestapprovalstep

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient struct {
	Client *msgraph.Client
}

func NewUserAppConsentRequestsForApprovalUserConsentRequestApprovalStepClientWithBaseURI(api sdkEnv.Api) (*UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userappconsentrequestsforapprovaluserconsentrequestapprovalstep", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient: %+v", err)
	}

	return &UserAppConsentRequestsForApprovalUserConsentRequestApprovalStepClient{
		Client: client,
	}, nil
}
