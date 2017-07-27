<img src="logo.png" alt="go-staticgen" height="75">

A simpler re-thinking of static site generators in golang.

## Features
- [ace](https://github.com/yosssi/ace) for html (similar to haml)
- [libsass](https://github.com/wellington/go-libsass) for css
- livereload (use a plugin on the browser, eg [chrome](https://chrome.google.com/webstore/detail/livereload/jnihajbhpnppcggbcgedagnkighmdlei?hl=en))
- errors shown in-browser
- SCREAMINGLY FAST

## Pitch
This generator is unlike other generators. It's a tool designed for efficiency with zero bloat. No 'downloading templates', or  complicated documentation. Give it a chance, you will get the hang of this style of development within minutes, and never want to use another generator.

If you don't like the code-safety of haml-like or sass-like preprocessors or don't know how to write basic markup, this isn't for you. If you like wordpress, this isn't for you. If you're a developer/designer who manages a group of static sites, and you want a better workflow, this might be a good option for you. It revolves around the idea of code/module re-use with minimal fuss. You build up a 'reserve' of shared assets such as fonts, placeholder images, code chunks, regularly used styles and framework files that are always available and you can update or manage them all in one simple directory structure without crazy dependencies or messing with the black hole of stupid that is node.js.

It compiles to a single directory and tells you instantly if there are missing files, broken code, or filesystem problems across your entire set of sites, so you know exactly what's what.

It becomes so easy and frictionless to whip up new sites, and keep them up to date, you'll find yourself addicted to pushing changes and making new sites.

![screenshot](screenshot.png)

## OMFG ANOTHER STATIC SITE GENERATOR?
Wow. Calm down. After trying many static generators I realised there's a core problem with seemingly all of them - they focus on non-technical users who want to make blogs, and/or don't know how to write markup. Nothing against this kind of user, but it's not me. I know how to use a variety of front end languages, why can't I have the freedom to create sites using them without delving into templates? I want to make actual websites that have varied structure. I don't need the program to force me into a structure I dont want to work in, which usually is:

- you are told to write markdown files as blog posts, as a first priority.
- the design itself is done as a 'template', separate to your site, why?!
- if don't like the way the tool works, changing it requires reading some huge documentation library and it probably doesn't support it because you aren't the target market, being all smart with your actual development skills and all.

## Installation
For mac, grab the latest binary from the [releases](https://github.com/robsaunders/go-staticgen/releases) page.

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
