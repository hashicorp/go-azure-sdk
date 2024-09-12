package operationapprovalrequest

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

type CancelOperationApprovalRequestApprovalOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *CancelOperationApprovalRequestApprovalResult
}

// CancelOperationApprovalRequestApproval - Invoke action cancelApproval. Cancels an already approved instance of an
// operationApprovalRequest.
func (c OperationApprovalRequestClient) CancelOperationApprovalRequestApproval(ctx context.Context, id beta.DeviceManagementOperationApprovalRequestId, input CancelOperationApprovalRequestApprovalRequest) (result CancelOperationApprovalRequestApprovalOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/cancelApproval", id.ID()),
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

	var model CancelOperationApprovalRequestApprovalResult
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
