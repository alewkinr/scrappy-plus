package internal

// Event — структура события
type Event struct {
	// IsAvailable — флаг что событие доступно
	IsAvailable bool
	// Title — навзание события
	Title string
	// BookURL — ссылка на бронирование
	BookURL string
}

// EventList — список событий
type EventList struct {
	// Events — события включенные в список
	Events []Event
}
