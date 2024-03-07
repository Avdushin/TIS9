#!/bin/bash

# Сборка программы
cargo build --release

# Измерение времени выполнения и использования памяти
hyperfine target/release/Rust.exe  -i
