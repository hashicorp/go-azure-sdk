package todolisttaskattachmentsessioncontent

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

type CreateUpdateTodoListTaskAttachmentSessionContentOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *stable.AttachmentSession
}

// CreateUpdateTodoListTaskAttachmentSessionContent ...
func (c TodoListTaskAttachmentSessionContentClient) CreateUpdateTodoListTaskAttachmentSessionContent(ctx context.Context, id UserIdTodoListIdTaskIdAttachmentSessionId, input []byte) (result CreateUpdateTodoListTaskAttachmentSessionContentOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodPut,
		Path:       fmt.Sprintf("%s/content", id.ID()),
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

	var model stable.AttachmentSession
	result.Model = &model

	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
