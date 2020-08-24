package neuro_getero

// Возбуждаемый объект (аналог нейрона)
// Принимает возбуждение и обрабатывает его
// В итоге - отправляет (или не отправляет) результатирующее возбуждение в пучок коннекторов ConnectorCluster
type Excitable interface {
	SetConnectorCluster(ConnectorCluster)
	GetConnectorCluster() ConnectorCluster

	Excite(Excitation)
}
