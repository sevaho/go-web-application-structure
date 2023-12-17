package integration

import (
	"net/http/httptest"
)

func (suite *TestSuite) TestIndex() {
	// given
	req := httptest.NewRequest("GET", "/", nil)

	// when
	resp, _ := suite.httpIngress.Server.Test(req)

	// then
	suite.Assert().Equal(302, resp.StatusCode)
}
