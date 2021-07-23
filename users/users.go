package users

import (
	g "github.com/allanger/gitlab-remove-user/third_party/gitlab"
	log "github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

func Search(user string, token string) {
	if len(user) == 0 {
		log.Fatal("specify gitlab user id with --user flag")
	} else if len(token) == 0 {
		log.Fatal("specify gitlab token with --token flag")
	}

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
