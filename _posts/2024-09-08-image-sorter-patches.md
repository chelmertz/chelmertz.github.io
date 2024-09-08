---
layout: post
title: "image-sorter patches"
date: 2024-09-08 15:37

tags: rust
comments: true
published: true
---

[image-sorter](https://github.com/jgalat/image-sorter/) is a smart program that
lets you organize your images.

You give it a couple of directories to look for images in, and some alternative
directories those images should move to. When viewing a "to be moved" image
inside the program window, you press a shortcut for the best matching output
directory, and move on to the next image.

Besides having a smart way of configuring the bindings to the output
directories, it also persists the "move" commands in a shell file, for you to
inspect before executing it. This is [dry
run](https://en.wikipedia.org/wiki/Dry_run_(testing)) being builtin and the
default, and it's so smart to me.

I tried adding a couple of patches, and luckily for me, the author was very
friendly and accepted them:

- [fix: Guarantee key mapping sort order](https://github.com/jgalat/image-sorter/commit/40ac1649abee3384c9149de8e694d8210c391ba2)
- [fix(path): Expand ~ (tilde) in target paths](https://github.com/jgalat/image-sorter/commit/0b79ab2051db057ff1433b87cce3c5c8e23ad972)
- [feat(delete): Delete an image with backspace](https://github.com/jgalat/image-sorter/commit/d72c98decb7b52f299d8ddd64980fe5f2b832c64)
- [feat: Recurse into target paths if -r is given](https://github.com/jgalat/image-sorter/commit/ccfb588400386fbde6a70bbd38f77ce23e55f579)
- [chore: Decorate errors with useful information](https://github.com/jgalat/image-sorter/commit/1ac60977a2a1444d66c439766dfb723398fc6318)
- [fix(filetype): infer after checking file ext](https://github.com/jgalat/image-sorter/commit/1532584003acce3d30f4f6f8034bb6debae60c03)

Besides some very small contributions at work, this is the first patches I've
written in Rust. Last time I tried writing rust with vscode-out-of-the-box
settings, the process was very slow (auto completion and type checking took
very much time, not ideal when coding). For this size of project, the tooling
was super quick, so I could just enjoy the nice things about Rust.

Some things to keep in mind for the next excursion with Rust:

- Don't overuse the catch-all `_` in pattern matching. When adding new (for
  example) enum cases, you want to get compiler errors to fix.
- Use linters. `cargo fmt --all -- --check` and `cargo clippy --all-targets
  --all-features -- -D warnings` were used in this project, and they offer good
hints à la "yes, your code compiles but _this_ is more semantic".
- [structopt](https://docs.rs/structopt/latest/structopt/) was nice for CLI
  parsing, I'm sure [clap](https://docs.rs/clap/latest/clap/) is good, too.
- The dependencies in rust are so many and so small (and I added more of them
  🙈).

------

And also, a debugging experience:

I wanted to add a feature for looking through folders recursively for images.
Once the code was written, and I tested it, the program just died. This was due
to the alignment of some nice error conditions:

1. "Is it an image or not" was checked by file type and nothing else.
2. The directory being looked through, contained node_modules somewhere in the
   file tree. One of those node_modules contained an invalid image file
   (basically `echo hello>a.jpg`) as a test case. Because, of course, having a
   dependency in node means importing the whole repository.
3. The TUI library code swallowed error messages, in order to control what's
   rendered by the client.

