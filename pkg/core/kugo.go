package core

type Kudo struct {
	UID         string `json:"id" bson:"uId"`
	UserID      string `json:"user_id" bson:"userId"`
	UserName    string `json:"login" bson:"userName"`
	UserURL     string `json:"html_url" bson:"userUrl"`
	Avatar      string `json:"avatar_url" bson:"avatar"`
	Description string `json:"subscriptions_url" bson:"description"`
}
