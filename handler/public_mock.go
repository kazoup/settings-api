package handler

import "encoding/json"

type ResponsePublic struct {
	APPLIANCE_IS_REGISTERED bool   `json:"APPLIANCE_IS_REGISTERED"`
	APPLIANCE_IS_CONFIGURED bool   `json:"APPLIANCE_IS_CONFIGURED"`
	APPLIANCE_IS_DEMO       bool   `json:"APPLIANCE_IS_DEMO"`
	GIT_COMMIT_STRING       string `json:"GIT_COMMIT_STRING"`
	SMB_USER_EXISTS         bool   `json:"SMB_USER_EXISTS"`
}

func (r *ResponsePublic) GetResponse() string {
	rp := &ResponsePublic{
		APPLIANCE_IS_REGISTERED: true,
		APPLIANCE_IS_CONFIGURED: true,
		APPLIANCE_IS_DEMO:       false,
		GIT_COMMIT_STRING:       "Build d150e88449 - 2016-05-04 11:52:53",
		SMB_USER_EXISTS:         true,
	}

	b, _ := json.Marshal(rp)

	return string(b)
}
