package neuro_getero

// Интерфейс тикера
// Реализуется объектами, работа которых должна быть синхронизирована с тикером
type Ticker interface {
	Tick(tick interface{})
}
