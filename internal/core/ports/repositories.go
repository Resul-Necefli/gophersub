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
	GetPlanByName(name string) (*domain.Plan, error)
}

// GetPlanByName(name string) (*domain.Plan, error)

// bu metodun  sebebi odurki  biz  abunelik ucun  planlarimiz var ve biz istfadeciye  plan adina gore  abunelik yaratmaq isteyirik
// men  bunu sonradan elave etdim  normalda heqiqi projecte hersey qabaxcadan cox cidi sekilde planlanmalidir  amma biz  bu projece
// faza faza  irlesdiririk ve bu metodun  elave edilmesi  bizim  planlarimiz ucun  repository de  bir metodun  elave edilmesi  lazim idi
//  cunki biz  plan adina gore  abunelik yaratmaq isteyirik ve bu metod bize  plan adina gore  plan obyektini qaytaracaq ki bizde o obyektin
//  icinde olan deyerlerden istfade ederek  abunelik yarada bilik

//  bu sekildede ayira bilerdim amma  helelik ayrmiram ve   SOLID  I sini pozuram bilerekden

// type Subscribe interface {
// 	Subscribe(userID, planName string, amount int64, currency string) error
// }
