package privilegedaccessgroupassignmentapprovalstage

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreatePrivilegedAccessGroupAssignmentApprovalStageOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.ApprovalStage
}

type CreatePrivilegedAccessGroupAssignmentApprovalStageOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreatePrivilegedAccessGroupAssignmentApprovalStageOperationOptions() CreatePrivilegedAccessGroupAssignmentApprovalStageOperationOptions {
	return CreatePrivilegedAccessGroupAssignmentApprovalStageOperationOptions{}
}

func (o CreatePrivilegedAccessGroupAssignmentApprovalStageOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreatePrivilegedAccessGroupAssignmentApprovalStageOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreatePrivilegedAccessGroupAssignmentApprovalStageOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreatePrivilegedAccessGroupAssignmentApprovalStage - Create new navigation property to stages for identityGovernance
func (c PrivilegedAccessGroupAssignmentApprovalStageClient) CreatePrivilegedAccessGroupAssignmentApprovalStage(ctx context.Context, id stable.IdentityGovernancePrivilegedAccessGroupAssignmentApprovalId, input stable.ApprovalStage, options CreatePrivilegedAccessGroupAssignmentApprovalStageOperationOptions) (result CreatePrivilegedAccessGroupAssignmentApprovalStageOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/stages", id.ID()),
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

	var model stable.ApprovalStage
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
