---
title: How to write generics in C 
description: Implementing type-safe generics in C step by step
date: 2025-11-15
---

C is barebones and "doesn't support" generics, but it's actually quite easy to implement with the tools we already have. 
There's many ways you might find them being implemented in the wild. Some of the common ones are:

- **Using function-like macros**

```c
#define vector_push(vector, item) vector.buf[vector.idx++] = item;
```
Con: This will cause everything to be inlined and loosely typed.

- **Using void pointers**

```c
void vector_push(Vector vec, void *item);
```
Con: You rely on type erasure to achieve generics, so you'll have to recast everything back to access them.
You might also run into UB.

- **Dispatch specializations through a function-like macro**

```c
#define DECLARE_VECTOR(type) \
  void vector_push(Vector##_##type, type item) {\
    // do stuff here \
    }

DECLARE_VECTOR(int)
```

\

The approach I recommend is not really novel, but it's the one I've found to be the nicest to use and to work better with existing tooling.
Since it doesn't rely on type erasure and everything isn't wrapped in macros.

## Achieving generics through header instantiation

The way this works is, you instantiate a specialization for a type by defining your type and possibly a suffix if the that type is not a valid name for an identifier.
To do this, we'll have to rely on a small macro trick, but other than that, everything is quite straight forward. The header instantiation should look like this:

```c
#define VEC_ITEM_TYPE long long
#define VEC_SUFFIX num
#include "vector.h"
```

With the `VEC_SUFFIX` being optional.

Once instantiated this would give you a `vector_push_num` function you could use.
To do this we need to be able to append `_num` to our symbol names. For we create a `G` function-like macro.

```c
#define G(name) name##_##VEC_ITEM_TYPE
```

Except, this doesn't work. You can't just expand and append them like that. We need to expand the definitions first and then append them, so we'll rely on the following trick:

```c
#define __G(name, type) name##_##type
#define _G(name, type) __G(name, type)
#define G(name) _G(name, VEC_ITEM_TYPE)
```

It works now! Now we can wrap our identifier names in a `G` macro so that they're [*name mangled*](https://en.wikipedia.org/wiki/Name_mangling) for the types they're for.  
We can use it to define our functions and structs like this: `struct G(something)`, `void G(do_something)()`.

### Enforcing type definition and allowing custom suffixes

```c
#ifndef VEC_ITEM_TYPE
  #error VEC_ITEM_TYPE was not defined
#endif

#ifndef VEC_SUFFIX
  #define VEC_SUFFIX VEC_ITEM_TYPE
#endif
```

Now we can emit an error if `VEC_ITEM_TYPE` and we know that if this header is compiled it means we have a `VEC_ITEM_TYPE` defined.

This fixes wrong use of the header but what if we want to use something like `long long` or `atomic int`? It won't work.
So we have to introduce a way of programmatically overriding the suffix that's added to functions. So we'll introduce a `VEC_SUFFIX`:

```c
#ifndef VEC_ITEM_TYPE
  #error VEC_ITEM_TYPE was not defined
#endif

#ifndef VEC_SUFFIX
  #define VEC_SUFFIX VEC_ITEM_TYPE
#endif

#define __G(name, type) name##_##type
#define _G(name, type) __G(name, type)
#define G(name) _G(name, VEC_SUFFIX)
```

Now we're free to implement our library. For example, here's a generic `vec_pop` implementation:

```c
bool G(vec_pop)(G(vec_Vector) *vec, VEC_ITEM_TYPE *dest) {
   if (!vec->vec || vec->len <= 0 || vec->capacity <= 0) {
      return false;
   }

   vec->len--;
   if (dest) {
      *dest = vec->vec[vec->len];
   }
   G(vec_fit)(vec);
   return true;
}
```

As you can see, every time we call anything generic we need to wrap it in `G` so that it can be specialized and name mangled properly.

### The elephants in the room
#### Redeclaration error

You can try and use this implementation but unless you only have one C file and nothing else, you're likely to eventually run into a redeclaration error especially if you're
using this generic library on `.c` and `.h` files.

A potential fix is to make everything `static` but then you would end up duplicating everything throughout many translation units in your program.

The actual fix is to be able to choose when to forward declare and when to use the implementation:

```c
#ifndef VEC_IMPLEMENTATION
  bool G(vec_pop)(G(vec_Vector) *vec, VEC_ITEM_TYPE *dest);
#else
  bool G(vec_pop)(G(vec_Vector) *vec, VEC_ITEM_TYPE *dest) {
     if (!vec->vec || vec->len <= 0 || vec->capacity <= 0) {
        return false;
     }

     vec->len--;
     if (dest) {
        *dest = vec->vec[vec->len];
     }
     G(vec_fit)(vec);
     return true;
  }
#endif
```

You can then just include the header in header files
```c
// included like any other header
#define VEC_ITEM_TYPE int
#include "vector.h"
#include <stdio.h>
#include <stdlib.h>
```
And then on your C files you can tell it that you need the implementations instantiated
```c
#define VEC_IMPLEMENTATION
#define VEC_ITEM_TYPE int
#include "vector.h"
```

--- 

Side Note: if you include the header in another header, you can still just call `#define VEC_IMPLEMENTION` without reincluding the file to get the implementation.
Which is a nice side effect of this method, but it can end up looking a bit implicit because the include is hidden in another include. It's up to you what you like.

Example:
```c
#define VEC_IMPLEMENTATION
#include "flag.h"
```

#### Multiple includes inheriting types
Since we're relying on preprocessor definitions to define our types.
If we do:
```c
#define VEC_ITEM_TYPE int
#include "vector.h"
#include "vector.h"
```
We'll have a redeclaration error, because both includes see `VEC_ITEM_TYPE` as `int`. We can fix it by undefining it at the end of our generic vector header:
```c
#undef VEC_ITEM_TYPE
#undef VEC_SUFFIX
```
This way we still get a "undeclared VEC_ITEM_TYPE" error if we include vector again and forget to tell it what type we want instantiated.

NOTE: We don't undefine `VEC_IMPLEMENTATION` so that we can just declare it once in our `.c` for convenience.

#### Include guards
You might be asking about the missing include guard, but you don't need it! The reason is that we actually want to be able to include the same header multiple times.
Each time for a different type.

## Putting everything together
The final header would look something like this:

```c
#ifndef VEC_ITEM_TYPE
   #error VEC_ITEM_TYPE was not defined
#endif

#ifndef VEC_SUFFIX
   #define VEC_SUFFIX VEC_ITEM_TYPE
#endif

#define __G(name, type) name##_##type
#define _G(name, type) __G(name, type)
#define G(name) _G(name, VEC_SUFFIX)

#include <math.h>
#include <stdbool.h>
#include <stdint.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

typedef struct G(vec_vector) {
   size_t capacity;
   size_t len;
   VEC_ITEM_TYPE *vec;
} G(vec_Vector);

#ifndef VEC_IMPLEMENTATION

G(vec_Vector) * G(vec_new)(void);
bool G(vec_push)(G(vec_Vector) *vec, VEC_ITEM_TYPE item);

#else

// Initializes a new vector with items of sizeof(T)
G(vec_Vector) * G(vec_new)(void) {
   G(vec_Vector) *vector = malloc(sizeof(*vector));
   if (!vector) {
      return NULL;
   }

   *vector = (G(vec_Vector)) {
      .capacity = 0,
      .len = 0,
      .vec = NULL,
   };

   return vector;
}

// Resizes vector to fit length
bool G(vec_fit)(G(vec_Vector) *vec) {
   const size_t power = ceilf(log2f(vec->len + 1));
   const size_t new_capacity = sizeof(vec->vec[0]) * powf(2, power);

   if (new_capacity < vec->capacity || new_capacity > vec->capacity) {
      void *tmp = realloc(vec->vec, new_capacity);
      if (!tmp) {
         return false;
      }
      vec->vec = tmp;
      vec->capacity = new_capacity;
   }
   return true;
}

// Pushes a value to vector
bool G(vec_push)(G(vec_Vector) *vec, VEC_ITEM_TYPE item) {
   if (vec->len == 0 && vec->capacity == 0) {
      vec->vec = malloc(sizeof(vec->vec[0]));
      if (!vec->vec) {
         return false;
      }

      vec->vec[0] = item;
      vec->len = 1;
      vec->capacity = sizeof(vec->vec[0]);
      return true;
   }

   G(vec_fit)(vec);
   vec->vec[vec->len] = item;
   vec->len++;
   return true;
}

#endif
#undef VEC_ITEM_TYPE
#undef VEC_SUFFIX
```

