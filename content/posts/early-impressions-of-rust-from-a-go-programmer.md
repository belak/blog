---
title: "Early Impressions of Rust from a Go Programmer"
date: 2020-02-26T03:30:00-07:00
tags: ["go", "rust"]
hide_updated: true
draft: true
---

F$&k the borrow checker, or "How I learned to stop worrying and love the compiler".

<!--more-->

After reading [Early Impressions of Go From a Rust Programmer](https://pingcap.com/blog/early-impressions-of-go-from-a-rust-programmer/), I started thinking about my experiences and how it's been for me in the other direction. In the past I've worked on a number of large Go projects and have been enjoying my time with Rust.

## My Go Projects

Even though most of my day-to-day work is in Python, I still consider myself a Go programmer. I co-wrote the SSH server powering Bitbucket.org in Go and maintain a number of nontrivial projects, such as an [easy to use SSH git server](https://github.com/belak/gitdir), an [IRC library](https://github.com/go-irc/irc), and an [IRC bot](https://github.com/belak/go-seabird).

## Why Rust?

One of my favorite projects for learning a new language is writing an IRC bot. My current bot project has been alive for quite some time. It's survived a few additional rewrites (mostly python with asyncio) and I've stuck with this one. Rust seemed like it might be a good fit and I've been looking for a lower level language that's easier to grok than C++ but more powerful than C.

At an initial glance, the features of Rust seem to line up well with things I consider pain points in Go, so I figured I'd give it a shot.

## Things I Like

- First class `enum` support. Being able to match on an enum variant is really nice, especially because you can include different values with different variants.
- `Result` and `Option`. Being able to propagate errors using `?` rather than `if err != nil { return err }` is really nice. There are some oddities around conversions between error types (and having to use an external crate like `anyhow`, `failure`, or `snafu` to get a truely generic error type), but having the convenience in propagation makes that easy to look past.
- Generics. This has become a contentious topic in the Go community. I had resigned myself to never getting them in Go because at least for me, interfaces are generally good enough. Rust does some pretty impressive things with them (allowing generic conversions by implementing `From`, extension traits, etc) and they've gone a long way to showing me what I was missing with Go.
- Explicit interfaces. In Go, you cannot have multiple functions with the same name. Because of how this interacts with implicit interfaces, if two interfaces you need to implement require a method with the same name, it gets very hard to manage. Normally this is not an issue, but I like being explicit about things.
- Explicit visibility and privacy. It's really nice to know that only exported functions are going to be exported. In Go the upper case for exported, lower case for non-exported seems like a very strange design decision.
- Matching on multiple things at once. This is more of a syntactic trick, but being able to match on a tuple and have each of the match arms destructure the tuple is really nice.
- Package management. The clustercuss of Go package management (vendoring, vgo, dep, modules, etc) has long been a pain point. Cargo was a breath of fresh air. There are some oddities around abandoned crates with nice top-level names, but overall it's been a pleasure to work with and for the most part just works.

## Things I'm Not Completely Sold On

I debated putting the borrow checker in here, but decided against it. I have a number of friends who have heard me rant endlessly while trying to figure out issues with ownership and the borrow checker, but at some point it started making more sense so while I wouldn't put it in "things I like", it's also not something that belongs here either. It's a useful property of the language which was initially hard to understand.

- Async/Await. I have no doubt this will improve in the future, but it has been a fairly big pain point so far. As an example, what do you do if you need to run some blocking code, like an ORM? It seems like the current recommendation is to submit that as a task on another thread pool. But then you need to think about what sorts of thread pools you have going on, resource constraints between them, etc. Plus, if you do it wrong, you could still manage to block the main thread pool. In go, you fire off a goroutine and it just works. Additionally, there are multiple async runtimes, which makes it hard to write portable async code.
- Switches must be exhaustive. Because one of my main personal projects deals with strings, I'm often dealing with matching against them. With IRC, you only have a small number of message types you probably want to match on, but Rust enforces the default case.
- It's been a long while since I've had to think about actual memory management, so having to wrap everything in an Arc in order to make it work with async was frustrating. Go lets you be very loose about memory (`&SomeStruct{}` will return a pointer to that struct, but because of lifetimes, this is not really possible in rust).
- Lifetimes. In particular, it is very hard to make a type which can escape a lifetime. I wanted to make an IRC message type which could be read and used inside an async task, but if you want it to survive outside that lifetime, you need to use `Cow<str>` for every string type which gets painful.
- `String`/`&str`/`Cow<str>`. There are places for all of these, but understanding where they fit in has taken a lot of time and experimentation.

## Things I Dislike

- Compile times. Definitely much slower than Go.
- Compilation errors. Sometimes these are indecipherable with tons of nested types when the end reason was that you forgot to import a Trait somewhere. The messages have been getting better, but they've got a long way to go.
- No varargs. The only way to implement something in the same pattern as `format!` is to make a macro, which can't really be attached to a type. You can cheat and use `format_args!` but this isn't a clean solution.

## Takeaways

Both languages can learn quite a lot from each other. Rust's enums, `Result`, and `Option` types would work amazingly in Go. Additionally, more powerful built-in async support would be incredible in Rust. At this point I think both languages have their place, but I think Rust fits better for lower level, resource constrained or performance critical projects, while Go is much easier for higher level projects.

## Discussion

Want to talk more about this? There are discussion threads at the following places:

- [lobste.rs](#)
- [Hacker News](#)

Thanks for reading!
