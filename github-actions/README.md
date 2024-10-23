We can use multi stage-github action

A -> B -> C

```yaml
name: Node.js CI

on:
  push:
    branches: [ "main" ]

jobs:
  build-on-ubuntu: # stage 1
    runs-on: ubuntu-latest

    strategy:
      matrix: # by using matrix we use it's variabe into the projects
        node-version: [20]
    env:
      EMAILJS_SERVICE_ID: ${{secrets.EMAILJS_SERVICE_ID}}   # these secret is from github. (repo settings -> secrets and variables -> actions)
      EMAILJS_CONTACT_TEMPLATE_ID: ${{secrets.EMAILJS_CONTACT_TEMPLATE_ID}}
      EMAILJS_PUBLIC_KEY: ${{secrets.EMAILJS_PUBLIC_KEY}}
    
    steps:
      - uses: actions/checkout@v4 # you can find in github repo (actions) there will available some repo that we can use 
                                  # here cehckout@v4(v4 is version) pull the code from where this action is triggured
      - name: Use Node.js ${{ matrix.node-version }} #1. see here we use matrix value #2. here there is step
        uses: actions/setup-node@v4
        with:
          node-version: ${{ matrix.node-version }} # again here we use matrix value
          cache: 'npm'
      - run: npm install --global yarn
      - run: yarn
      - run: yarn build
      - run: npm i sharp
      - run: rm -rf ./.git
      - run: tar -czvf archive_name.tar.gz .[^.]* ./*

      # Upload project files as artifacts, excluding node_modules
      - name: Upload project files
        uses: actions/upload-artifact@v4 # this action upload files into artifact
        with:
          name: project-files
          path: ./archive_name.tar.gz  # Upload this file in the current directory

  deploy-on-self-hosted: # stage 2
    needs: build-on-ubuntu # it says to run this stage it needs (build-on-ubuntu(in our ctx stage 1)).
    runs-on: self-hosted

    steps:
      # Download the artifact containing project files
      - name: Download project files
        uses: actions/download-artifact@v4
        with:
          name: project-files # download this file(project-files) from atrifact
          path: .  # Download into the current directory
      - name: Untar file
        run: tar -xzvf archive_name.tar.gz
      - name: Change directory
        run: cd /home/ubuntu/actions-runner/_work/amirat-lube/amirat-lube
      - name: remove tar file
        run: rm -rf archive_name.tar.gz
      - run: pm2 restart 0
```