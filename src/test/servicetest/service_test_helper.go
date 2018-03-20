package servicetest

import (
	"encoding/json"
	gm "github.com/onsi/gomega"
	"github.com/sanksons/reflorest/src/common/constants"
	utilhttp "github.com/sanksons/reflorest/src/common/utils/http"
)

func validateHealthCheckResponse(responseBody string) {
	var utilResponse utilhttp.Response
	err := json.Unmarshal([]byte(responseBody), &utilResponse)
	gm.Expect(err).To(gm.BeNil())

	if _, ok := utilResponse.Data.(map[string]interface{}); ok {
		gm.Expect(ok).To(gm.Equal(true))
	}
}

func validateRequestValidationResponse(responseBody string, errmsg []string) {
	var utilResponse utilhttp.Response
	_ = json.Unmarshal([]byte(responseBody), &utilResponse)
	errleng := len(errmsg)
	m := make(map[string]string, errleng)
	for i := 0; i < errleng; i++ {
		temp := errmsg[i]
		m[temp] = temp
	}
	var leng int = len(utilResponse.Status.Errors)
	for i := 0; i < leng; i++ {
		var err constants.AppError = utilResponse.Status.Errors[i]
		_, found := m[err.Message]
		gm.Expect(true).To(gm.Equal(found))
	}
}
