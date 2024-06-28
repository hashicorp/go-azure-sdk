package ueinformationlist

import (
	"context"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UeInformationListOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]UeInfo
}

type UeInformationListCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []UeInfo
}

type UeInformationListCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *UeInformationListCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// UeInformationList ...
func (c UeInformationListClient) UeInformationList(ctx context.Context, id PacketCoreControlPlaneId) (result UeInformationListOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &UeInformationListCustomPager{},
		Path:       fmt.Sprintf("%s/ues", id.ID()),
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
		Values *[]UeInfo `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	result.Model = values.Values

	return
}

// UeInformationListComplete retrieves all the results into a single object
func (c UeInformationListClient) UeInformationListComplete(ctx context.Context, id PacketCoreControlPlaneId) (UeInformationListCompleteResult, error) {
	return c.UeInformationListCompleteMatchingPredicate(ctx, id, UeInfoOperationPredicate{})
}

// UeInformationListCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c UeInformationListClient) UeInformationListCompleteMatchingPredicate(ctx context.Context, id PacketCoreControlPlaneId, predicate UeInfoOperationPredicate) (result UeInformationListCompleteResult, err error) {
	items := make([]UeInfo, 0)

	resp, err := c.UeInformationList(ctx, id)
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

	result = UeInformationListCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
