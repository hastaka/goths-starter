# goths-starter
GoTTHS starter pack

A system for building web apps without using React, Angular, Svelte, Vue or any other super-heavy framework. 

### Stack
* [Go](https://go.dev) - web server, backend system, integration with databases (such as [PocketBase](https://pocketbase.io/)!)
* [Templ](https://templ.guide/) - HTML composition engine to build HTML using Go, entirely on server-side
* [Tailwind CSS](https://tailwindcss.com/docs/) - CSS styling framework with minimal footprint
* [HTMX](https://htmx.org/docs/) - interaction library for AJAX, WebSockets, and more
* [Shoelace](https://shoelace.style/) - frontend web-components library

Supporting Tooling
* [Air](https://github.com/air-verse/air) - live reload development tool
* [Iconify](https://iconify.design/) - multi-icon framework for any icon set

_Thanks to [Murtaza Udaipurwala](https://blog.murtazau.xyz/templ-tailwind-htmx) and [Dave Eddy](https://www.daveeddy.com/) for inspiration_

### Basic Usage
* Generate template: `templ generate`
* Tailwind live rebuild: `npm run watch`
* Run Go backend: `go run .`
* Run Air (live reload): `air`

### Installation
> [!IMPORTANT]
> #### Follow the below instructions if starting from this template.
> 1. On Github click "Use this template" on the top right and create your own repository.
> 2. Clone repository to your machine, and ensure Go, Templ, NPM, Node.js, and Air are installed (following the installation directions below). Configurations do not need to be changed or created.
>     * Go: Install Go from the [Go website](https://go.dev).
>     * Templ: Run `go install github.com/a-h/templ/cmd/templ@latest` to install templ.
>     * NPM/Node.js: Install the latest version from [NPM and Node.js](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm).
>     * Air: Install Air using `go install github.com/air-verse/air@latest`.
> 3. Install required packages for Tailwind using `npm install`.
> 4. Run `air`, then navigate to [localhost:8080](http://localhost:8080).
#### Go
1. Install Go from the [Go website](https://go.dev).
2. Initialize Go project using `go mod init <project>/<module>`
> [!NOTE]
> If you are using Github, I recommend using the Go template for your `.gitignore` file. Otherwise you can copy the below:
> ```gitignore
> # If you prefer the allow list template instead of the deny list, see community template:
> # https://github.com/github/gitignore/blob/main/community/Golang/Go.AllowList.gitignore
> #
> # Binaries for programs and plugins
> *.exe
> *.exe~
> *.dll
> *.so
> *.dylib
> 
> # Test binary, built with `go test -c`
> *.test
> 
> # Output of the go coverage tool, specifically when used with LiteIDE
> *.out
> 
> # Dependency directories (remove the comment below to include it)
> # vendor/
> 
> # Go workspace file
> go.work
> go.work.sum
> 
> # env file
> .env
> ```
#### Templ
1. Run `go install github.com/a-h/templ/cmd/templ@latest` to install templ.
2. Run `go get github.com/a-h/templ` to add to project.
3. You can now generate templates by running `templ generate`.
> [!NOTE]
> If you are using Visual Studio Code, I recommend hiding the generated *_templ.go files using the following `.vscode/settings.json` configuration:
> ```json
> {
>     "files.exclude": {
>         "**/*_templ.go": true
>     }
> }
> ```
#### Tailwind CSS
1. Ensure you have installed the latest version of [NPM and Node.js](https://docs.npmjs.com/downloading-and-installing-node-js-and-npm) (only needed for installation and generation of CSS).
2. Initialize node module using `npm init -y`.
3. Install Tailwind using `npm install -D tailwindcss`.
> [!NOTE]
> If you are using git, make sure to add `node_modules` to your `.gitignore` file!
4. Initialize Tailwind `npx tailwindcss init`.
5. Modify `tailwind.config.js` as follows to check for templ files:
    ```javascript
    ...
    module.exports = {
        content: ["./app/**/*.templ"], // Check all *.templ files
            theme: {
                extend: {},
            },
        plugins: [],
    }
    ...
    ```
6. Modify `package.json` to create convenience scripts:
    ```javascript
    ...

    "scripts": {
        ...

        "build": "tailwindcss build -o static/css/tailwind.css --minify",
        "watch": "tailwindcss build -o static/css/tailwind.css --watch"
    }
    
    ...
    ```
7. You can now run Tailwind live rebuild with `npm run watch`.
#### Air
1. Install Air using `go install github.com/air-verse/air@latest`.
2. Create the `.air.toml` file as follows:
    ```toml
    root = "."
    tmp_dir = "air"

    [build]
    bin = "./air/main.exe"
    cmd = "templ generate && npm run dev && go build -o ./air/main.exe ."
    delay = 1000
    exclude_dir = ["static", "node_modules"]
    exclude_regex = [".*_templ.go"]
    exclude_unchanged = false
    follow_symlink = false
    include_ext = ["go", "tpl", "tmpl", "templ", "html"]
    kill_delay = "0s"
    log = "build-errors.log"
    send_interrupt = false
    stop_on_error = true

    [color]
    build = "yellow"
    main = "magenta"
    runner = "green"
    watcher = "cyan"

    [log]
    time = false

    [misc]
    clean_on_exit = true
    ```
3. You can now run Air for live reload by using `air`. If you have not run `templ generate` yet, do so first to create the initial files. After this setup, only `air` needs to be run to generate TEMPL files, compile Tailwind, and run the Go server.
> [!NOTE]
> If you are using git, make sure to add `air` to your `.gitignore` file!
### Project Structure
#### `main.go`
* Contains the main routing logic for the app.
* Uses `net/http` for handling. Can be replaced with other routing libraries if so desired.
#### `static/`
* Contains statically served files, like CSS and images.
* The file `static/css/tailwind.css` is automatically generated by Tailwind and should not be modified!
* The file `static/css/site.css` is used for global classes, and is imported in the root layout.
#### `utils/`
* Contains utility functions to be used in templ templates or elsewhere.
* Contains `If` function as example.
#### `server/`
* Here lives server-side code. Code can be called from other files, like templ templates, and executed server-side.
#### `app/`
* Here is all the code related to the front-end application, mostly made of templ files.
#### `app/components/`
* This folder contains re-usable components of the application.
#### `app/layout/`
* This folder contains layouts, which wrap around individual pages. This can contain things like headers and footers which are re-used.
> [!IMPORTANT]
> The file `layout.templ` is the root layout of the application. In this layout, we add the imports for the JS libraries used for the stack, as well as CSS:
> ```html
> <!-- Tailwind CSS -->
> <link href="/static/css/tailwind.css" rel="stylesheet"/>
> <!-- Site-specific global CSS -->
> <link href="/static/css/site.css" rel="stylesheet"/>
> <!-- HTMX -->
> <script src="https://unpkg.com/htmx.org@1.9.12"></script>
> <!-- Iconfiy-Icon -->
> <script src="https://cdn.jsdelivr.net/npm/iconify-icon@2.1.0/dist/iconify-icon.min.js"></script>
> <!-- Shoelace Components -->
> <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.18.0/cdn/themes/light.css"/>
> <script type="module" src="https://cdn.jsdelivr.net/npm/@shoelace-style/shoelace@2.18.0/cdn/shoelace-autoloader.js"></script>
> ```
#### `app/pages/`
* This folder contains individual pages of the application.
* Together with layouts, these are called in `main.go`.
#### `app/partial/`
* This folder contains partials, which can be composed in pages, components, or other partials using HTMX and Templ.
