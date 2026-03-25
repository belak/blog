---
title: "Early Impressions of Rust from a Go Programmer"
date: 2020-02-27T02:50:00-07:00
tags: ["go", "rust"]
---

F$&k the borrow checker, or "How I learned to stop worrying and love the compiler".

<!--more-->

Even though most of my day-to-day work is in Python, I still consider myself a Go programmer. I co-wrote the SSH server powering Bitbucket.org in Go and maintain a number of nontrivial projects, such as an [easy to use SSH git server](https://github.com/belak/gitdir), an [IRC library](https://github.com/go-irc/irc), and an [IRC bot](https://github.com/belak/go-seabird).

After reading [Early Impressions of Go From a Rust Programmer](https://pingcap.com/blog/early-impressions-of-go-from-a-rust-programmer/), I started thinking about my experiences and how it's been for me learning in the other direction and wanted to see what a similar article would look like.

## Why Rust?

One of my favorite projects for learning a new language is writing an IRC bot. It handles network IO, text parsing, and potentially dispatching of some type. For me, this is usually the perfect level of complexity for playing around and learning, but complex enough to expand in the future. My current bot project has been alive for quite some time. It's survived a few additional rewrites (mostly python with asyncio) and I've stuck with this one. Rust seemed like it might be a good fit and I had been looking for a lower level language that was easier to grok than C++ but more powerful than C.

At an initial glance, the features of Rust seemed to line up well with things I consider pain points in Go, so I figured I'd give it a shot.

## Learnability

Rust is not a simple language. It has generics, traits, macros, async/await, and a borrow checker (among many other things). In particular that last one was hard to learn. I have a number of friends who have heard me rant endlessly while trying to figure out issues with ownership and the borrow checker, but at some point it just started making more sense. Unfortunately I don't have better advice other than "it doesn't make sense until it does". It's a useful property of the language which was initially hard to understand.

Generics, traits, and macros all had a similar learning curve for me. Lots and lots of initial pain followed by eventual understanding. For me, Rust has been worth the pain, but it definitely won't be for everyone.

## Things I Like

Rust has so many conveniences and niceties that come from a newer, more complex language. Many of these are direct improvements on things I consider pain points in go.

- First class `enum` support. Being able to match on an enum variant is really nice, especially because you can include different values with different variants.
- `Result` and `Option`. Being able to propagate errors using `?` rather than `if err != nil { return err }` is really nice. There are some oddities around conversions between error types (and having to use an external crate like `anyhow`, `failure`, or `snafu` to get a truely generic error type), but having the convenience in propagation makes that easy to look past.
- Generics. This has become a contentious topic in the Go community. I had resigned myself to never getting them in Go because at least for me, interfaces are generally good enough. Rust does some pretty impressive things with them (allowing generic conversions by implementing `From`, extension traits, etc) and they've gone a long way to showing me what I was missing with Go.
- Explicit interfaces. In Go, you cannot have multiple functions with the same name. Because of how this interacts with implicit interfaces, if two interfaces you need to implement require a method with the same name, it gets hard to manage. Normally this is not an issue, but I also like being explicit about things.
- Explicit visibility and privacy. It's really nice to know that only exported functions are going to be exported. In Go the upper case for exported, lower case for non-exported seems like a strange design decision.
- Matching on multiple things at once. This is more of a syntactic trick, but being able to match on a tuple and have each of the match arms destructure the tuple is really nice.
- Package management. The clustercuss of Go package management (vendoring, vgo, dep, modules, etc) has long been a pain point. Cargo was a breath of fresh air. There are some oddities around abandoned crates with nice top-level names, but overall it's been a pleasure to work with and for the most part just works.
- Multiple iterators can be implemented on a single type and you're not limited to only iterating over maps and slices. Strings are a good example which have both `.chars()` and `.bytes()`.

## Things I'm Not Completely Sold On

In addition to things that were definitely positive, there were also some that were frustrating, but have a clear reason.

- I have no doubt Async/Await will improve in the future, but it has been a fairly big pain point so far. There are multiple async runtimes, which makes it hard to write portable async code. Additionally, if you need to run some blocking code, like an ORM it seems like the current recommendation is to submit that as a task on another thread pool. But then you need to think about what sorts of thread pools you have going on, resource constraints between them, etc. Plus, if you do it wrong, you could still manage to block tasks in the main thread pool. In Go, you start a goroutine and it just works. Async is fairly new in Rust, so I look forward to seeing this get easier to use in the future.
- Switches must be exhaustive. Because one of my main personal projects deals with strings, I'm often dealing with matching against them. With IRC, you only have a small number of message types you probably want to match on, but Rust enforces you to cover all cases.
- It's been a long time since I've had to think about actual memory management, so having to wrap everything in an `Arc` (and make sure it's owned rather than borrowed) in order to make it work with async has been frustrating. Go lets you be fairly loose about memory (`&SomeStruct{}` will return a pointer to that struct, but because of lifetimes, this is not really possible in Rust).
- Lifetimes. In particular, it is hard to make a type which can be used inside a lifetime without additional allocations but can still escape a lifetime. I wanted to make an IRC message type which could be read and used inside an async task, but if you want it to survive outside that lifetime, you need to use `Cow<str>` for every string type which gets painful.
- `String`/`&str`/`Cow<str>`. There are places for all of these, but understanding where they fit in has taken a lot of time and experimentation.

## Things I Dislike

- Compile times are definitely much slower than Go. It's a trade-off because generics, macros and a complex borrow checker can't be implemented without a cost. I do worry a bit that this will start veering towards the long compilation times of C++ because of complex macros, but I'd be happy to be proved wrong.
- Compilation errors. Sometimes these are indecipherable with tons of nested types when the end reason was that you forgot to import a Trait somewhere. Sometimes this is unavoidable - generics do not lend themselves to simple compile errors. The messages have been getting better, but they've got a long way to go.
- Self-contained types. You cannot have an IRC Message which contains its own byte buffer with references to components of the message. This can be worked around by storing ranges rather than the actual references and creating the references in functions when needed, but it's not convenient.
- No varargs. The only way to implement something in the same pattern as `format!` is to make a macro, which can't really be attached to a type. You can cheat and use `format_args!` but this isn't a clean solution.

## Takeaways

Both languages can learn quite a lot from each other. Rust's enums, `Result`, and `Option` types would work amazingly in Go. Additionally, more powerful built-in async support would be incredible in Rust. At this point I think both languages have their place, but I think Rust fits better for lower level, resource constrained or performance critical projects, while Go is much easier for higher level projects.

Even though both languages could be thought of as replacements for C on a range of C to Python, Go is closer to the Python side (convenience while maintaining good performance) and Rust is closer to the C side (performance and safety at the expense of convenience). Both languages have value, and each comes with a separate set of trade-offs. In particular, I plan on eventually moving gitdir to Rust (as it has a higher expectation of security) but I'm still on the fence in regards to my IRC bot.

## Discussion

Want to talk more about this? There are discussion threads at the following locations:

- [lobste.rs](https://lobste.rs/s/gu3s9u/early_impressions_rust_from_go)
- [Hacker News](https://news.ycombinator.com/item?id=22438880)

Thanks for reading!
