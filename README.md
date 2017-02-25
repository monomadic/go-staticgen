# go-staticgen
A simpler re-thinking of static site generators in golang.

## OMFG ANOTHER STATIC SITE GENERATOR?
Wow. Calm down. After trying many static generators I realised there's a core problem with all of them - they focus on non-technical users who want to make blogs. Nothing against this kind of user, but it's not me.

## How is go-staticgen different?
The code is small and simple and designed to be easy to edit. It's strongly opinionated, following a simple directory structure. You're free to enforce your own opinions as you see fit.

If you're a developer/designer who manages a group of static sites, and you want a better workflow, this might be a good option for you, as it revolves around the idea of code/module re-use across static sites. You could build up a 'reserve' of assets such as fonts, placeholder images, regularly used styles and framework files that are always available and you can update or manage them all in one simple directory structure without crazy dependencies or messing with the black hole of stupid that is node.js.

Just a warning though, I've only added things that golang currently supports natively. There's an example of libsass support in the project, but it is an external dependency so slows compile times and is disabled.

### Features
- manage all (or just one if you like) of your static sites in one clean directory structure, with a shared directory for code shared across all of them.
- build everything in one go 
- run a single server to host all sites at once (though configuration overrides via optional yaml file are coming)
- live reload (coming very soon)
- strong compile time error checking (coming soon)
- strong shared asset support (coming soon)

### Technical Features
- compiles to single binary with no external dependencies
- can create its own directory structure from anywhere by typing `go-staticgen new <sitename>`
- one template language (ace) and one stylesheet language (gcss)
- one general purpose preprocessor (templates can easily be passed through template/text so you can add your own functions to markup easily)

## Usage
For a usage guide simply run the command:
```bash
$ ./go-staticgen

  NAME:
    go-staticgen

  DESCRIPTION:
    An opinionated multi-site static generator written in golang.

  COMMANDS:
    new <sitename>       Creates a new site scaffolding.
    build <sitename>     Process a specific site only.
    build                Process all sites.
    serve                Serve your site locally.
```

Sites are organised into a directory structure as follows:
```
public                  - your built files
sites                   - your source files
sites/_shared           - anything you want to share between projects
sites/example           - an example site
sites/example/images    - where the processor will assume images are
sites/example/scripts   - where the processor will assume scripts are
sites/example/styles    - where the processor will assume stylesheets are
```

## Compiling
```bash
# make sure your gopath is set up correctly, or just use something like:
export GOPATH=$PWD
export PATH=$PWD/src:$PATH

# install dependencies
go get

# compile
go build
```
