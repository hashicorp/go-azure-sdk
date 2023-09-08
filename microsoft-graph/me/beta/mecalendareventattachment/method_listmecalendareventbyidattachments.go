package mecalendareventattachment

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

type ListMeCalendarEventByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListMeCalendarEventByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListMeCalendarEventByIdAttachments ...
func (c MeCalendarEventAttachmentClient) ListMeCalendarEventByIdAttachments(ctx context.Context, id MeCalendarEventId) (result ListMeCalendarEventByIdAttachmentsOperationResponse, err error) {
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

// ListMeCalendarEventByIdAttachmentsComplete retrieves all the results into a single object
func (c MeCalendarEventAttachmentClient) ListMeCalendarEventByIdAttachmentsComplete(ctx context.Context, id MeCalendarEventId) (ListMeCalendarEventByIdAttachmentsCompleteResult, error) {
	return c.ListMeCalendarEventByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListMeCalendarEventByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarEventAttachmentClient) ListMeCalendarEventByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarEventId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListMeCalendarEventByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListMeCalendarEventByIdAttachments(ctx, id)
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

	result = ListMeCalendarEventByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
