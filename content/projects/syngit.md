---
title: "Syngit"
description: "A git repo synchronization tool"
date: 2022-11-08T20:42:21-01:00
tags: [Python, Utility, Automation]
---

## [Find it on GitHub](https://github.com/RaphGL/Syngit)

<!-- ABOUT THE PROJECT -->

Syngit is a CLI and init service that let's you synchronize repositories across different clients (Github, Codeberg, Gitlab, etc) with a very simple interface and configuration file.

Built With:

- [Python](https://www.python.org/)
- [Poetry](https://python-poetry.org/)
- [AioHttp](https://docs.aiohttp.org/en/stable/)
- [TOML](https://github.com/uiri/toml)

## Usage

1. Create a `$HOME/.config/syngit.toml` file
2. Make something akin to this:

```toml
main_client = "github"

[github]
username = "RaphGL"

[codeberg]
username = "RaphGL"

[gitlab]
username = "RaphGL"
```

3. Enable the syngit service:

```sh
$ systemctl enable syngit --now
```