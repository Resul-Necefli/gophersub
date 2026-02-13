package main

// Faza 4: Application Assembly (Main.go)

// Adapterimiz artıq "Thread-Safe"dir. İndi o böyük an gəldi. Bütün hissələri birləşdirib proqramı işə salaq.

// Task 9: Dependency Injection (cmd/api/main.go)

// Bu faylı yarat və aşağıdakı addımları yerinə yetir:

//     Repo-nu yarat: repo := db.NewInMemorySubscriptionRepository()

//     Servisi yarat: svc := services.NewSubscriptionService(repo)

//     Test et: Kodla birbaşa svc.Subscribe(...) çağır.

//     Nəticə: Ekrana (Log) nəticəni yazdır.

// Sual: Subscribe metodunda userID olaraq "resul_123" göndərdin. Proqramı bağlayıb təzədən açsan və yenə eyni ID ilə abunə olmağa çalışsan, xəta alacaqsan? (InMemory repo-nun təbiətini düşün).

func main() {

	// repo := db.NewInMemorySubscriptionRepository()

	// // servs := services.NewSubscriptionService(repo)

	// handle := https.NewSubscriptionHandler(*servs)

	// http.HandleFunc("/subscribe", handle.Subscribe)

	// log.Fatal(http.ListenAndServe(":8080", nil))
}
