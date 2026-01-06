package ports

import "github.com/Resul-Necefli/gophersub/internal/core/domain"

//  burada biz   DDD  ve hexegonal  strukturnda gordyumuz kimi  hem servis hemdeki   bu  port hissede proqramizin
// biznes mentiqini dirildecek onu bir  butun edecek davranislar ucun qebul edilme desikleri   noqteler yaradiriq  yenu portlar
// bunlari interface sekilde  teyin edirk ki inject  formasinda qebul edek ve bizim hecbir texniki  asliligimiz olmasin
// biz abunelik sistemi uzerinde islediyimiz ucun sadece olaraq  abuneliyin yaradilmazi abunecinin
//
//  spesfik olaraq tapilmasi  ve onun yenilenmesi
// metodlarini elave edirik amma bura delete  elave ede bilerek hele bu beledir amma ola bilerki elave edel cunki istfadeci Cancled olunduqda bunu
// ayrica bir handlerden  ayrica bir resurs kimi yanasilmasini menim fikrimce teyin etmeliyik

// yuxardaki  cumleni ayrdim cunki DDD ruhuna ziddir   buz   DDD zamani CRUD api cox  fokus olmaliyiq yenu bir sozle port da buna cox foks olmagimiz
// DDD  texniki  detallaa fokslamis olur  buda DDD ye zidd meseledir

type SubscriptionRepository interface {
	Save(sub *domain.Subscription) error
	GetByID(id string) (*domain.Subscription, error)
	GetByUserID(userID string) ([]*domain.Subscription, error)
}

// - `SubscriptionRepository` adlı bir `interface` yarat.

//     - Metodlar:

//         - `Create(sub *domain.Subscription) error`

//         - `GetByID(id string) (*domain.Subscription, error)`

//         - `Update(sub *domain.Subscription) error`
