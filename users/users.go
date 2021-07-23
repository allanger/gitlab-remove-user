package users

import (
	g "github.com/allanger/gitlab-remove-user/third_party/gitlab"
	log "github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

func Search(user string, token string) {
	git := g.GetCLient(token)
	userOpt := &gitlab.ListUsersOptions{
		Search: gitlab.String(user),
	}
	users, _, err := git.Users.ListUsers(userOpt)
	if err != nil {
		log.Fatal(err)
	}
	for _, user := range users {
		log.Infof("id - %d / name - %s / username - %s", user.ID, user.Name, user.Username)
	}
	return
}
