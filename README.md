# goapp-clock

## Deps
1. Golang 1.21
2. Node 21 / NPM
3. [Taskfile](https://taskfile.dev/)

## Build Process
1. Install Deps
2. Clone Repo
3. `task npm-deps`
4. `task run`
5. Visit [http://127.0.0.1:8000](http://127.0.0.1:8000) in a web browser

`task run` Breakdown
- Build the webapp client (ie app.wasm)
- Run the server executable


## GH Pages Build Process

Much thanks to [pojntfx/liwasc](https://github.com/pojntfx/liwasc) for demystifying the go-app static build process a bit

(in the GH action workflow)
1. Clone repo
2. Install Deps
3. `task gh-pages`
4. Release the newly created `/docs` directory to the gh-pages branch

`task gh-pages` Breakdown
- Build the webapp client (ie app.wasm)
- Run the build executable to use go-app's internal static site generation function (creates `/docs`)
- Move the contents of the web folder (the output of the the webapp client build) into the `/docs` directory