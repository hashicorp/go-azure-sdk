package outlooktaskgrouptaskfolder

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

type CreateOutlookTaskGroupTaskFolderPermanentDeleteOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type CreateOutlookTaskGroupTaskFolderPermanentDeleteOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultCreateOutlookTaskGroupTaskFolderPermanentDeleteOperationOptions() CreateOutlookTaskGroupTaskFolderPermanentDeleteOperationOptions {
	return CreateOutlookTaskGroupTaskFolderPermanentDeleteOperationOptions{}
}

func (o CreateOutlookTaskGroupTaskFolderPermanentDeleteOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o CreateOutlookTaskGroupTaskFolderPermanentDeleteOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o CreateOutlookTaskGroupTaskFolderPermanentDeleteOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// CreateOutlookTaskGroupTaskFolderPermanentDelete - Invoke action permanentDelete. Permanently delete an outlook task
// folder and remove its items from the user's mailbox. For more information about item retention, see Configure Deleted
// Item retention and Recoverable Items quotas.
func (c OutlookTaskGroupTaskFolderClient) CreateOutlookTaskGroupTaskFolderPermanentDelete(ctx context.Context, id beta.UserIdOutlookTaskGroupIdTaskFolderId, options CreateOutlookTaskGroupTaskFolderPermanentDeleteOperationOptions) (result CreateOutlookTaskGroupTaskFolderPermanentDeleteOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/permanentDelete", id.ID()),
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
