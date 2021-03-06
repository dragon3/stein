# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

This file is automatically generated by [git-chglog](github.com/git-chglog/git-chglog). Don't edit by hand.

<a name="unreleased"></a>
## [Unreleased]


<a name="v0.2.11"></a>
## [v0.2.11] - 2019-12-10
### Pull Requests
- Merge pull request [#17](https://github.com/b4b4r07/stein/issues/17) from dragon3/fix-jsonpath-empty-string-value


<a name="v0.2.10"></a>
## [v0.2.10] - 2019-10-16
### Added
- **core:** Add tests for internal functions
- **core:** Add tests for internal functions

### Fixed
- **docs:** Update some wrong pages
- **docs:** Use edit_uri to edit pages on browser


<a name="v0.2.9"></a>
## [v0.2.9] - 2019-10-02
### Fixed
- **core:** Make jsonpath type detection strict

### Pull Requests
- Merge pull request [#16](https://github.com/b4b4r07/stein/issues/16) from b4b4r07/hugo-to-mkdocs


<a name="v0.2.8"></a>
## [v0.2.8] - 2019-09-27
### Added
- **core:** Add tests for jsonpath

### Fixed
- **core:** Use ctyjson.Unmarshal for jsonpath internal logic to detect typo more strictly


<a name="v0.2.7"></a>
## [v0.2.7] - 2019-09-20
### Fixed
- **core:** fix [#12](https://github.com/b4b4r07/stein/issues/12) fundamentally

### Reverts
- GJSON handle parsed arg as just string everytime

### Pull Requests
- Merge pull request [#15](https://github.com/b4b4r07/stein/issues/15) from b4b4r07/fix-14


<a name="v0.2.6"></a>
## [v0.2.6] - 2019-08-29

<a name="v0.2.5"></a>
## [v0.2.5] - 2019-08-27
### Pull Requests
- Merge pull request [#13](https://github.com/b4b4r07/stein/issues/13) from b4b4r07/fix-12
- Merge pull request [#8](https://github.com/b4b4r07/stein/issues/8) from b4b4r07/use-release-go
- Merge pull request [#7](https://github.com/b4b4r07/stein/issues/7) from b4b4r07/techdoc


<a name="v0.2.4"></a>
## [v0.2.4] - 2019-02-08
### Fixed
- **core:** Fix bug related to the file name cache parsed as HCL (to show source code when facing some errors)


<a name="v0.2.3"></a>
## [v0.2.3] - 2019-02-06
### Added
- **core:** Add log mechanism for debugging like Terraform
- **docs:** Add more docs and update
- **func:** Add exist built-in function

### Pull Requests
- Merge pull request [#6](https://github.com/b4b4r07/stein/issues/6) from b4b4r07/log-design


<a name="v0.2.2"></a>
## [v0.2.2] - 2019-01-29
### Changed
- **core:** Rename "expressions" in rule attribute name to "conditions"

### Code Refactoring
- **core:** Make topological sort package to one for lint internal
- **core:** Fix import path
- **core:** Move lang to lint/internal/policy

### Pull Requests
- Merge pull request [#5](https://github.com/b4b4r07/stein/issues/5) from b4b4r07/expressions-to-conditions
- Merge pull request [#4](https://github.com/b4b4r07/stein/issues/4) from b4b4r07/refactor-lang

### BREAKING CHANGE

We introduced "precondition" block before, similarly to "conditions".
It means pre-conditions. So "conditions" is more easier to understand
and write policies than "expressions".


<a name="v0.2.1"></a>
## [v0.2.1] - 2019-01-24
### Pull Requests
- Merge pull request [#3](https://github.com/b4b4r07/stein/issues/3) from b4b4r07/introduce-policy-dir


<a name="v0.2.0"></a>
## [v0.2.0] - 2019-01-21
### Pull Requests
- Merge pull request [#2](https://github.com/b4b4r07/stein/issues/2) from b4b4r07/precondition
- Merge pull request [#1](https://github.com/b4b4r07/stein/issues/1) from b4b4r07/jsonpath-default


<a name="v0.1.0"></a>
## v0.1.0 - 2019-01-07

[Unreleased]: https://github.com/b4b4r07/stein/compare/v0.2.11...HEAD
[v0.2.11]: https://github.com/b4b4r07/stein/compare/v0.2.10...v0.2.11
[v0.2.10]: https://github.com/b4b4r07/stein/compare/v0.2.9...v0.2.10
[v0.2.9]: https://github.com/b4b4r07/stein/compare/v0.2.8...v0.2.9
[v0.2.8]: https://github.com/b4b4r07/stein/compare/v0.2.7...v0.2.8
[v0.2.7]: https://github.com/b4b4r07/stein/compare/v0.2.6...v0.2.7
[v0.2.6]: https://github.com/b4b4r07/stein/compare/v0.2.5...v0.2.6
[v0.2.5]: https://github.com/b4b4r07/stein/compare/v0.2.4...v0.2.5
[v0.2.4]: https://github.com/b4b4r07/stein/compare/v0.2.3...v0.2.4
[v0.2.3]: https://github.com/b4b4r07/stein/compare/v0.2.2...v0.2.3
[v0.2.2]: https://github.com/b4b4r07/stein/compare/v0.2.1...v0.2.2
[v0.2.1]: https://github.com/b4b4r07/stein/compare/v0.2.0...v0.2.1
[v0.2.0]: https://github.com/b4b4r07/stein/compare/v0.1.0...v0.2.0
