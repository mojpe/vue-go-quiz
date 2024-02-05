# vue-go-quiz
A small app with Vue.js and GO

## Install
``` bash
# Download and install Vue client NPM packages
npm create vue@latest
Done. Now run:

cd quiz
npm install
npm run format
npm run dev
```
# Vue packages
```cd web/
npm install vue-resource
npm bootstrap-vue
```

# Download and install this package
go get github.com/gorilla/mux

# Run
### Go app
Server will enable CORS headers from localhost:8888. Running the Vue client will automatically open the browser at http://localhost:8888/
``` bash
# Start API server at :8888
quiz

### Vue
# Start Vue client with at :8080
cd web/
npm run dev
```

Now navigate to http://localhost:8080/.
