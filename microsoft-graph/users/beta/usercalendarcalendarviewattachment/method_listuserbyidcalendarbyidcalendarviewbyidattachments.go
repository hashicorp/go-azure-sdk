package usercalendarcalendarviewattachment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common/beta/models"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListUserByIdCalendarByIdCalendarViewByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListUserByIdCalendarByIdCalendarViewByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListUserByIdCalendarByIdCalendarViewByIdAttachments ...
func (c UserCalendarCalendarViewAttachmentClient) ListUserByIdCalendarByIdCalendarViewByIdAttachments(ctx context.Context, id UserCalendarCalendarViewId) (result ListUserByIdCalendarByIdCalendarViewByIdAttachmentsOperationResponse, err error) {
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

// ListUserByIdCalendarByIdCalendarViewByIdAttachmentsComplete retrieves all the results into a single object
func (c UserCalendarCalendarViewAttachmentClient) ListUserByIdCalendarByIdCalendarViewByIdAttachmentsComplete(ctx context.Context, id UserCalendarCalendarViewId) (ListUserByIdCalendarByIdCalendarViewByIdAttachmentsCompleteResult, error) {
	return c.ListUserByIdCalendarByIdCalendarViewByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListUserByIdCalendarByIdCalendarViewByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UserCalendarCalendarViewAttachmentClient) ListUserByIdCalendarByIdCalendarViewByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id UserCalendarCalendarViewId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListUserByIdCalendarByIdCalendarViewByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListUserByIdCalendarByIdCalendarViewByIdAttachments(ctx, id)
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

	result = ListUserByIdCalendarByIdCalendarViewByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
