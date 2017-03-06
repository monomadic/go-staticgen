# go-staticgen
A simpler re-thinking of static site generators in golang. If you're a developer/designer who manages a group of static sites, and you want a better workflow, this might be a good option for you, as it revolves around the idea of code/module re-use across static sites. You could build up a 'reserve' of assets such as fonts, placeholder images, regularly used styles and framework files that are always available and you can update or manage them all in one simple directory structure without crazy dependencies or messing with the black hole of stupid that is node.js.

![screenshot][screenshot.png]

## OMFG ANOTHER STATIC SITE GENERATOR?
Wow. Calm down. After trying many static generators I realised there's a core problem with all of them - they focus on non-technical users who want to make blogs. Nothing against this kind of user, but it's not me. I want to make actual websites that have varied structure. I don't need the program to force me into a structure I dont want to work in, which usually is:

- you are told to write markdown files as blog posts, as a first priority.
- the idea of actual design and layout is pushed into garbagey 'templates' you're supposed to download separately and muck around with.
- you don't like the way the tool works? changing it requires reading some huge documentation library and it probably doesn't support it because you aren't the target market.
- you wrestle with the primary layout file (for example, in jekyll hybrids, and find it's impossible to create your enclosing layout in haml, you have to futz around with erb and html tags - gross - I use a static generator to avoid that!).
- you put your head into your hands and quietly weep.

## How is go-staticgen different?
The code is small and simple and designed to be easy to edit. It's strongly opinionated, following a simple directory structure. You're free to enforce your own opinions as you see fit.

Just a warning though, I've only added things that golang currently supports natively. There's an example of libsass support in the project, but it is an external dependency so slows compile times and is disabled.

### Features
- manage all (or just one if you like) of your static sites in one clean directory structure, with a shared directory for code shared across all of them.
- build everything in one go 
- run a single server to host all sites at once (though configuration overrides via optional yaml file are coming)
- live reload

#### Wishlist
- error messages shown in the browser (up next)
- strong compile time error checking (coming soon)
- strong shared asset support (coming soon)
- markdown partial support (coming soon)
- more examples of extensibility (so you can easily add a garbage blog if you really need to)

### Technical Features
- compiles to single binary with no external dependencies
- can create its own directory structure from anywhere by typing `go-staticgen new <sitename>`
- one template language (ace)[https://github.com/yosssi/ace] and one stylesheet language (gcss)[https://github.com/yosssi/gcss]
- one general purpose preprocessor (templates can easily be passed through template/text so you can add your own functions to markup easily)

## Usage

Sites are organised with a directory structure as follows:
```
public                  - your built files
sites                   - your source files
sites/_shared           - anything you want to share between projects
sites/example           - an example site
sites/example/images    - where the processor will assume images are
sites/example/scripts   - where the processor will assume scripts are
sites/example/styles    - where the processor will assume stylesheets are
```

The idea is that all your sites sit directly (no subdirs) within `/sites`, any directories starting with a `.` or `_` will be ignored. These directories are expected to have certain subdirectories (eg. `styles`, `images`, `scripts`, `fonts`, and `files`) which are processed accordingly.

## Compiling
```bash
# an example of compiling, set GOPATH however you prefer.
export GOBIN=$PWD/bin
export GOPATH=$PWD
go build -v gen
mkdir ./bin
mv ./gen ./bin/gen
# executable will be placed in bin/gen, you should copy it to /bin or within your path.
```
