# sonic-fuzz

Fuzzing test for [sonic](https://github.com/bytedance/sonic)

# Getting Started

## Install dependency tools and modules

```
make install
```

## Run fuzzing test for stdjson

This step is to create many corpus for sonic.
```
make stdjson
```

## Run fuzzing test for sonic

The problems found in fuzzing are stored at wordir/crashers.
```
make sonic
```
