package servicetest

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	gk "github.com/onsi/ginkgo"
	gm "github.com/onsi/gomega"
	"github.com/sanksons/reflorest/src/test/api"
	testUtil "github.com/sanksons/reflorest/src/test/utils"
)

func testRequestValidation() {
	var tstData = new(api.TestData)
	bt, err := ioutil.ReadFile("testdata/requestValidationNode.json")
	if err != nil {
		fmt.Sprintf("Error loading Request Validation Node Data file %s \n %s", err)
	}
	err = json.Unmarshal(bt, tstData)
	apiName := "florest"
	gk.Describe("GET"+"/"+apiName+"/V1/REQVD/", func() {
		request := testUtil.CreateTestRequest("GET", "/"+apiName+"/V1/REQVD/")
		response := GetResponse(request)
		gk.Context("then the response", func() {
			gk.It("should return validation failure message", func() {
				gm.Expect(response.HeaderMap.Get("Content-Type")).To(gm.Equal("application/json"))
				gm.Expect(response.HeaderMap.Get("Cache-Control")).To(gm.Equal(""))
				gm.Expect(response.Code).To(gm.Equal(400))
				validateRequestValidationResponse(response.Body.String(), tstData.RequestValidationMessage)
			})
		})
	})

}
