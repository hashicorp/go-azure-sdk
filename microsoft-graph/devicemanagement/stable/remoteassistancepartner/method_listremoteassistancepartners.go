package remoteassistancepartner

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

type ListRemoteAssistancePartnersOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]stable.RemoteAssistancePartner
}

type ListRemoteAssistancePartnersCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []stable.RemoteAssistancePartner
}

type ListRemoteAssistancePartnersCustomPager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *ListRemoteAssistancePartnersCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ListRemoteAssistancePartners ...
func (c RemoteAssistancePartnerClient) ListRemoteAssistancePartners(ctx context.Context) (result ListRemoteAssistancePartnersOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ListRemoteAssistancePartnersCustomPager{},
		Path:       "/deviceManagement/remoteAssistancePartners",
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
		Values *[]stable.RemoteAssistancePartner `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// ListRemoteAssistancePartnersComplete retrieves all the results into a single object
func (c RemoteAssistancePartnerClient) ListRemoteAssistancePartnersComplete(ctx context.Context) (ListRemoteAssistancePartnersCompleteResult, error) {
	return c.ListRemoteAssistancePartnersCompleteMatchingPredicate(ctx, RemoteAssistancePartnerOperationPredicate{})
}

// ListRemoteAssistancePartnersCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c RemoteAssistancePartnerClient) ListRemoteAssistancePartnersCompleteMatchingPredicate(ctx context.Context, predicate RemoteAssistancePartnerOperationPredicate) (result ListRemoteAssistancePartnersCompleteResult, err error) {
	items := make([]stable.RemoteAssistancePartner, 0)

	resp, err := c.ListRemoteAssistancePartners(ctx)
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

	result = ListRemoteAssistancePartnersCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
