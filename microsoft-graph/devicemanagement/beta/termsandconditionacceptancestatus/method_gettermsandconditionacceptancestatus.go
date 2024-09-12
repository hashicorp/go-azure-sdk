package termsandconditionacceptancestatus

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetTermsAndConditionAcceptanceStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TermsAndConditionsAcceptanceStatus
}

type GetTermsAndConditionAcceptanceStatusOperationOptions struct {
	Expand *odata.Expand
	Select *[]string
}

func DefaultGetTermsAndConditionAcceptanceStatusOperationOptions() GetTermsAndConditionAcceptanceStatusOperationOptions {
	return GetTermsAndConditionAcceptanceStatusOperationOptions{}
}

func (o GetTermsAndConditionAcceptanceStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o GetTermsAndConditionAcceptanceStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Expand != nil {
		out.Expand = *o.Expand
	}
	if o.Select != nil {
		out.Select = *o.Select
	}
	return &out
}

func (o GetTermsAndConditionAcceptanceStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// GetTermsAndConditionAcceptanceStatus - Get acceptanceStatuses from deviceManagement. The list of acceptance statuses
// for this T&C policy.
func (c TermsAndConditionAcceptanceStatusClient) GetTermsAndConditionAcceptanceStatus(ctx context.Context, id beta.DeviceManagementTermsAndConditionIdAcceptanceStatusId, options GetTermsAndConditionAcceptanceStatusOperationOptions) (result GetTermsAndConditionAcceptanceStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: options,
		Path:          id.ID(),
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

	var model beta.TermsAndConditionsAcceptanceStatus
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
