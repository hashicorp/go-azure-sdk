package outlooktaskfoldertaskattachment

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

type ListOutlookTaskFolderTaskAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]beta.Attachment
}

type ListOutlookTaskFolderTaskAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []beta.Attachment
}

type ListOutlookTaskFolderTaskAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListOutlookTaskFolderTaskAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListOutlookTaskFolderTaskAttachments ...
func (c OutlookTaskFolderTaskAttachmentClient) ListOutlookTaskFolderTaskAttachments(ctx context.Context, id UserIdOutlookTaskFolderIdTaskId) (result ListOutlookTaskFolderTaskAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListOutlookTaskFolderTaskAttachmentsCustomPager{},
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
		Values *[]beta.Attachment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListOutlookTaskFolderTaskAttachmentsComplete retrieves all the results into a single object
func (c OutlookTaskFolderTaskAttachmentClient) ListOutlookTaskFolderTaskAttachmentsComplete(ctx context.Context, id UserIdOutlookTaskFolderIdTaskId) (ListOutlookTaskFolderTaskAttachmentsCompleteResult, error) {
	return c.ListOutlookTaskFolderTaskAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListOutlookTaskFolderTaskAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c OutlookTaskFolderTaskAttachmentClient) ListOutlookTaskFolderTaskAttachmentsCompleteMatchingPredicate(ctx context.Context, id UserIdOutlookTaskFolderIdTaskId, predicate AttachmentOperationPredicate) (result ListOutlookTaskFolderTaskAttachmentsCompleteResult, err error) {
	items := make([]beta.Attachment, 0)

	resp, err := c.ListOutlookTaskFolderTaskAttachments(ctx, id)
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

	result = ListOutlookTaskFolderTaskAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
