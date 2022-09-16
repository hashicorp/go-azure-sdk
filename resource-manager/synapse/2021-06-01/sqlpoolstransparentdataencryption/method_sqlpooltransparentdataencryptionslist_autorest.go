package sqlpoolstransparentdataencryption

import (
	"context"
	"fmt"
	"net/http"
	"net/url"

	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolTransparentDataEncryptionsListOperationResponse struct {
	HttpResponse *http.Response
	Model        *[]TransparentDataEncryption

	nextLink     *string
	nextPageFunc func(ctx context.Context, nextLink string) (SqlPoolTransparentDataEncryptionsListOperationResponse, error)
}

type SqlPoolTransparentDataEncryptionsListCompleteResult struct {
	Items []TransparentDataEncryption
}

func (r SqlPoolTransparentDataEncryptionsListOperationResponse) HasMore() bool {
	return r.nextLink != nil
}

func (r SqlPoolTransparentDataEncryptionsListOperationResponse) LoadMore(ctx context.Context) (resp SqlPoolTransparentDataEncryptionsListOperationResponse, err error) {
	if !r.HasMore() {
		err = fmt.Errorf("no more pages returned")
		return
	}
	return r.nextPageFunc(ctx, *r.nextLink)
}

// SqlPoolTransparentDataEncryptionsList ...
func (c SqlPoolsTransparentDataEncryptionClient) SqlPoolTransparentDataEncryptionsList(ctx context.Context, id SqlPoolId) (resp SqlPoolTransparentDataEncryptionsListOperationResponse, err error) {
	req, err := c.preparerForSqlPoolTransparentDataEncryptionsList(ctx, id)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsList", nil, "Failure preparing request")
		return
	}

	resp.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsList", resp.HttpResponse, "Failure sending request")
		return
	}

	resp, err = c.responderForSqlPoolTransparentDataEncryptionsList(resp.HttpResponse)
	if err != nil {
		err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsList", resp.HttpResponse, "Failure responding to request")
		return
	}
	return
}

// preparerForSqlPoolTransparentDataEncryptionsList prepares the SqlPoolTransparentDataEncryptionsList request.
func (c SqlPoolsTransparentDataEncryptionClient) preparerForSqlPoolTransparentDataEncryptionsList(ctx context.Context, id SqlPoolId) (*http.Request, error) {
	queryParameters := map[string]interface{}{
		"api-version": defaultApiVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(fmt.Sprintf("%s/transparentDataEncryption", id.ID())),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// preparerForSqlPoolTransparentDataEncryptionsListWithNextLink prepares the SqlPoolTransparentDataEncryptionsList request with the given nextLink token.
func (c SqlPoolsTransparentDataEncryptionClient) preparerForSqlPoolTransparentDataEncryptionsListWithNextLink(ctx context.Context, nextLink string) (*http.Request, error) {
	uri, err := url.Parse(nextLink)
	if err != nil {
		return nil, fmt.Errorf("parsing nextLink %q: %+v", nextLink, err)
	}
	queryParameters := map[string]interface{}{}
	for k, v := range uri.Query() {
		if len(v) == 0 {
			continue
		}
		val := v[0]
		val = autorest.Encode("query", val)
		queryParameters[k] = val
	}

	preparer := autorest.CreatePreparer(
		autorest.AsContentType("application/json; charset=utf-8"),
		autorest.AsGet(),
		autorest.WithBaseURL(c.baseUri),
		autorest.WithPath(uri.Path),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// responderForSqlPoolTransparentDataEncryptionsList handles the response to the SqlPoolTransparentDataEncryptionsList request. The method always
// closes the http.Response Body.
func (c SqlPoolsTransparentDataEncryptionClient) responderForSqlPoolTransparentDataEncryptionsList(resp *http.Response) (result SqlPoolTransparentDataEncryptionsListOperationResponse, err error) {
	type page struct {
		Values   []TransparentDataEncryption `json:"value"`
		NextLink *string                     `json:"nextLink"`
	}
	var respObj page
	err = autorest.Respond(
		resp,
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&respObj),
		autorest.ByClosing())
	result.HttpResponse = resp
	result.Model = &respObj.Values
	result.nextLink = respObj.NextLink
	if respObj.NextLink != nil {
		result.nextPageFunc = func(ctx context.Context, nextLink string) (result SqlPoolTransparentDataEncryptionsListOperationResponse, err error) {
			req, err := c.preparerForSqlPoolTransparentDataEncryptionsListWithNextLink(ctx, nextLink)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsList", nil, "Failure preparing request")
				return
			}

			result.HttpResponse, err = c.Client.Send(req, azure.DoRetryWithRegistration(c.Client))
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsList", result.HttpResponse, "Failure sending request")
				return
			}

			result, err = c.responderForSqlPoolTransparentDataEncryptionsList(result.HttpResponse)
			if err != nil {
				err = autorest.NewErrorWithError(err, "sqlpoolstransparentdataencryption.SqlPoolsTransparentDataEncryptionClient", "SqlPoolTransparentDataEncryptionsList", result.HttpResponse, "Failure responding to request")
				return
			}

			return
		}
	}
	return
}

// SqlPoolTransparentDataEncryptionsListComplete retrieves all of the results into a single object
func (c SqlPoolsTransparentDataEncryptionClient) SqlPoolTransparentDataEncryptionsListComplete(ctx context.Context, id SqlPoolId) (SqlPoolTransparentDataEncryptionsListCompleteResult, error) {
	return c.SqlPoolTransparentDataEncryptionsListCompleteMatchingPredicate(ctx, id, TransparentDataEncryptionOperationPredicate{})
}

// SqlPoolTransparentDataEncryptionsListCompleteMatchingPredicate retrieves all of the results and then applied the predicate
func (c SqlPoolsTransparentDataEncryptionClient) SqlPoolTransparentDataEncryptionsListCompleteMatchingPredicate(ctx context.Context, id SqlPoolId, predicate TransparentDataEncryptionOperationPredicate) (resp SqlPoolTransparentDataEncryptionsListCompleteResult, err error) {
	items := make([]TransparentDataEncryption, 0)

	page, err := c.SqlPoolTransparentDataEncryptionsList(ctx, id)
	if err != nil {
		err = fmt.Errorf("loading the initial page: %+v", err)
		return
	}
	if page.Model != nil {
		for _, v := range *page.Model {
			if predicate.Matches(v) {
				items = append(items, v)
			}
		}
	}

	for page.HasMore() {
		page, err = page.LoadMore(ctx)
		if err != nil {
			err = fmt.Errorf("loading the next page: %+v", err)
			return
		}

		if page.Model != nil {
			for _, v := range *page.Model {
				if predicate.Matches(v) {
					items = append(items, v)
				}
			}
		}
	}

	out := SqlPoolTransparentDataEncryptionsListCompleteResult{
		Items: items,
	}
	return out, nil
}
