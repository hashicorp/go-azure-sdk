package meappconsentrequestsforapprovaluserconsentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeAppConsentRequestsForApprovalUserConsentRequestClient struct {
	Client *msgraph.Client
}

func NewMeAppConsentRequestsForApprovalUserConsentRequestClientWithBaseURI(api sdkEnv.Api) (*MeAppConsentRequestsForApprovalUserConsentRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "meappconsentrequestsforapprovaluserconsentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating MeAppConsentRequestsForApprovalUserConsentRequestClient: %+v", err)
	}

	return &MeAppConsentRequestsForApprovalUserConsentRequestClient{
		Client: client,
	}, nil
}
