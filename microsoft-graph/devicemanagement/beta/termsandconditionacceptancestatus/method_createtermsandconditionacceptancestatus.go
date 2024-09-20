package termsandconditionacceptancestatus

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

type CreateTermsAndConditionAcceptanceStatusOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.TermsAndConditionsAcceptanceStatus
}

type CreateTermsAndConditionAcceptanceStatusOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultCreateTermsAndConditionAcceptanceStatusOperationOptions() CreateTermsAndConditionAcceptanceStatusOperationOptions {
	return CreateTermsAndConditionAcceptanceStatusOperationOptions{}
}

func (o CreateTermsAndConditionAcceptanceStatusOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTermsAndConditionAcceptanceStatusOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTermsAndConditionAcceptanceStatusOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTermsAndConditionAcceptanceStatus - Create new navigation property to acceptanceStatuses for deviceManagement
func (c TermsAndConditionAcceptanceStatusClient) CreateTermsAndConditionAcceptanceStatus(ctx context.Context, id beta.DeviceManagementTermsAndConditionId, input beta.TermsAndConditionsAcceptanceStatus, options CreateTermsAndConditionAcceptanceStatusOperationOptions) (result CreateTermsAndConditionAcceptanceStatusOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/acceptanceStatuses", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	if err = req.Marshal(input); err != nil {
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
