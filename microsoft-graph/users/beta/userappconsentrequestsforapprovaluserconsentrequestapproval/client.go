package userappconsentrequestsforapprovaluserconsentrequestapproval

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAppConsentRequestsForApprovalUserConsentRequestApprovalClient struct {
	Client *msgraph.Client
}

func NewUserAppConsentRequestsForApprovalUserConsentRequestApprovalClientWithBaseURI(api sdkEnv.Api) (*UserAppConsentRequestsForApprovalUserConsentRequestApprovalClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userappconsentrequestsforapprovaluserconsentrequestapproval", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserAppConsentRequestsForApprovalUserConsentRequestApprovalClient: %+v", err)
	}

	return &UserAppConsentRequestsForApprovalUserConsentRequestApprovalClient{
		Client: client,
	}, nil
}
