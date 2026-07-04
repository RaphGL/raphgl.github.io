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
know which block thingy to wrap your html around, etc.

All this rigmarole is simply a waste of time. All you wanted to do was render some html and display some information but now you're forced to understand all this machinery,
possibly debug why you can't get it to do exactly what you wanted, and they might've broken the API and now you're forced
to go and update your website's template just because of someone else's whims.  

## You do not need a general purpose static site generator

What you want is to be able to write some markdown or some other format of your choice and have it be translated to the verbose html equivalent with a little bit of styling.
Static site generators are glorified string appenders. You can write your own in half an hour and just add more functionality as you need them. 

The initial set up might be slightly annoying but once you have the base built, building on top of it is way easier.
There's no more arbitrary breaking changes because of things you never even cared about.
You no longer have to spend time digging documentation and forums for answers on how to do something, you can just do it. You're in total control.

