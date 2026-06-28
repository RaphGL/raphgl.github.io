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
Con: Everything is wrapped in a macro and might break autocomplete

\
\

The approach I recommend is not really novel, but it's the one I've found to be the nicest to use and to work better with existing tooling.
Since it doesn't rely on type erasure and everything isn't wrapped in macros.

\
\

Pros:

- Type safe
- Doesn't have to use pointers
- Minimal use of Macros
- Can be deduplicated by linkers
- Can be entirely put in a single file to reduce changes triggering compilation for multiple files

Cons:

- Relies on a macro trick
- Can be a bit verbose

## Achieving generics through header instantiation

The way this works is, you instantiate a specialization for a type by defining your type and possibly a suffix if that type is not a valid identifier name.
To do this, we'll have to rely on a small macro trick, but other than that, everything is quite straight forward. The header instantiation should look like this:

```c
#define VEC_ITEM_TYPE long long
#define VEC_SUFFIX num
#include "vector.h"
```

With the `VEC_SUFFIX` being optional.

Once instantiated this would give you a `vector_push_num` function you could use.
To do this we need to be able to append `_num` to our symbol names. So we create a `G` function-like macro.

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

Now we can emit an error if `VEC_ITEM_TYPE` is not defined.

This fixes wrong uses of the header but what if we want to use something like `long long` or `atomic int`? It won't work.
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
bool G(vec_pop)(G(Vector) *vec, VEC_ITEM_TYPE *dest) {
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

As you can see, every time we call anything generic we need to wrap it in `G` so that it can properly call the generic function.

### The elephants in the room
#### Redeclaration error

You can try and use this implementation but unless you only have one C file and nothing else, you're likely to eventually run into a redeclaration error especially if you're
using this generic library on `.c` and `.h` files.

A potential fix is to make everything `static` but then you would end up duplicating everything throughout many translation units in your program.

The actual fix is to be able to choose when to forward declare and when to use the implementation:

```c
#ifndef VEC_IMPLEMENTATION
  bool G(vec_pop)(G(Vector) *vec, VEC_ITEM_TYPE *dest);
#else
  bool G(vec_pop)(G(Vector) *vec, VEC_ITEM_TYPE *dest) {
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
And then on your C files you can tell it that you need the implementations, not the declarations.
```c
#define VEC_IMPLEMENTATION
#define VEC_ITEM_TYPE int
#include "vector.h"
```

--- 

Side Note: if you include the header in another header, you can still just call `#define VEC_IMPLEMENTATION` without reincluding the generic header library to get the implementation.
Which is a nice side effect of this approach to generics, but it can end up looking a bit implicit because the include is hidden in another include.

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
This way we still get a "VEC_ITEM_TYPE is not defined" error if we include the vector again and forget to tell it what type we want instantiated.

NOTE: We don't undefine `VEC_IMPLEMENTATION` so that we can just declare it once in our `.c` files for convenience's sake.

#### Include guards
You might be asking about the missing include guard, but you don't need it! The reason is that we actually want to be able to include the same header multiple times.
Each time with a different type.

You might still want to have an include guard for a section of the header if you want to declare some types or code
that is not generic in the header. A good example of this would be if your allocator would receive an allocator context
which does not care about types. For example, you might want to have something like this:

```c
#ifndef VECTOR_H
#define VECTOR_H

typedef struct vec_allocator {
   void *(*malloc)(size_t);
   void (*free)(void *);
} VecAllocator;

#endif
```

That way, every single vector could still share the same allocator type.

## Putting everything together
The final header should look something like this:

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

#ifndef VEC_IMPLEMENTATION

typedef struct G(vec_vector) {
   size_t capacity;
   size_t len;
   VEC_ITEM_TYPE *vec;
} G(Vector);

G(Vector) * G(vec_new)(void);
bool G(vec_push)(G(Vector) *vec, VEC_ITEM_TYPE item);

#else

// Initializes a new vector with items of sizeof(T)
G(Vector) * G(vec_new)(void) {
   G(Vector) *vector = malloc(sizeof(*vector));
   if (!vector) {
      return NULL;
   }

   *vector = (G(Vector)) {
      .capacity = 0,
      .len = 0,
      .vec = NULL,
   };

   return vector;
}

// Resizes vector to fit length
bool G(vec_fit)(G(Vector) *vec) {
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
bool G(vec_push)(G(Vector) *vec, VEC_ITEM_TYPE item) {
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

