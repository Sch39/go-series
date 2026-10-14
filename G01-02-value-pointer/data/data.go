package data

type Data struct {
	Tags   map[string]bool
	Scores []int
}

func (s Data) AppendScore() {
	s.Scores = append(s.Scores, 1)
}

func (s Data) ChangeFirstScore() {
	s.Scores[0] = 99
}
