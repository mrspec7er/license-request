package util

// type key string

// const (
// 	UserContextKey key = "user"
// )

// var Store *sessions.CookieStore

// func AuthInit() {

// 	Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_SECRET")))
// 	maxAge := 86400 * 1 // 1 days
// 	isProd := false     // Set to true when serving over https

// 	Store.MaxAge(maxAge)
// 	Store.Options.Path = "/"
// 	Store.Options.HttpOnly = true
// 	Store.Options.Secure = isProd

// 	gothic.Store = Store

// 	goth.UseProviders(
// 		google.New(os.Getenv("GOOGLE_AUTH_KEY"), os.Getenv("GOOGLE_AUTH_SECRET"), os.Getenv("API_URL")+"/auth/callback?provider=google"),
// 	)
// }
