package greet

import "strings"

// Greet приймає ім'я користувача і повертає рядок привітання.
//
// Правила:
//   - для непорожнього імені (пробіли на краях обрізаються):
//     "Hello, <name>! Welcome to Go."
//   - якщо після обрізання пробілів ім'я порожнє:
//     "Hello, stranger! Welcome to Go."
func Greet(name string) string {
	// Обрізаємо пробіли з обох боків рядка
	trimmedName := strings.TrimSpace(name)

	// Перевіряємо, чи рядок порожній після обрізання
	if trimmedName == "" {
		return "Hello, stranger! Welcome to Go."
	}

	// Повертаємо привітання з ім'ям
	return "Hello, " + trimmedName + "! Welcome to Go."
}
