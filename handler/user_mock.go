package handler

import "encoding/json"

type company struct {
	Plan        string `json:"plan"`
	PlanStarted int64  `json:"plan_started"`
	ID          int64  `json:"id"`
	Name        string `json:"name"`
}

type intercom struct {
	CreatedAt     int64   `json:"created_at"`
	AppID         string  `json:"app_id"`
	ApplianceUUID string  `json:"Appliance UUID"`
	Build         string  `json:"Build"`
	Email         string  `json:"email"`
	UserHash      string  `json:"user_hash"`
	Company       company `json:"company"`
}

// ResponseUser ...
type ResponseUser struct {
	URL        string   `json:"url"`
	ID         int64    `json:"id"`
	UserName   string   `json:"username"`
	GroupName  string   `json:"group_name"`
	LastLogin  string   `json:"last_login"`
	Email      string   `json:"email"`
	DateJoined string   `json:"date_joined"`
	Created    string   `json:"created"`
	Updated    string   `json:"updated"`
	Intercom   intercom `json:"intercom"`
}

// GetResponse ...
func (r *ResponseUser) GetResponse() string {
	ru := ResponseUser{
		URL:        "https://kazoup.com/api/v1/users/2/",
		ID:         2,
		UserName:   "pablo.aguirre@kazoup.com",
		GroupName:  "admin",
		LastLogin:  "2015-12-16T15:52:32.276907Z",
		Email:      "",
		DateJoined: "2015-12-16T15:52:32.276968Z",
		Created:    "2015-12-16T15:52:32.281424Z",
		Updated:    "2016-05-17T13:20:22.981312Z",
		Intercom: intercom{
			CreatedAt:     1450281152,
			AppID:         "udh2t9k4",
			ApplianceUUID: "a7d91565-f99c-4f4f-a396-e7ff6ae082b6",
			Build:         "Build f054e71e2f - 2016-05-16 21:53:08",
			Email:         "pablo.aguirre@kazoup.com",
			UserHash:      "8ad1df17fff865e90bc21a30b36875d1d4da15c611b077aaff378f6753235dd4",
			Company: company{
				Plan:        "Business",
				PlanStarted: 1452164082,
				ID:          1,
				Name:        "Kazoup",
			},
		},
	}

	b, _ := json.Marshal(ru)

	return string(b)
}
