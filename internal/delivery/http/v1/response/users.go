package response

type UserAccessToken struct {
	Email       string `json:"email"`
	AccessToken string `json:"access_token"`
}
