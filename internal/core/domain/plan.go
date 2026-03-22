package domain

// Plan struct-ı Reference Data (sabit məlumat) olduğu üçün
// sahələri Açıq (Public) saxlamaq daha rahatdır.
// Beləliklə Repository onlara rahatca yaza biləcək.
type Plan struct {
	Name     string
	Price    Money
	Duration int
}
