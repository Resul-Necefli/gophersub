package ports

type Subscribe interface {
	Subscribe(userID, planName string, amount int64, currency string) error
}
