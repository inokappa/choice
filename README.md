# choice
[![CircleCI](https://circleci.com/gh/inokappa/choice.svg?style=svg)](https://circleci.com/gh/inokappa/choice)

```
$ echo 'foo,bar,baz' | choice
bar
$ echo 'foo|bar|baz' | choice --separate='|'
baz
```
