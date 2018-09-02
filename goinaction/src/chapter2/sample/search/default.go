package search

func init() {
	var matcher defaultMatcher
	Register("default", matcher)
}

type defaultMatcher struct{}

func (m defaultMatcher) Search(feed *Feed, searchTerm string) ([]*Result, error) {
	return nil, nil
}
