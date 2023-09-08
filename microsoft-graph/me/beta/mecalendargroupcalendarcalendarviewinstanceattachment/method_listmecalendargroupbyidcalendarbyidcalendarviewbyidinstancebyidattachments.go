package mecalendargroupcalendarcalendarviewinstanceattachment

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

type ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]models.AttachmentCollectionResponse
}

type ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult struct {
	Items []models.AttachmentCollectionResponse
}

// ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachments ...
func (c MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient) ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachments(ctx context.Context, id MeCalendarGroupCalendarCalendarViewInstanceId) (result ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsOperationResponse, err error) {
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

// ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsComplete retrieves all the results into a single object
func (c MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient) ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsComplete(ctx context.Context, id MeCalendarGroupCalendarCalendarViewInstanceId) (ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult, error) {
	return c.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsCompleteMatchingPredicate(ctx, id, models.AttachmentCollectionResponseOperationPredicate{})
}

// ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c MeCalendarGroupCalendarCalendarViewInstanceAttachmentClient) ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsCompleteMatchingPredicate(ctx context.Context, id MeCalendarGroupCalendarCalendarViewInstanceId, predicate models.AttachmentCollectionResponseOperationPredicate) (result ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult, err error) {
	items := make([]models.AttachmentCollectionResponse, 0)

	resp, err := c.ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachments(ctx, id)
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

	result = ListMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdAttachmentsCompleteResult{
		Items: items,
	}
	return
}
