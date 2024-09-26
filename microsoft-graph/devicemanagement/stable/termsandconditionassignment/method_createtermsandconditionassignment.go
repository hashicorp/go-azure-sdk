package termsandconditionassignment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateTermsAndConditionAssignmentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.TermsAndConditionsAssignment
}

type CreateTermsAndConditionAssignmentOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateTermsAndConditionAssignmentOperationOptions() CreateTermsAndConditionAssignmentOperationOptions {
	return CreateTermsAndConditionAssignmentOperationOptions{}
}

func (o CreateTermsAndConditionAssignmentOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateTermsAndConditionAssignmentOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateTermsAndConditionAssignmentOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateTermsAndConditionAssignment - Create termsAndConditionsAssignment. Create a new termsAndConditionsAssignment
// object.
func (c TermsAndConditionAssignmentClient) CreateTermsAndConditionAssignment(ctx context.Context, id stable.DeviceManagementTermsAndConditionId, input stable.TermsAndConditionsAssignment, options CreateTermsAndConditionAssignmentOperationOptions) (result CreateTermsAndConditionAssignmentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusAccepted,
			http.StatusCreated,
			http.StatusNoContent,
			http.StatusOK,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/assignments", id.ID()),
		RetryFunc:     options.RetryFunc,
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

	var model stable.TermsAndConditionsAssignment
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
