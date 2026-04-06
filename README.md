#  GopherSub

**GopherSub** — Go ilə yazılmış, **Hexagonal Architecture** və **Domain-Driven Design (DDD)** prinsiplərinə əsaslanan abunəlik idarəetmə sistemidir.

Bu layihə, texniki detalları biznes məntiqindən tam ayıran, genişlənə bilən və test edilə bilən bir arxitektura nümunəsi kimi hazırlanmışdır.

---

##  Arxitektura

Layihə **Hexagonal Architecture (Ports & Adapters)** modelinə əsaslanır:

```
cmd/
└── api/             → Tətbiqin giriş nöqtəsi (main.go)

internal/
├── core/
│   ├── domain/      → Biznes məntiqinin özəyi (Entity, Value Object-lər)
│   ├── ports/       → Interfeyslər (Repository, Service, Notification)
│   └── services/    → Use case-lər (abunəlik yaratma məntiqi)
│
└── adapters/
    ├── driving/
    │   └── https/   → HTTP handler, router, DTO-lar
    └── driven/
        ├── db/      → PostgreSQL və In-Memory repository-lər
        └── notify/  → E-mail bildiriş adapteri (stub)
```

**Əsas prinsip:** Domain qatı heç bir xarici asılılıq tanımır. Bütün əlaqə portlar (interfeyslər) vasitəsilə həyata keçirilir.

---

## Domain Modelləri

| Model | Növ | Təsvir |
|---|---|---|
| `Subscription` | Aggregate Root | İstifadəçinin abunəliyini təmsil edir |
| `Plan` | Entity | Abunəlik planı (ad, qiymət, müddət) |
| `Money` | Value Object | Məbləğ və valyuta (hal-hazırda: `azn`) |
| `SubscriptionPeriod` | Value Object | Abunəliyin başlanğıc və bitmə tarixi |
| `Status` | Value Object | Abunəlik statusu (`active`, `canceled`, `expired`) |

---

## API

### `POST /api/subscribe`

Yeni abunəlik yaradır. İstifadəçinin aktiv abunəliyi varsa, əməliyyat rədd edilir.

**Request body:**
```json
{
  "user_id": "user-123",
  "plan_name": "premium",
  "amount": 1000,
  "currency": "azn"
}
```

**Uğurlu cavab:** `200 OK` — `subscription created successfully`

**Xəta halları:**
- İstifadəçinin aktiv abunəliyi artıq mövcuddursa → `500`
- Yanlış JSON formatı → `400`
- `POST` olmayan metodla müraciət → `400`

---

##  Verilənlər Bazası

Layihə **PostgreSQL** istifadə edir. Aşağıdakı cədvəllər tələb olunur:

### `subscriptions`
```sql
CREATE TABLE subscriptions (
    id             TEXT PRIMARY KEY,
    user_id        TEXT NOT NULL,
    plan_name      TEXT NOT NULL,
    start_date     TIMESTAMP NOT NULL,
    end_date       TIMESTAMP NOT NULL,
    price_amount   BIGINT NOT NULL,
    price_currency TEXT NOT NULL,
    status         TEXT NOT NULL
);
```

### `plans`
```sql
CREATE TABLE plans (
    name            TEXT PRIMARY KEY,
    price_amount    BIGINT NOT NULL,
    price_currency  TEXT NOT NULL,
    duration_days   INT NOT NULL
);
```

---

##  Quraşdırma və İşə Salma

### Tələblər
- Go `1.24+`
- PostgreSQL

### Addımlar

```bash
# 1. Repo-nu klonlayın
git clone https://github.com/Resul-Necefli/gophersub.git
cd gophersub

# 2. Asılılıqları yükləyin
go mod tidy

# 3. main.go-da connection string-i öz bazanıza uyğunlaşdırın
# "postgres://postgres:<şifrə>@localhost:5432/gophersub?sslmode=disable"

# 4. Tətbiqi işə salın
go run cmd/api/main.go
```

Server `8080` portunda işə düşür.

---

## Layihənin Dizayn Qərarları

- **Domain izolasiyası** — `Subscription` aggregate root-u öz statusunu və dövrünü özü idarə edir. Xarici qatlar heç vaxt daxili sahələrə birbaşa müraciət edə bilmir; yalnız getter metodlar vasitəsilə oxuyur.

- **`now` parametri** — Domain metodları (`IsExpired`, `IsActive`, `Renew`) vaxtı `time.Now()` ilə daxildən almır, əvəzinə `now time.Time` parametri qəbul edir. Bu, testləri deterministik edir.

- **`RestoreSubscription`** — Verilənlər bazasından gələn məlumatları yeni abunəlik kimi deyil, mövcud vəziyyəti bərpa etmək kimi modelləşdirir. ID generasiyası və ya status dəyişikliyi baş vermir.

- **In-Memory repository** — `db/in_memory_repo.go` — test və ya sürətli prototipləşdirmə üçün alternativ implementasiya mövcuddur.

- **Email adapter (stub)** — Bildiriş portu hazırdır; real implementasiya sonradan əlavə ediləcək.

---

## Asılılıqlar

| Paket | Məqsəd |
|---|---|
| `github.com/lib/pq` | PostgreSQL sürücüsü |

---

## Yol Xəritəsi

- [ ] Abunəliyin ləğvi (`/api/cancel`) endpoint-i
- [ ] Abunəliyin yenilənməsi (`/api/renew`)
- [ ] E-mail bildiriş adapterinin real implementasiyası
- [ ] JWT əsaslı autentifikasiya
- [ ] Unit testlər
- [ ] Docker & docker-compose dəstəyi

---


**Rəsul Nəcəfli** — [@Resul-Necefli](https://github.com/Resul-Necefli)
