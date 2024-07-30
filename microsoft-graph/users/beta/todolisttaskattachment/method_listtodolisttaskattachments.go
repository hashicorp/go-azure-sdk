package todolisttaskattachment

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

type ListTodoListTaskAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.AttachmentBase
}

type ListTodoListTaskAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.AttachmentBase
}

type ListTodoListTaskAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListTodoListTaskAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListTodoListTaskAttachments ...
func (c TodoListTaskAttachmentClient) ListTodoListTaskAttachments(ctx context.Context, id UserIdTodoListIdTaskId) (result ListTodoListTaskAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListTodoListTaskAttachmentsCustomPager{},
		Path:       fmt.Sprintf("%s/attachments", id.ID()),
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
		Values *[]beta.AttachmentBase `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListTodoListTaskAttachmentsComplete retrieves all the results into a single object
func (c TodoListTaskAttachmentClient) ListTodoListTaskAttachmentsComplete(ctx context.Context, id UserIdTodoListIdTaskId) (ListTodoListTaskAttachmentsCompleteResult, error) {
	return c.ListTodoListTaskAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentBaseOperationPredicate{})
}

// ListTodoListTaskAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c TodoListTaskAttachmentClient) ListTodoListTaskAttachmentsCompleteMatchingPredicate(ctx context.Context, id UserIdTodoListIdTaskId, predicate AttachmentBaseOperationPredicate) (result ListTodoListTaskAttachmentsCompleteResult, err error) {
	items := make([]beta.AttachmentBase, 0)

	resp, err := c.ListTodoListTaskAttachments(ctx, id)
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

	result = ListTodoListTaskAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
