# Gitlab user remove 

Use this tool to delete user from provate repos and groups which you can access with gitlab token.

## Use with docker 

```
# -- pull the image
$ docker pull ghcr.io/allanger/gitlab-remove-user
# -- get the user id
# -- it will print id / name / username 
$ docker run ghcr.io/allanger/gitlab-remove-user search -u $USERNAME -t $GITLAB_TOKEN
# -- remove runs dry-run if --dry-run false is not specified 
# -- it will print from which project the user will be removed
$ docker run ghcr.io/allanger/gitlab-remove-user remove -u $USER_ID $-t $GITLAB_TOKEN
# -- if you're happy for this tool to remove user, run
$ docker run ghcr.io/allanger/gitlab-remove-user remove -u $USER_ID $-t $GITLAB_TOKEN --dry-run false
```

It also can read GTILAB_TOKEN from environment variables