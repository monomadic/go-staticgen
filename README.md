# go-staticgen
A simpler re-thinking of static site generators in golang.

## OMFG ANOTHER STATIC SITE GENERATOR?
Wow. Calm down. After trying many static generators I realised there's a core problem with all of them - they focus on non-technical users who want to make blogs. Nothing against this kind of user, but it's not me, and it's not many people. There are designers out there who can code and vice versa, and static site generators don't cater to them at all.

## How is go-staticgen different?
The code is organised in a clean, extremely simple way to understand. It is built with the intention of being able to read and modify it to suit your own needs, but by default has a lot of convention-over-configuration and is very strongly opinionated. You're free to enforce your own opinions as you see fit. Namely:

### Features
- manage all (or just one if you like) of your static sites in one clean directory structure, with a shared directory for code shared across all of them.
- build everything in one go 
- run a single server to host all sites at once (though configuration overrides via optional yaml file are coming)

### Technical Features
- compiles to single binary with no external dependencies
- can create its own directory structure from anywhere by typing `go-staticgen new <sitename>`
- one template language (ace) and one stylesheet language (gcss)
- one general purpose preprocessor (templates can easily be passed through template/text so you can add your own functions to markup easily)

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
