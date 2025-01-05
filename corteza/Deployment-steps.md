Corteza deployment steps:

open project 


1. cd lib/vue/
	- yarn
	- yarn build

2. cd client/web/compose/
	- yarn
	- yarn link @cortezaproject/corteza-vue
	- yarn link @cortezaproject/corteza-js
	- yarn build --production
	- a `dist` file will created
	- move `dist` to  /cortezafe/
	- delete /cortezafe/compose/
	- rename `dist` to `compose`

3. cd server/
	- make release
	- a `build` directory will create or modify
	- move `server/build/pkg/corteza-server/bin/corteza-server`  to `corteza-server/bin/`
	
	
4. go to the project root directory
	- docker build --no-cache -t devopsaes/corteza:1.17 .
	- docker push devopsaes/corteza:1.17
