#!/bin/bash

# Сборка программы
cargo build --release

# Измерение времени выполнения и использования памяти
hyperfine --prepare "cargo build --release" 'target/release/Rust.exe'
