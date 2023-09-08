package userinformationprotectionthreatassessmentrequest

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/msgraph"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserInformationProtectionThreatAssessmentRequestClient struct {
	Client *msgraph.Client
}

func NewUserInformationProtectionThreatAssessmentRequestClientWithBaseURI(api sdkEnv.Api) (*UserInformationProtectionThreatAssessmentRequestClient, error) {
	client, err := msgraph.NewMsGraphClient(api, "userinformationprotectionthreatassessmentrequest", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating UserInformationProtectionThreatAssessmentRequestClient: %+v", err)
	}

	return &UserInformationProtectionThreatAssessmentRequestClient{
		Client: client,
	}, nil
}
