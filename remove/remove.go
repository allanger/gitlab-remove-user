package remove

import (
	"sync"

	g "github.com/allanger/gitlab-remove-user/third_party/gitlab"
	log "github.com/sirupsen/logrus"
	"github.com/xanzy/go-gitlab"
)

func Remove(userID int, dry bool, token string) {
	if userID < 0 {
		log.Fatal("specify gitlab user id with --user flag")
	} else if len(token) == 0 {
		log.Fatal("specify gitlab token with --token flag")
	}
	var wg sync.WaitGroup
	git := g.GetCLient(token)
	for _, g := range getGroups(git) {
		wg.Add(1)
		go removeFromGroup(git, g, userID, dry, &wg)
		for _, p := range getProjects(git, g) {
			wg.Add(1)
			go removeFromProject(git, p, userID, dry, &wg)
		}
	}
	wg.Wait()
	return
}

var getGroups = func(git *gitlab.Client) []*gitlab.Group {
	var (
		page    = 1
		perPage = 10
		groups  []*gitlab.Group
	)
	for {
		groupsOpt := &gitlab.ListGroupsOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: perPage,
				Page:    page,
			}}
		groupsFromPage, status, err := git.Groups.ListGroups(groupsOpt)
		if err != nil {
			log.Fatal(err, status)
		}
		groups = append(groups, groupsFromPage...)
		page += 1
		if len(groupsFromPage) == 0 {
			break
		}
	}
	return groups
}

var getProjects = func(git *gitlab.Client, group *gitlab.Group) []*gitlab.Project {
	var (
		page     = 1
		perPage  = 10
		projects []*gitlab.Project
	)
	for {
		projectsOpt := &gitlab.ListGroupProjectsOptions{
			ListOptions: gitlab.ListOptions{
				PerPage: perPage,
				Page:    page,
			}}
		projectsFromPage, status, err := git.Groups.ListGroupProjects(group.ID, projectsOpt)
		if err != nil {
			log.Fatal(err, status)
		}
		projects = append(projects, projectsFromPage...)
		page += 1
		if len(projectsFromPage) == 0 {
			break
		}
	}
	return projects
}

func removeFromGroup(git *gitlab.Client, group *gitlab.Group, userID int, dry bool, wg *sync.WaitGroup) {
	defer wg.Done()
	member, _, err := git.GroupMembers.GetGroupMember(group.ID, userID)
	if err == nil {
		log.Infof("removing %s from %s", member.Name, group.Name)
		if !dry {
			_, err = git.GroupMembers.RemoveGroupMember(group.ID, userID)
			log.Infof("removed %s from %s", member.Name, group.Name)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}

func removeFromProject(git *gitlab.Client, project *gitlab.Project, userID int, dry bool, wg *sync.WaitGroup) {
	defer wg.Done()
	member, _, err := git.ProjectMembers.GetProjectMember(project.ID, userID)
	if err == nil {
		log.Infof("removing %s from project %s", member.Name, project.NameWithNamespace)
		if !dry {
			_, err = git.ProjectMembers.DeleteProjectMember(project.ID, userID)
			log.Infof("removed %s from project %s", member.Name, project.NameWithNamespace)
			if err != nil {
				log.Fatal(err)
			}
		}
	}
}
