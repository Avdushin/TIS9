## Confugure testing

## Apps

- ### Rust

```bash
cd Rust
cargo run --release
# for testing
hyperfine target/release/Rust.exe  -i
```

- ### Go

```bash
cd Go
go run main.go
go tool pprof mem.prof
(pprof) top
(pprof) list main.main
(pprof) web
(pprof) quit

```
- ### Python

```bash
cd Python
pip install memory_profiler
python -m memory_profiler app.py
```

