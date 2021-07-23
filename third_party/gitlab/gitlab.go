package gitlab

import (
	log "github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

var GetCLient = func(token string) *gitlab.Client {
	git, err := gitlab.NewClient(token)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	return git
}
