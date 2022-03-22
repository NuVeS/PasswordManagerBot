package handler

import (
	"fmt"
	"strings"

	"github.com/NuVeS/PasswordMangerBot/storage"
)

// Формат ввода команды
// get apple.com - получить пароль для apple.com
// get all - получить все пароли
// add apple.com Querty12 - установить пароль для apple.com на Querty12
// update apple.com Querty123 - обновить пароль для apple.com на Querty123
// remove apple.com - удалить пароль для apple.com
func HandleMessage(message string, chatId string) string {

	storage := storage.NewInstance()
	fields := strings.Fields(message)

	switch fields[0] {
	case "get":
		if res := storage.Get(chatId, fields[1]); len(res) > 0 {
			return res
		} else {
			return "У меня его нет"
		}
	case "getall":
		if res := storage.GetAll(chatId); len(res) > 0 {
			var arr string = ""
			for key, value := range res {
				temp := fmt.Sprintf("%s : %s", key, value)
				arr = fmt.Sprintf("%s\n%s", arr, temp)
			}
			return arr
		} else {
			return "Ты пока не сохранялся"
		}
	case "update":
		if res := storage.Update(chatId, fields[1], fields[2]); res {
			return "Готово"
		} else {
			return "Ошибка!"
		}
	case "remove":
		if res := storage.Remove(chatId, fields[1]); res {
			return "Готово"
		} else {
			return "Ошибка!"
		}
	default:
		return ""
	}
}
