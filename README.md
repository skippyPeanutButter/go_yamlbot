# goyamlbot

A golang implementation of
[yamlbot](https://github.com/skippyPeanutButter/yaml_bot)


### Why

Mistakes can often be made when working with yaml-based configuration. Systems
such as travis-ci, rubocop, and other tools that utilize yaml files to govern
how they work often present users with a multitude of keys and options that can
take on many possible values. `yamlbot` allows users to feed it a set of rules
that a yaml-based system follows and then validate any yaml file against those
rules.

If a tool works off of a yaml-based configuration then a rules file can be
crafted for validating the configuration is correct. Create a rules file, share
it with others, and have them use `yamlbot` to validate their configuration
against the specified rules.

See the [RULES FILE](RULES_DEFINITION.md) specification for details on creating
a `.yamlbot.yml` rules file.
