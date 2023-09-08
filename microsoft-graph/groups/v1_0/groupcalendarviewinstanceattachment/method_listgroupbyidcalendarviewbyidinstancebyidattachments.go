package groupcalendarviewinstanceattachment

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

type ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListGroupByIdCalendarViewByIdInstanceByIdAttachments ...
func (c GroupCalendarViewInstanceAttachmentClient) ListGroupByIdCalendarViewByIdInstanceByIdAttachments(ctx context.Context, id GroupCalendarViewInstanceId) (result ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsOperationResponse, err error) {
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

// ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsComplete retrieves all the results into a single object
func (c GroupCalendarViewInstanceAttachmentClient) ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsComplete(ctx context.Context, id GroupCalendarViewInstanceId) (ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult, error) {
	return c.ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c GroupCalendarViewInstanceAttachmentClient) ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id GroupCalendarViewInstanceId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListGroupByIdCalendarViewByIdInstanceByIdAttachments(ctx, id)
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

	result = ListGroupByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
