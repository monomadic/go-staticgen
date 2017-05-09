<img src="logo.png" alt="go-staticgen" height="75">

A simpler re-thinking of static site generators in golang.

- doesn't treat you like a wordpress template blogger - site development is first-class, rather than markdown and templates.
- livereload with errors shown in-browser
- simple plugin system, with well-written best practice golang, so it's easy to extend
  - [libsass](https://github.com/wellington/go-libsass) for css
  - [ace](https://github.com/yosssi/ace) for html
- compiles to single binary with no external dependencies
- manage multiple static sites at once, with shared code/assets
- extremely fast
- simple, easy to understand and extend code
- error checking as a fundamental principle - you will immediately know if any of your sites have bugs.

If you're a developer/designer who manages a group of static sites, and you want a better workflow, this might be a good option for you, as it revolves around the idea of code/module re-use across static sites. You could build up a 'reserve' of assets such as fonts, placeholder images, regularly used styles and framework files that are always available and you can update or manage them all in one simple directory structure without crazy dependencies or messing with the black hole of stupid that is node.js.

It becomes so easy and frictionless to whip up new sites, and keep them up to date, you'll find yourself making more that you otherwise would.

![screenshot](screenshot.png)

## OMFG ANOTHER STATIC SITE GENERATOR?
Wow. Calm down. After trying many static generators I realised there's a core problem with seemingly all of them - they focus on non-technical users who want to make blogs. Nothing against this kind of user, but it's not me. I know how to use a variety of front end languages, why can't I have the freedom to create sites using them without delving into templates? I want to make actual websites that have varied structure. I don't need the program to force me into a structure I dont want to work in, which usually is:

- you are told to write markdown files as blog posts, as a first priority.
- the design itself is done as a 'template', separate to your site, why?!
- if don't like the way the tool works, changing it requires reading some huge documentation library and it probably doesn't support it because you aren't the target market, being all smart with your actual development skills and all.

## Usage

```
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

## Folder Structure
```
public        # compiled files get generated into public
sites         # source files here
  _shared     # shared resources go here.
    images
    pages
  site1
    images
    styles
    pages
    scripts
    files
  site2
    images
    styles
    pages
```

The idea is that all your sites sit directly (no subdirs) within `/sites`, any directories starting with a `.` or `_` will be ignored. These directories are expected to have certain subdirectories (eg. `styles`, `images`, `scripts`, `fonts`, and `files`) which are processed accordingly.

## Installation (MacOS)
Grab the latest release on the [releases](releases) page. Unzip and add to your path. Or, from a terminal:

```bash
wget https://github.com/robsaunders/go-staticgen/releases/download/0.4/gen-0.4-mac.zip
unzip gen-0.4-mac.zip
chmod 777 gen
sudo mv gen /usr/local/bin
```

## Compilation

*Important!!!!* make sure you don't forget the linker flags, otherwise libsass *will* cause a segfault. Yes, the go linker is a bit shit.

```bash
# use gb tool, see: getgb.io
git clone https://github.com/robsaunders/go-staticgen.git
cd go-staticgen
gb vendor restore
gb build -ldflags -s
# executable will be placed in bin/gen, you should copy it to /bin or within your path.
```

## Templates

### Images, fonts and files (the copy helper)
Static binary files are treated a little differently in this processor. You need to use a template helper to copy them across to your project. This is intentional for these reasons:
- The destination directory will only contain used files and not copy unintended or unused files.
- If a file does not exist or gets renamed / corrupted, the compiler will pick up on that and tell you. You won't launch with broken images!
- Allows for cache busting on file updates. Super important for static site hosting!

To use these templates, use the following helper:
```haml
link href="{{ copy "styles/pure-min.css" }}" rel="stylesheet" type="text/css"
```

The files can exist in your `_shared` directory or your local directory.
