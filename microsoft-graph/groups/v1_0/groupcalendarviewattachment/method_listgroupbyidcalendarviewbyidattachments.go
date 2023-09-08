package groupcalendarviewattachment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/v1_0/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListGroupByIdCalendarViewByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListGroupByIdCalendarViewByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListGroupByIdCalendarViewByIdAttachments ...
func (c GroupCalendarViewAttachmentClient) ListGroupByIdCalendarViewByIdAttachments(ctx context.Context, id GroupCalendarViewId) (result ListGroupByIdCalendarViewByIdAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
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
		Values *[]models.AttachmentCollectionResponse `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListGroupByIdCalendarViewByIdAttachmentsComplete retrieves all the results into a single object
func (c GroupCalendarViewAttachmentClient) ListGroupByIdCalendarViewByIdAttachmentsComplete(ctx context.Context, id GroupCalendarViewId) (ListGroupByIdCalendarViewByIdAttachmentsCompleteResult, error) {
	return c.ListGroupByIdCalendarViewByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListGroupByIdCalendarViewByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupCalendarViewAttachmentClient) ListGroupByIdCalendarViewByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id GroupCalendarViewId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListGroupByIdCalendarViewByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListGroupByIdCalendarViewByIdAttachments(ctx, id)
	if err != nil {
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

	result = ListGroupByIdCalendarViewByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
