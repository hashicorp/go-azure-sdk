package pendingaccessreviewinstancedecisioninstancecontactedreviewer

import (
	"context"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewerOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewerOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultUpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewerOperationOptions() UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewerOperationOptions {
	return UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewerOperationOptions{}
}

func (o UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewerOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewerOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewerOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewer - Update the navigation property
// contactedReviewers in me
func (c PendingAccessReviewInstanceDecisionInstanceContactedReviewerClient) UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewer(ctx context.Context, id beta.MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId, input beta.AccessReviewReviewer, options UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewerOperationOptions) (result UpdatePendingAccessReviewInstanceDecisionInstanceContactedReviewerOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPatch,
		OptionsObject: options,
		Path:          id.ID(),
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

	return
}
