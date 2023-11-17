# Kaniko Build Args
Takes a json key-value object and returns a string that can be used in Kaniko build.

This was built to dynamically add data to kaniko when building images in CI pipeline.

## Usage
```
$ kaniko-build-args --json '{"my_variable": "my_value", "my_variable_two": "my_value_two"}'
--build-arg my_variable=my_value --build-arg my_variable_two=my_value_two
```
