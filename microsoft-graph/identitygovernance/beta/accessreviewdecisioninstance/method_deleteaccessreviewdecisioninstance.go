package accessreviewdecisioninstance

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

type DeleteAccessReviewDecisionInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type DeleteAccessReviewDecisionInstanceOperationOptions struct {
	IfMatch   *string
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultDeleteAccessReviewDecisionInstanceOperationOptions() DeleteAccessReviewDecisionInstanceOperationOptions {
	return DeleteAccessReviewDecisionInstanceOperationOptions{}
}

func (o DeleteAccessReviewDecisionInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}
	if o.IfMatch != nil {
		out.Append("If-Match", fmt.Sprintf("%v", *o.IfMatch))
	}
	return &out
}

func (o DeleteAccessReviewDecisionInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o DeleteAccessReviewDecisionInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// DeleteAccessReviewDecisionInstance - Delete navigation property instance for identityGovernance
func (c AccessReviewDecisionInstanceClient) DeleteAccessReviewDecisionInstance(ctx context.Context, id beta.IdentityGovernanceAccessReviewDecisionId, options DeleteAccessReviewDecisionInstanceOperationOptions) (result DeleteAccessReviewDecisionInstanceOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodDelete,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance", id.ID()),
		RetryFunc:     options.RetryFunc,
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
