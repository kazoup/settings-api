package handler

import "encoding/json"

// ResponsePublic ...
type ResponsePublic struct {
	ApplianceIsRegistered bool   `json:"APPLIANCE_IS_REGISTERED"`
	ApplianceIsConfigured bool   `json:"APPLIANCE_IS_CONFIGURED"`
	ApplianceIsDemo       bool   `json:"APPLIANCE_IS_DEMO"`
	GitCommitString       string `json:"GIT_COMMIT_STRING"`
	SmbUserExists         bool   `json:"SMB_USER_EXISTS"`
}

// GetResponse ...
func (r *ResponsePublic) GetResponse() string {
	rp := &ResponsePublic{
		ApplianceIsRegistered: true,
		ApplianceIsConfigured: true,
		ApplianceIsDemo:       false,
		GitCommitString:       "Build d150e88449 - 2016-05-04 11:52:53",
		SmbUserExists:         true,
	}

	b, _ := json.Marshal(rp)

	return string(b)
}
