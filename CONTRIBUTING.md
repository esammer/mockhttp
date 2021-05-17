# Contributor Guide

We welcome community contributions to mockhttp. While we try and make it as simple as possible for folks to get
involved, we do have a few standards for contributions. This project has an accepted and documented code of conduct that
we all contributors and community participants to follow.

This document uses RFC-style definitions of _must_, _may_, and _should_.

## Requesting a Feature

Anyone may request a feature by opening an issue. The more information you provide, the better. Feature requests that
come with PRs are obviously even better.

There may be features that do not align with the general principles of this project or its maintainers' view of what
is "best." If we decide that a particular feature request falls into this category, we'll do our best to let you know
why we feel that way. We appreciate your understanding.

## Reporting an Issue

Bug reports are welcome. If you're not sure if something is a bug, please file it, and we'll do our best to validate it.
More information is typically more helpful than less.

## Contributing

Contributors _must_ have the legal right to make contributions to the project. That is, contributions must be
unencumbered by employment contracts, patents, other restrictive software licenses, or governmental regulations. All
contributions must be licensable under this project's license. See LICENSE.

Contributions _must_ follow the [Standards](#Standards) below.

Contributions that do not align with the general principles of this project or its maintainers' views of what is "best"
for the project may be declined. If we decide that a particular contribution falls into this category, we'll do our best
to let you know why we feel that way. We appreciate your understanding.

## Standards

### Code

We follow the standard Go formatting rules. Code _must_ be formatted with `go fmt`.

This project follows semver and typical Go mod standards with respect to versioning. Backward incompatible API changes
_must not_ be made unless the major version is incremented.

All structs and methods that are part of the public API _must_ be documented with Go doc. In cases where its helpful,
examples methods _may_ be included.

Code _must_ have associated unit tests.

If tradeoffs must be made, code _should_ prioritize public API usability > internal readability > testability >
performance > write-time efficiency unless there is a good reason.

### Commits

Commits _must_ have a descriptive 1 line summary, and _may_ have additional commentary describing the change, tradeoffs,
alternatives, and potential issues. While we understand that "descriptive" can be subjective, "typo," "bug fix," "
update" are never considered descriptive. When in doubt, see `git log` for examples.

Every commit _must_ leave the `main` branch in a releasable state. We _must_ be able to release at any point in
committed history, including retroactively.

Commits on `main` _must_ be linear. Merge commits are not allowed.

Each commit _should_ be a discrete, atomic feature, fix, or refactor. This makes bugs easier to detect and revert,
minimizes merge conflicts, and makes project state easier to understand.

### Issues

Issues _must- have a descriptive 1 line summary, and _may_ have additional commentary describing the feature request,
issue, tradeoffs, alternatives, reproduction steps, and potential issues. While we understand that "descriptive" can be
subjective, "typo," "bug fix," "update" are never considered descriptive. When in doubt, see existing issues for
examples.

Issues _should_ have the appropriate labels attached to them, and follow existing patterns.
