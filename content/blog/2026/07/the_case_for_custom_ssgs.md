---
title: The case for custom static site generators
description: Some reasons to ditch static site generators and write your own
date: 2026-07-04
---

Static site generators (SSGs) are fairly popular. Some of the most common ones out there are:
[Hugo](https://gohugo.io/), [Jekyll](https://jekyllrb.com/), [Gatsby](https://www.gatsbyjs.com/), [Zola](https://www.getzola.org/) 

These generators usually rely on some form of directory structure to make sense of your website.
They also try to be very generic and provide as much functionality as possible out of the box.
In this quest for broad utility they all end up making generating a static site more complicated than it should be.
There's many functions, many taxonomies, you need to learn *their* templating language (or use a javascript frontend framework which is then prerendered),
know which block thingy to wrap your HTML around, etc.

This makes using many (most?) static site generators complicated..
All you wanted to do was render some HTML and display some information in the exact way you wanted it, but now you have to debug your templates.
These tools also often break APIs behind your back so you have to keep up with the changes.

## You do not need a general purpose static site generator

Static site generators are glorified string appenders. You can write your own in half an hour and just add more functionality as you need them. 

The initial set up is slightly annoying but once you have a base to build on, it gets much easier.
You don't have to keep up with breaking changes anymore and you can implement whatever you want in whatever way you want to very easily
since the code base is yours and you completely understand it.
