package pendingaccessreviewinstancestagedecisioninstance

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

type StopPendingAccessReviewInstanceStageDecisionInstanceOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type StopPendingAccessReviewInstanceStageDecisionInstanceOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultStopPendingAccessReviewInstanceStageDecisionInstanceOperationOptions() StopPendingAccessReviewInstanceStageDecisionInstanceOperationOptions {
	return StopPendingAccessReviewInstanceStageDecisionInstanceOperationOptions{}
}

func (o StopPendingAccessReviewInstanceStageDecisionInstanceOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o StopPendingAccessReviewInstanceStageDecisionInstanceOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o StopPendingAccessReviewInstanceStageDecisionInstanceOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// StopPendingAccessReviewInstanceStageDecisionInstance - Invoke action stop. Stop a currently active
// accessReviewInstance. After the access review instance stops, the instance status will be Completed, the reviewers
// can no longer give input, and the access review decisions can be applied. Stopping an instance will not effect future
// instances. To prevent a recurring access review from starting future instances, update the schedule definition to
// change its scheduled end date.
func (c PendingAccessReviewInstanceStageDecisionInstanceClient) StopPendingAccessReviewInstanceStageDecisionInstance(ctx context.Context, id beta.UserIdPendingAccessReviewInstanceIdStageIdDecisionId, options StopPendingAccessReviewInstanceStageDecisionInstanceOperationOptions) (result StopPendingAccessReviewInstanceStageDecisionInstanceOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/instance/stop", id.ID()),
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
