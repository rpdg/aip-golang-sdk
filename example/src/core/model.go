package core

// page meta
type Page struct {
	UserID   int
	Username string
	Title    string
	Account  interface{}
	CSS      []string
	JS       []string
	Message  []string
	Error    string
}

type UserItem struct {
	Group_id  string
	Score     float64
	User_id   string
	User_info string
}

type Result struct {
	Face_token string
	User_list  []UserItem
	Timestamp  float64
}

type SearchResult struct {
	Cached     float64
	Error_code float64
	Error_msg  string
	Log_id     float64
	Result
}
