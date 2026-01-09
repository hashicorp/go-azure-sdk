package mailfolderchildfolder

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateMailFolderChildFolderAllMessagesReadStateOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
}

type UpdateMailFolderChildFolderAllMessagesReadStateOperationOptions struct {
	Metadata  *odata.Metadata
	RetryFunc client.RequestRetryFunc
}

func DefaultUpdateMailFolderChildFolderAllMessagesReadStateOperationOptions() UpdateMailFolderChildFolderAllMessagesReadStateOperationOptions {
	return UpdateMailFolderChildFolderAllMessagesReadStateOperationOptions{}
}

func (o UpdateMailFolderChildFolderAllMessagesReadStateOperationOptions) ToHeaders() *client.Headers {
	out := client.Headers{}

	return &out
}

func (o UpdateMailFolderChildFolderAllMessagesReadStateOperationOptions) ToOData() *odata.Query {
	out := odata.Query{}
	if o.Metadata != nil {
		out.Metadata = *o.Metadata
	}
	return &out
}

func (o UpdateMailFolderChildFolderAllMessagesReadStateOperationOptions) ToQuery() *client.QueryParams {
	out := client.QueryParams{}

	return &out
}

// UpdateMailFolderChildFolderAllMessagesReadState - Invoke action updateAllMessagesReadState. Update the read state of
// all messages in a mailFolder object.
func (c MailFolderChildFolderClient) UpdateMailFolderChildFolderAllMessagesReadState(ctx context.Context, id beta.UserIdMailFolderIdChildFolderId, input UpdateMailFolderChildFolderAllMessagesReadStateRequest, options UpdateMailFolderChildFolderAllMessagesReadStateOperationOptions) (result UpdateMailFolderChildFolderAllMessagesReadStateOperationResponse, err error) {
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
		Path:          fmt.Sprintf("%s/updateAllMessagesReadState", id.ID()),
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

	return
}
