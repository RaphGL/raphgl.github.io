---
title: Stop (ab)using regex
description: Stop it, get some help!
date: 2025-11-23
---

Due to how widespread support for regex is in many languages, people feel like regex is a good tool to reach out for when you need to match over some text.
Whether it be while writing a parser, validating some input or whatever else. I'm here to argue that while there are some small valid uses for regex that are very useful,
most uses of regex are an abuse of them. This is because people writing regex usually don't really know how to parse themselves.
So they just use whatever tool they know how to use.

### How your regex is ran

To be run, your regex first goes through a few steps:

1. Tokenizing - the regex is read and converted into tokens
2. Parsing - the tokens are converted into a syntax tree that can be walked through and more easily worked with by code
3. Compilation - the syntax tree is read and compiled into instructions for a regex virtual machine
4. Optimization (optional) - the instructions are ran through an optimizer and then turned into more efficient code

By the time all of this is done your program has run through a couple thousand lines of code.
Now you have a compiled regex expression and can run it through the regex virtual machine.

## Regex is bad

From a performance standpoint, doing all those steps already is much slower than calling any small function. Running through a virtual machine means at least a few levels
of indirection which is likely to thrash the cache and just runs a lot more code to do the same thing. So just by using regex, you're letting go of performance.
But ok, you don't care about performance. I get it, no one cares about performance anymore. So let's look at the other reason why regex is bad.

#### Regex is unmaintainable

Comprehending any reasonable complex regex expression is a huge waste of time. Let's look at this example
I copied [from stackoverflow](https://stackoverflow.com/questions/19605150/regex-for-password-must-contain-at-least-eight-characters-at-least-one-number-a):
```js
"^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$"
```

Can you guess what it does at a glance without opening the link, seeing a comment or using one of those regex explainer tools? Probably not.
Here's what it does:

> Minimum eight characters, at least one letter and one number

Now look at what the alternative is:
```ts
// note: isAlphanumeric, isNumber and isAlpha are not real javascript functions
// but they're here to prove a point not to write absolutely correct code
function isValidPassword(password: string): bool {
  if (password.length < 8) {
    return false;
  }

  let charCount = 0;
  let numCount = 0;

  for (const c of password) {
    if (!c.isAlphanumeric()) {
      return false;
    }

    if (c.isNumber()) numCount += 1;
    if (c.isAlpha()) charCount += 1;
  }

  return charCount >= 1 && numCount >= 1;
}
```
It is more "verbose" but you can actually understand what it's doing and if you wanted to update the validation you can easily change it and reason about it.
Meanwhile for the regex version it might take you a couple tries and you might even need to completely scrap your current implementation. Even if you were a regex wizard,
other people aren't. So your regex is just objectively harder to understand and maintain.

#### Regex can't give good error messages

The only thing you get from using a regex is a match or not a match. You can't really give good error messages with it.
If you write the parser yourself and you know what is and isn't expected, you can give more meaningful error messages than "password doesn't match regex" or whatever.

#### Regex is good actually

Regex is not a good tool for validation and parsing. People use them as such but they're not a good tool for that. It's not a replacement for knowing how to parse.
If your program requires user interactivity to do matches, then regex is a good tool. Things like search functionality in text editors,
the find tool on browsers and so on are good use cases, for a regex engine in your program.

But if you already know the shape of the thing you want to match against and it's not supplied by the user,
then you're probably better off writing your own validation/parser.
You'll get better performance, better errors and more understandable and maintainable code.
