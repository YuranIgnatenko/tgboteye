package bot

func (b *Bot) MenuMain() []string {
	names_button := make([]string, 0)
	if RecoverFlag == true {
		names_button = append(names_button, " Восстановить")
	}
	names_button = append(names_button, "Добавить")
	for _, t := range b.Tasks {
		names_button = append(names_button, StyleStatus(t.TargetUser))
	}
	if len(b.Tasks) > 0 {
		names_button = append(names_button, "Остановить все")
	}
	names_button = append(names_button, "История")
	return names_button
}

func (b *Bot) MenuSelectUser() []string {
	names_button := make([]string, 0)
	names_button = append(names_button, "Активность")
	names_button = append(names_button, "Остановить")
	names_button = append(names_button, " Забыть")
	names_button = append(names_button, "Меню")
	return names_button
}

func (b *Bot) MenuAddUser() []string {
	names_button := make([]string, 0)
	names_button = append(names_button, "Сохранить")
	names_button = append(names_button, " Забыть")
	names_button = append(names_button, "Отмена")
	return names_button
}

func (b *Bot) MenuHistory() []string {
	names_button := make([]string, 0)
	names_button = append(names_button, "Удалить")
	names_button = append(names_button, "Отмена")
	return names_button
}
