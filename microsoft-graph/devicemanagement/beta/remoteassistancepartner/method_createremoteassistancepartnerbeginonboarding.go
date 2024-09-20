package remoteassistancepartner

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateRemoteAssistancePartnerBeginOnboardingOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateRemoteAssistancePartnerBeginOnboardingOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateRemoteAssistancePartnerBeginOnboardingOperationOptions() CreateRemoteAssistancePartnerBeginOnboardingOperationOptions {
	return CreateRemoteAssistancePartnerBeginOnboardingOperationOptions{}
}

func (o CreateRemoteAssistancePartnerBeginOnboardingOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateRemoteAssistancePartnerBeginOnboardingOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateRemoteAssistancePartnerBeginOnboardingOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateRemoteAssistancePartnerBeginOnboarding - Invoke action beginOnboarding. A request to start onboarding. Must be
// coupled with the appropriate TeamViewer account information
func (c RemoteAssistancePartnerClient) CreateRemoteAssistancePartnerBeginOnboarding(ctx context.Context, id beta.DeviceManagementRemoteAssistancePartnerId, options CreateRemoteAssistancePartnerBeginOnboardingOperationOptions) (result CreateRemoteAssistancePartnerBeginOnboardingOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/beginOnboarding", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.Execute(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	return
}
