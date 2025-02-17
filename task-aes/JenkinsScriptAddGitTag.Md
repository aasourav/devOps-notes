Add these in deploy stage

```md
sh """
	set -e
	cd ${REPO_NAME}
	git config --global user.email "bot@jenkins.com"
	git config --global user.name "Jenkins bot"
	git -c http.sslVerify=false tag -a ${DOCKER_USER_NAME}/${DOCKER_APP_NAME}-${IMAGE_TAG} -m "${DOCKER_USER_NAME}/${DOCKER_APP_NAME}:${GIT_TAG}"
	git -c http.sslVerify=false push origin HEAD --tags"""
```