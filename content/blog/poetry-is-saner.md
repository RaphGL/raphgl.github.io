---
title: Poetry is a saner package manager for Python
date: 2022-10-10T12:13:10-01:00
---

Like many, I enjoy Python and the versatility and speed of development it provides. But, to me Python tooling is a big deal breaker. It's insane, there's way too many ways to do things. Not to mention that documentation in the official python website is kind of convoluted, you end up going through hoops to understand whatever you need or relying on other people's tutorials, videos, etc, to understand them, even then it can be overwhelming the amount of choices there are in python. For python bundling you have: wheels, pyinstaller, py2exe. For configuration you have: setup.py, setup.cfg, pyproject.toml. Then you need to configure your environment: pyenv, pipenv, virtualenv. This goes on and on and the recommended approach seems to change every few years and we get a conflict between all of these being used in production.

For people that have come from other languages, this seems pure insanity and learning that you actually need all of these to actually spin up your project for prodution because python was never that well thought out to work in this way is kind of insane. I truly wished python had sane tooling and a de facto way of doing things, unfortunately that isn't the case and will continue to be like this for a while, thankfully (for me at least) there's a solution.

Someone recommended me to look at Poetry. Poetry is a third party package manager that's better than Pip and actually does what it's supposed to do, it wraps around all the insanity in python and let's you focus on what's important, packaging and developing software.

The advantages of using Poetry:

- Saner environment
- The virtual environment is dealt with for you
- Dependency management is actually enjoyable and you can pin and track dependencies' versions
- Building and publishing projects is a joy

## Getting started with Poetry

To initialize Poetry:

```sh
$ poetry new your_project
```

You'll get something like:

```sh
your_project
├── pyproject.toml
├── README.rst
├── tests
│   ├── __init__.py
│   └── test_your_project.py
└── your_project
    └── __init__.py
```

There there is a pyproject.toml that you can add some metadata about your project. If you try and install a program:
```sh
$ poetry add requests
```
You'll see that poetry adds a virtual environment and installs your dependency. If you look into the `pyproject.toml` file you should find that a dependency was added. To run your code you can either just enter the virtual environment (like you would normally do in normal python) by running `poetry shell` or you could ask poetry to run the commands for you like so:
```sh
$ poetry run python main.py
```

There's many cool things that poetry can do that makes managing a python project much easier but that are beyond the scope of this blog. Make sure to check out poetry [here](https://python-poetry.org/docs/basic-usage/).
