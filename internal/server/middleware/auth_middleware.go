package middleware

import "net/http"

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Попытка извлечь токен из куки
		token, err := r.Cookie("auth_token")
		if err != nil || token.Value == "" {
			// Попытка извлечь токен из заголовков
			tokenValue := r.Header.Get("Authorization")
			if tokenValue == "" {
				tokenValue = r.Header.Get("Bearer")
			}
			if tokenValue == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}
		}

		// Продолжение выполнения следующего обработчика
		next.ServeHTTP(w, r)
	})
}
