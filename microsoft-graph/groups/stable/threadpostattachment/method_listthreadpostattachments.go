package threadpostattachment

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/stable"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ListThreadPostAttachmentsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.Attachment
}

type ListThreadPostAttachmentsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.Attachment
}

type ListThreadPostAttachmentsCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListThreadPostAttachmentsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListThreadPostAttachments ...
func (c ThreadPostAttachmentClient) ListThreadPostAttachments(ctx context.Context, id GroupIdThreadIdPostId) (result ListThreadPostAttachmentsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListThreadPostAttachmentsCustomPager{},
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
		Values *[]stable.Attachment `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListThreadPostAttachmentsComplete retrieves all the results into a single object
func (c ThreadPostAttachmentClient) ListThreadPostAttachmentsComplete(ctx context.Context, id GroupIdThreadIdPostId) (ListThreadPostAttachmentsCompleteResult, error) {
	return c.ListThreadPostAttachmentsCompleteMatchingPredicate(ctx, id, AttachmentOperationPredicate{})
}

// ListThreadPostAttachmentsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c ThreadPostAttachmentClient) ListThreadPostAttachmentsCompleteMatchingPredicate(ctx context.Context, id GroupIdThreadIdPostId, predicate AttachmentOperationPredicate) (result ListThreadPostAttachmentsCompleteResult, err error) {
	items := make([]stable.Attachment, 0)

	resp, err := c.ListThreadPostAttachments(ctx, id)
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

	result = ListThreadPostAttachmentsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
