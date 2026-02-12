package iotcentrals

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExportsListDestinationsOperationResponse struct {
	HttpResponse *http.Response
	OData        *odata.OData
	Model        *[]Destination
}

type ExportsListDestinationsCompleteResult struct {
	LatestHttpResponse *http.Response
	Items              []Destination
}

type ExportsListDestinationsCustomPager struct {
	NextLink *odata.Link `json:"nextLink"`
}

func (p *ExportsListDestinationsCustomPager) NextPageLink() *odata.Link {
	defer func() {
		p.NextLink = nil
	}()

	return p.NextLink
}

// ExportsListDestinations ...
func (c IotcentralsClient) ExportsListDestinations(ctx context.Context, id ExportIdId) (result ExportsListDestinationsOperationResponse, err error) {
	opts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Pager:      &ExportsListDestinationsCustomPager{},
		Path:       fmt.Sprintf("%s/destinations", id.Path()),
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
		Values *[]json.RawMessage `json:"value"`
	}
	if err = resp.Unmarshal(&values); err != nil {
		return
	}

	temp := make([]Destination, 0)
	if values.Values != nil {
		for i, v := range *values.Values {
			val, err := UnmarshalDestinationImplementation(v)
			if err != nil {
				err = fmt.Errorf("unmarshalling item %d for Destination (%q): %+v", i, v, err)
				return result, err
			}
			temp = append(temp, val)
		}
	}
	result.Model = &temp

	return
}

// ExportsListDestinationsComplete retrieves all the results into a single object
func (c IotcentralsClient) ExportsListDestinationsComplete(ctx context.Context, id ExportIdId) (ExportsListDestinationsCompleteResult, error) {
	return c.ExportsListDestinationsCompleteMatchingPredicate(ctx, id, DestinationOperationPredicate{})
}

// ExportsListDestinationsCompleteMatchingPredicate retrieves all the results and then applies the predicate
func (c IotcentralsClient) ExportsListDestinationsCompleteMatchingPredicate(ctx context.Context, id ExportIdId, predicate DestinationOperationPredicate) (result ExportsListDestinationsCompleteResult, err error) {
	items := make([]Destination, 0)

	resp, err := c.ExportsListDestinations(ctx, id)
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

	result = ExportsListDestinationsCompleteResult{
		LatestHttpResponse: resp.HttpResponse,
		Items:              items,
	}
	return
}
