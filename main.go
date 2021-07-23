package main

import (
	"github.com/allanger/gitlab-remove-user/cmd"
)

func main() {
	cmd.Execute()
}

// const userID = 7322656

// func main() {

// 	argsWithProg := os.Args
// 	argsWithoutProg := os.Args[1:]
// 	arg := os.Args[3]

// 	fmt.Printf("%s, %s, %s", argsWithProg, argsWithoutProg, arg)
// 	git, err := gitlab.NewClient("vLZu6Gp4XtvmkL_NBQX5")
// 	if err != nil {
// 		log.Fatalf("Failed to create client: %v", err)
// 	}
// 	removeFromGroups(git)
// 	// removeFromProjects(git)
// }


// func removeFromProjects(git *gitlab.Client, group int, wg sync.WaitGroup) {
// 	var (
// 		page     = 0
// 		perPage  = 50
// 		projects []*gitlab.Project
// 		wg1      sync.WaitGroup
// 	)

// 	for {
// 		projectsOpt := &gitlab.ListGroupProjectsOptions{
// 			ListOptions: gitlab.ListOptions{
// 				PerPage: perPage,
// 				Page:    page,
// 			}}
// 		projectsFromPage, status, err := git.Groups.ListGroupProjects(group, projectsOpt)
// 		if err != nil {
// 			log.Fatal(err, status)
// 		}
// 		projects = append(projects, projectsFromPage...)
// 		page += 1
// 		if len(projectsFromPage) == 0 {
// 			break
// 		}
// 	}
// 	for _, p := range projects {
// 		log.Infof("project: %s", p.Name)
// 		wg1.Add(1)
// 		go func(id int, name string) {
// 			defer wg1.Done()
// 			member, _, err := git.ProjectMembers.GetProjectMember(id, userID)
// 			if err != nil {
// 				return
// 			}
// 			log.Infof("!! removing user %s from project %s\n\n", member.Name, name)
// 			return
// 		}(p.ID, p.Name)
// 	}
// 	wg1.Wait()
// 	return
// }
