package neurogg

/*
	Объединяет все интерфейсы в единую сеть
*/
type Conductor interface {
	Train()
	Predict()
}
