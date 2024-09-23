package appconsentrequestsforapprovaluserconsentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppConsentRequestsForApprovalUserConsentRequestClient struct {
	Client *msgraph.Client
}

func NewAppConsentRequestsForApprovalUserConsentRequestClientWithBaseURI(sdkApi sdkEnv.Api) (*AppConsentRequestsForApprovalUserConsentRequestClient, error) {
	client, err := msgraph.NewClient(sdkApi, "appconsentrequestsforapprovaluserconsentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AppConsentRequestsForApprovalUserConsentRequestClient: %+v", err)
	}

	return &AppConsentRequestsForApprovalUserConsentRequestClient{
		Client: client,
	}, nil
}
