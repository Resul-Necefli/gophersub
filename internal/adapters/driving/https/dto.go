package https

type SubscribeRequest struct {
	UserID   string `json:"user_id"`
	PlanName string `json:"plan_name"`
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}
