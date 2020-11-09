package local

type Signal struct {
	Value int64
}

func (s *Signal) GetValue() interface{} {
	return s.Value
}
