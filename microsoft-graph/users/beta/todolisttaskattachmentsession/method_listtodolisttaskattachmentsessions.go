package todolisttaskattachmentsession

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

type ListTodoListTaskAttachmentSessionsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AttachmentSession
}

type ListTodoListTaskAttachmentSessionsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AttachmentSession
}

type ListTodoListTaskAttachmentSessionsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTodoListTaskAttachmentSessionsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTodoListTaskAttachmentSessions ...
func (c TodoListTaskAttachmentSessionClient) ListTodoListTaskAttachmentSessions(ctx context.Context, id UserIdTodoListIdTaskId) (result ListTodoListTaskAttachmentSessionsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTodoListTaskAttachmentSessionsCustomPager{},
		Path:       fmt.Sprintf("%s/attachmentSessions", id.ID()),
	}

	req, err := c.Client.NewRequest(ctx, opts)
	if err != nil {
		return
	}

	var resp *client.Response
	resp, err = req.ExecutePaged(ctx)
	if resp != nil {
		result.OData = resp.OData
		result.HttpResponse = resp.Response
	}
	if err != nil {
		return
	}

	var values struct {
		Values *[]beta.AttachmentSession `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTodoListTaskAttachmentSessionsComplete retrieves all the results into a single object
func (c TodoListTaskAttachmentSessionClient) ListTodoListTaskAttachmentSessionsComplete(ctx context.Context, id UserIdTodoListIdTaskId) (ListTodoListTaskAttachmentSessionsCompleteResult, error) {
	return c.ListTodoListTaskAttachmentSessionsCompleteMatchingPredicate(ctx, id, AttachmentSessionOperationPredicate{})
}

// ListTodoListTaskAttachmentSessionsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TodoListTaskAttachmentSessionClient) ListTodoListTaskAttachmentSessionsCompleteMatchingPredicate(ctx context.Context, id UserIdTodoListIdTaskId, predicate AttachmentSessionOperationPredicate) (result ListTodoListTaskAttachmentSessionsCompleteResult, err error) {
	items := make([]beta.AttachmentSession, 0)

	resp, err := c.ListTodoListTaskAttachmentSessions(ctx, id)
	if err != nil {
		result.LatestHttpResponse = resp.HttpResponse
		err = fmt.Errorf("loading results: %+v", err)
		return
	}
	if resp.Model != nil {
		for _, v := range *resp.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	result = ListTodoListTaskAttachmentSessionsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
