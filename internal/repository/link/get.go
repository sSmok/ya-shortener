package link

import "errors"

func (r *repo) Get(short string) (string, error) {
	originalURL, ok := r.storage[short]
	if !ok {
		return "", errors.New("link not found")
	}

	return originalURL, nil
}
