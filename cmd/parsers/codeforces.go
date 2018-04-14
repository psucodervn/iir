package parsers

// CodeforcesParser parser for Codeforces
type CodeforcesParser struct {
}

func (*CodeforcesParser) ParseTaskFromHTML(html string) (Task, error) {
	return nil, nil
}

func (*CodeforcesParser) ParseTaskFromURL(url string) (Task, error) {
	panic("implement me")
}

func (*CodeforcesParser) ParseContestFromHTML(html string) (Contest, error) {
	panic("implement me")
}

func (*CodeforcesParser) ParseContestFromURL(url string) (Contest, error) {
	panic("implement me")
}
