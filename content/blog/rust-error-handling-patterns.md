---
title: Error handling patterns in Rust
description: Common error handling patterns that help reduce verbosity in Rust
date: 2025-01-09
---

These are some patterns I've found in Rust that tend to make error handling code more legible/less noisy.

## Handling errors
If you only care about one case in an error, instead of being forced to match against all errors, do this:
```rs
if let Some(val) = returns_option() {
    // do something
}
```
If the value that is returns doesn't matter, then you can do this:
```rs  
if returns_option().is_some() {
    // do something
}
```
There's corresponding versions for the other values as well: `Result::{is_err, is_ok}`, `Option::{is_some, is_none}`

If you expect the error to never happen, or if it happens you can't handle it, you can just `unwrap` or `expect` it.

- `unwrap` just panics over an error. So if you do this:
```rs
let x = returns_option().unwrap();
```
that would be equivalent to this:
```rs
let x = match returns_option() {
    Some(x) => x,
    None => panic!("some reason why it errored here"),
};
```

- `expect` allows you to change the panic message, you should write there what your expectation was for that error
```rs
let retcode = execute_file("path/to/file").expect("file should be executable");
```

## Propagating errors

```rs
fn func(x: &str) -> Result<&str, Error> {
    // instead of:
    let x = match returns_result(x) {
        Ok(other) => other,
        Err(err) => return err,
    };

    // do:
    let x = returns_result(x)?;
}
```

Sometimes though, the Result or Option returns by the function might not be 1-to-1 with the functions you're propagating the error from, in those cases
you can either use `let else` idiom:
```rs
let Ok(other) = returns_result(x) else {
    return None;
}
```

Or you can also convert from a Result to an Option or vice versa with these methods: `Result::{ok, err}` and `Option::ok_or` to convert to the appropriate type
```rs
fn func(x: &str) -> Option<&str> {
    let other = returns_result().ok()?;
}
```

If these are not sufficient, you might want to change the Error or Ok values so you can "massage" it into working,
 these functions are `Result::{map_err, map_ok}` and `Option::map_or`. You can use them like this:
```rs
let other = returns_result().map_err(|e| /* massage the value here */)?;
```

