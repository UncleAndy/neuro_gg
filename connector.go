package neurogg

// Передает полученное возбуждение после обработки цели типа Excitable
type Connector interface {
	SetTarget(e Excitable)

	ProcessExcitation(e Excitation)
}

// Пучок коннекторов для передачи полученного возбуждения группе целей
type ConnectorCluster interface {
	Append(c Connector)

	ProcessExcitation(e Excitation)
}
