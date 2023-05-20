# webcuss | extension

## Run dev
```shell
$ export REACT_APP_BACKEND_URL=http://localhost:8080
$ npm install
$ npm start
```

## Build
```shell
$ export REACT_APP_BACKEND_URL=http://localhost:8080
$ npm run build
```

## Package
Requires [build](#build)
```shell
$ npm run package
```
the package output is located at [build/webcuss.zip](build/webcuss.zip), use this file as release package to Chrome Webstore.
## Import extension
Goto Google Chrome > Extension > Enable `Developer Mode` > Load unpacked > select `webcuss/extension`

Next, open extension! 🚀

## Keyboard shorcut
* `Ctrl+Shift+Comma` - Toggle debugging mode
