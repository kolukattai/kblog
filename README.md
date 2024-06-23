
# KBlog
a simple static sight generator for blogging


## dependencies
- latest version of golang v1.22 or above is required

## getting started
use the following command to install kblog
```sh
    go install github.com/kolukattai/kblog@latest
```
to create a blog first create a empty folder and run the following command followed by your post title, post title can have space in between them
```sh
    kblog post my first post
```

this will create a project folder structure this
```sh
posts
    |-- my-first-post.md
config.yaml
```

if the above comment does not work then you should set golang path properly
You'll need to add `GOPATH/bin` to `PATH`

To set both, add this to your `.profile` file:

```
export GOPATH="$HOME/go"
PATH="$GOPATH/bin:$PATH"
```

## Features offers
- [x] Home page
- [x] Lazy loading post
- [x] Post page
- [x] Tag page
- [x] Category page
- [x] SocialMedia Share

- [x] Google Analytics
- [x] Design Revamp
- [x] Adding Robots.txt and sitemap files for build
- [x] faster build


## config.yaml file
config file is made in such a way that it is self explanatory, the only thing that is missing here is that the `public` folder will be the static assets folder, meaning all the static assets will be store inside this folder in order for the sight to access it

```yaml
name: "KBlog"
domainName: domain.com
logo: "/public/images/logo.png"
default:
    title: "KBlog"
    description: "this is webpage default description"
    keywords: "kblog, bloging"
    author: "you"
ga: ""
styles: []
scripts: []
perPage: 10
twitter: https://twitter.com/kolukattai
facebook: https://facebook.com/kolukattai
instagram: https://instagram.com/kolukattai
```

## Front Mater
the `yaml` data in the top of the markdown files is called front mater, it should be on top of the file and should be in between tripple eifen 

```yaml
---
title: Other stuff
description: this is post description
keywords: one, two, three
tags:
    - one
    - two
    - three
category: Technology
author: <your name>
landingImage: /static/images/placeholder.webp
date: Mon, 10 Jun 2024 00:51:57 IST
---

this is page content

```



