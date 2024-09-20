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

type SendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type SendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationOptions struct {
	Metadata *odata.Metadata
}

func DefaultSendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationOptions() SendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationOptions {
	return SendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationOptions{}
}

func (o SendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o SendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o SendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// SendPendingAccessReviewInstanceStageDecisionInstanceReminder - Invoke action sendReminder. Send a reminder to the
// reviewers of a currently active accessReviewInstance.
func (c PendingAccessReviewInstanceStageDecisionInstanceClient) SendPendingAccessReviewInstanceStageDecisionInstanceReminder(ctx context.Context, id beta.MePendingAccessReviewInstanceIdStageIdDecisionId, options SendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationOptions) (result SendPendingAccessReviewInstanceStageDecisionInstanceReminderOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodPost,
		OptionsObject: options,
		Path:          fmt.Sprintf("%s/instance/sendReminder", id.ID()),
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
