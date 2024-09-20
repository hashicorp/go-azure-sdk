package pendingaccessreviewinstance

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

type ResetPendingAccessReviewInstanceDecisionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type ResetPendingAccessReviewInstanceDecisionsOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultResetPendingAccessReviewInstanceDecisionsOperationOptions() ResetPendingAccessReviewInstanceDecisionsOperationOptions {
	return ResetPendingAccessReviewInstanceDecisionsOperationOptions{}
}

func (o ResetPendingAccessReviewInstanceDecisionsOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o ResetPendingAccessReviewInstanceDecisionsOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o ResetPendingAccessReviewInstanceDecisionsOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// ResetPendingAccessReviewInstanceDecisions - Invoke action resetDecisions. Resets decisions of all
// accessReviewInstanceDecisionItem objects on an accessReviewInstance to notReviewed.
func (c PendingAccessReviewInstanceClient) ResetPendingAccessReviewInstanceDecisions(ctx context.Context, id beta.MePendingAccessReviewInstanceId, options ResetPendingAccessReviewInstanceDecisionsOperationOptions) (result ResetPendingAccessReviewInstanceDecisionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/resetDecisions", id.ID()),
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
