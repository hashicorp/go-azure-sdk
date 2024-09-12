package conversationthread

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

type CreateConversationThreadOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *beta.ConversationThread
}

// CreateConversationThread - Create thread. Create a new thread in the specified conversation. A thread and post are
// created as specified. Use reply thread to further post to that thread. Or, if you get the post ID, you can also reply
// to that post in that thread. Note: You can also start a new conversation by first creating a thread.
func (c ConversationThreadClient) CreateConversationThread(ctx context.Context, id beta.GroupIdConversationId, input beta.ConversationThread) (result CreateConversationThreadOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusCreated,
		},
		HttpMethod: http.MethodPost,
		Path:       fmt.Sprintf("%s/threads", id.ID()),
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

	var model beta.ConversationThread
	result.Model = &model
	if err = resp.Unmarshal(result.Model); err != nil {
		return
	}

	return
}
