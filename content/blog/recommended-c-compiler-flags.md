---
title: C compiler flags you ought to know
description: A compiled list of compiler flags that make development and catching errors easier
date: 2025-11-17
---

C compilers have bad defaults due to the need to be "backwards compatible".
This means that despite improvements in C compilers, you just have to know what to use.
If you're not in the know, you might subject yourself to pain needlessly.

This is a compiled list of flags you might want to use.
Heavily inspired by [Compiler Options Hardening Guide for C and C++](https://best.openssf.org/Compiler-Hardening-Guides/Compiler-Options-Hardening-Guide-for-C-and-C++.html).
I felt that this adaptation was needed because that post suffered from "too much yapping".

## The Essentials

No matter what you do, always enable `-Wall -Wextra -pedantic`, these flags are actually a group of other flags that enable other errors.
If you want to see what they can do you can check the manpage for them:

```
-Wall
  This enables all the warnings about constructions that some users consider questionable, and that are easy to avoid (or modify to prevent the warning), even in conjunction with
  macros.  This also enables some language-specific warnings described in C++ Dialect Options and Objective-C and Objective-C++ Dialect Options.

  -Wall turns on the following warning flags:

  -Waddress -Waligned-new (C++ and Objective-C++ only) -Warray-bounds=1 (only with -O2) -Warray-compare -Warray-parameter=2 -Wbool-compare -Wbool-operation -Wc++11-compat
  -Wc++14-compat  -Wc++17compat  -Wc++20compat -Wcatch-value (C++ and Objective-C++ only) -Wchar-subscripts -Wclass-memaccess (C++ and Objective-C++ only) -Wcomment -Wdangling-else
  -Wdangling-pointer=2 -Wdelete-non-virtual-dtor (C++ and Objective-C++ only) -Wduplicate-decl-specifier (C and Objective-C only) -Wenum-compare (in C/ObjC; this is on by default in
  C++) -Wenum-int-mismatch (C and Objective-C only) -Wformat=1 -Wformat-contains-nul -Wformat-diag -Wformat-extra-args -Wformat-overflow=1 -Wformat-truncation=1 -Wformat-zero-length
  -Wframe-address -Wimplicit (C and Objective-C only) -Wimplicit-function-declaration (C and Objective-C only) -Wimplicit-int (C and Objective-C only) -Winfinite-recursion
  -Winit-self (C++ and Objective-C++ only) -Wint-in-bool-context -Wlogical-not-parentheses -Wmain (only for C/ObjC and unless -ffreestanding) -Wmaybe-uninitialized -Wmemset-elt-size
  -Wmemset-transposed-args -Wmisleading-indentation (only for C/C++) -Wmismatched-dealloc -Wmismatched-new-delete (C++ and Objective-C++ only) -Wmissing-attributes -Wmissing-braces
  (only for C/ObjC) -Wmultistatement-macros -Wnarrowing  (C++ and Objective-C++ only) -Wnonnull -Wnonnull-compare -Wopenmp-simd (C and C++ only) -Woverloaded-virtual=1 (C++ and
  Objective-C++ only) -Wpacked-not-aligned -Wparentheses -Wpessimizing-move (C++ and Objective-C++ only) -Wpointer-sign (only for C/ObjC) -Wrange-loop-construct (C++ and
  Objective-C++ only) -Wreorder (C++ and Objective-C++ only) -Wrestrict -Wreturn-type -Wself-move (C++ and Objective-C++ only) -Wsequence-point -Wsign-compare (C++ and Objective-C++
  only) -Wsizeof-array-div -Wsizeof-pointer-div -Wsizeof-pointer-memaccess -Wstrict-aliasing -Wstrict-overflow=1 -Wswitch -Wtautological-compare -Wtrigraphs -Wuninitialized
  -Wunknown-pragmas -Wunused -Wunused-but-set-variable -Wunused-const-variable=1 (only for C/ObjC) -Wunused-function -Wunused-label -Wunused-local-typedefs -Wunused-value
  -Wunused-variable -Wuse-after-free=2 -Wvla-parameter -Wvolatile-register-var -Wzero-length-bounds

  Note that some warning flags are not implied by -Wall.  Some of them warn about constructions that users generally do not consider questionable, but which occasionally you might
  wish to check for; others warn about constructions that are necessary or hard to avoid in some cases, and there is no simple way to modify the code to suppress the warning. Some of
  them are enabled by -Wextra but many of them must be enabled individually.
```

```
-Wextra
  This enables some extra warning flags that are not enabled by -Wall. (This option used to be called -W.  The older name is still supported, but the newer name is more descriptive.)

  -Wabsolute-value (only for C/ObjC) -Walloc-size -Wcalloc-transposed-args -Wcast-function-type -Wclobbered -Wdangling-reference (C++ only) -Wdeprecated-copy (C++ and Objective-C++
  only) -Wempty-body -Wenum-conversion (only for C/ObjC) -Wexpansion-to-defined -Wignored-qualifiers  (only for C/C++) -Wimplicit-fallthrough=3 -Wmaybe-uninitialized
  -Wmissing-field-initializers -Wmissing-parameter-name (C/ObjC only) -Wmissing-parameter-type (C/ObjC only) -Wold-style-declaration (C/ObjC only) -Woverride-init (C/ObjC only)
  -Wredundant-move (C++ and Objective-C++ only) -Wshift-negative-value (in C++11 to C++17 and in C99 and newer) -Wsign-compare (C++ and Objective-C++ only) -Wsized-deallocation (C++
  and Objective-C++ only) -Wstring-compare -Wtype-limits -Wuninitialized -Wunterminated-string-initialization -Wunused-parameter (only with -Wunused or -Wall)
  -Wunused-but-set-parameter (only with -Wunused or -Wall)

  The option -Wextra also prints warning messages for the following cases:

  *   A pointer is compared against integer zero with "<", "<=", ">", or ">=".

  *   (C++ only) An enumerator and a non-enumerator both appear in a conditional expression.

  *   (C++ only) Ambiguous virtual bases.

  *   (C++ only) Subscripting an array that has been declared "register".

  *   (C++ only) Taking the address of a variable that has been declared "register".

  *   (C++ only) A base class is not initialized in the copy constructor of a derived class.
```

```
-pedantic
  Issue all the warnings demanded by strict ISO C and ISO C++; diagnose all programs that use forbidden extensions, and some other programs that do not follow ISO C and ISO C++.
  This follows the version of the ISO C or C++ standard specified by any -std option used.
  
  Valid ISO C and ISO C++ programs should compile properly with or without this option (though a rare few require -ansi or a -std option specifying the version of the standard).
  However, without this option, certain GNU extensions and traditional C and C++ features are supported as well.  With this option, they are diagnosed (or rejected with
  -pedantic-errors).
  
  -Wpedantic does not cause warning messages for use of the alternate keywords whose names begin and end with __.  This alternate format can also be used to disable warnings for
  non-ISO __intN types, i.e. __intN__.  Pedantic warnings are also disabled in the expression that follows "__extension__".  However, only system header files should use these escape
  routes; application programs should avoid them.
  
  Some warnings about non-conforming programs are controlled by options other than -Wpedantic; in many cases they are implied by -Wpedantic but can be disabled separately by their
  specific option, e.g. -Wpedantic -Wno-pointer-sign.
  
  Where the standard specified with -std represents a GNU extended dialect of C, such as gnu90 or gnu99, there is a corresponding base standard, the version of ISO C on which the GNU
  extended dialect is based.  Warnings from -Wpedantic are given where they are required by the base standard.  (It does not make sense for such warnings to be given only for
  features not in the specified GNU C dialect, since by definition the GNU dialects of C include all features the compiler supports with the given option, and there would be nothing
  to warn about.)
```

Some people recommend adding `-Werror` which turns every warning into an error, but I don't really think that's reasonable. When you're iterating on a project, being forced to 
fix any minutia slows you down. If you're going to enable it at all, I suggest you enable it for release builds.

#### Sanitizers
Sanitizers are your friend. They'll help you with detecting memory leaks, undefined behavior and other common bugs. They can be enabled with the `-fsanitize` flag.
Also don't forget to add `-g` when compiling to add debug symbols, otherwise the sanitization errors might not be as useful.
(You can also use valgrind but I won't get into it in this post).

There's many sanitizers you can use. I recommend looking at the manpage and seeing what sanitizers are available. Bear in mind that not all sanitizers can be used together.
So you might have to compile your code and test it with each sanitizer individually. As far as I know, you can use `-fsanitize=leak,undefined,address` and these 3 can work together just fine.
So I recommend using this at least on your debug builds. (do not that if you use a library that has many leaks and was never used with sanitizers, you might have to add an ignore list so that you
can sanitize your own code, it's just the state of things sadly)

#### C Standard
A lot of people swear by C89 but personally I'm not a fan. I recommend C99+. C99 is a nice standard if you want to support as many compilers as possible and/or you don't need atomics.
If you do need atomics, C11 is fine, Linux for example defaults to it (technically gnu11, but still).

I recommend you always set your standard with `-std=c99` or whatever version you're using. GCC defaults to their own dialect which has a bunch of extensions. I'd rather use the standard
and not a dialect, only use it if you know you need the extensions (for some reason).

## Compiler Flags
The following are flags that you might benefit from, they either add compile-time or runtime checks to make the compiler/program even stricter.

Flags that enable compile-time checks:

|Compiler Flag                                                                         |Supported since      |Description                                                                                                                               |
|---                                                                                   |---                  |---                                                                                                                                       |
| -Wformat <br/>-Wformat=2                                                             |GCC 2.95.3Clang 4.0.0|Enable additional format function warnings                                                                                                |
| -Wconversion <br/>-Wsign-conversion                                                  |GCC 2.95.3Clang 4.0.0|Enable implicit conversion warnings                                                                                                       |
| -Wtrampolines                                                                        |GCC 4.3.0            |Enable warnings about trampolines that require executable stacks                                                                          |
| -Wimplicit-fallthrough                                                               |GCC 7.0.0Clang 4.0.0 |Warn when a switch case falls through                                                                                                     |
| -Wbidi-chars=any                                                                     |GCC 12.0.0           |Enable warnings for possibly misleading Unicode bidirectional control characters                                                          |
| -Werror <br/>-Werror=<warning-flag>                                                  |GCC 2.95.3Clang 2.6.0|Treat all or selected compiler warnings as errors. Use the blanket form  <br/>-Werror only during development, not in source distribution.|
| -Werror=format-security                                                              |GCC 2.95.3Clang 4.0.0|Treat format strings that are not string literals and used without arguments as errors                                                    |
| -Werror=implicit <br/>-Werror=incompatible-pointer-types <br/>-Werror=int-conversion |GCC 2.95.3Clang 2.6.0|Treat obsolete C constructs as errors                                                                                                     |


Flags that enable runtime checks:
|Compiler Flag                                              |Supported since            |Description                                                                                                                                                                                                            |
|---                                                        |---                        |---                                                                                                                                                                                                                    |
| -D_FORTIFY_SOURCE=3                                       |GCC 12.0.0Clang 9.0.02     |Fortify sources with compile- and run-time checks for unsafe libc usage and buffer overflows. Some fortification levels can impact performance. Requires -O1 or higher, may require prepending -U_FORTIFY_SOURCE.      |
| -D_GLIBCXX_ASSERTIONS                                     |libstdc++ 6.0.0            |Precondition checks for C++ standard library calls. Can impact performance.                                                                                                                                            |
| -fstrict-flex-arrays=3                                    |GCC 13.0.0Clang 16.0.0     |Consider a trailing array in a struct as a flexible array if declared as []                                                                                                                                            |
| -fstack-clash-protection                                  |GCC 8.0.0Clang 11.0.0      |Enable run-time checks for variable-size stack allocation validity. Can impact performance.                                                                                                                            |
| -fstack-protector-strong                                  |GCC 4.9.0Clang 6.0.0       |Enable run-time checks for stack-based buffer overflows. Can impact performance.                                                                                                                                       |
| -fcf-protection=full                                      |GCC 8.0.0Clang 7.0.0       |Enable control-flow protection against return-oriented programming (ROP) and jump-oriented programming (JOP) attacks on x86_64                                                                                         |
| -mbranch-protection=standard                              |GCC 9.0.0Clang 8.0.0       |Enable branch protection against ROP and JOP attacks on AArch64                                                                                                                                                        |
|  <br/>-Wl,-z,nodlopen                                     |Binutils 2.10.0            |Restrict dlopen(3) calls to shared objects                                                                                                                                                                             |
|  <br/>-Wl,-z,noexecstack                                  |Binutils 2.14.0            |Enable data execution prevention by marking stack memory as non-executable                                                                                                                                             |
|  <br/>-Wl,-z,relro <br/>-Wl,-z,now                        |Binutils 2.15.0            |Mark relocation table entries resolved at load-time as read-only.  <br/>-Wl,-z,now can impact startup performance.                                                                                                     |
| -fPIE -pie                                                |Binutils 2.16.0Clang 5.0.0 |Build as position-independent executable. Can impact performance on 32-bit architectures.                                                                                                                              |
| -fPIC -shared                                             |< Binutils 2.6.0Clang 5.0.0|Build as position-independent code. Can impact performance on 32-bit architectures.                                                                                                                                    |
| -fno-delete-null-pointer-checks                           |GCC 3.0.0Clang 7.0.0       |Force retention of null pointer checks                                                                                                                                                                                 |
| -fno-strict-overflow                                      |GCC 4.2.0                  |Define behavior for signed integer and pointer arithmetic overflows                                                                                                                                                    |
| -fno-strict-aliasing                                      |GCC 2.95.3Clang 2.9.0      |Do not assume strict aliasing                                                                                                                                                                                          |
| -ftrivial-auto-var-init                                   |GCC 12.0.0Clang 8.0.0      |Initialize automatic variables that lack explicit initializers                                                                                                                                                         |
| -fexceptions                                              |GCC 2.95.3Clang 2.6.0      |Enable exception propagation to harden multi-threaded C code                                                                                                                                                           |
| -fhardened                                                |GCC 14.0.0                 |Enable pre-determined set of hardening options in GCC                                                                                                                                                                  |
|  <br/>-Wl,--as-needed <br/>-Wl,--no-copy-dt-needed-entries|Binutils 2.20.0            |Allow linker to omit libraries specified on the command line to link against if they are not used                                                                                                                      |
| -fzero-init-padding-bits=all                              |GCC 15.0.0                 |Guarantee zero initialization of padding bits in all automatic variable initializers                                                                                                                                   |

I suggest you read the manpages and search online for the flags you're interested in using.
