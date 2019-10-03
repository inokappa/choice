# choice

```
$ echo 'foo,bar,baz' | choice
bar
$ echo 'foo|bar|baz' | choice --separate='|'
baz
```
