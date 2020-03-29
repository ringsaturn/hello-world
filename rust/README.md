# Rust Notes

rust-numpy, rust-ndarray, orjson 是有机会改进含有大量 NumPy 类型 Python 程序性能的方法，
问题的瓶颈往往不在 NumPy 的部分，而是需要利用 NumPy 计算结果的 Python 代码中。
如果可以将 Python 代码的逻辑移植到 Rust 上， 配合 Rust-NumPy，也许能改善部分程序性能。

## Setup

```sh
# [rust-numpy] 需要 Python 环境
# 进入一个有 Python3 的虚拟环境
# 为了偷懒可以 Symlinks 到一个已有的环境中, 但是不要在真的「实际开发」中这么弄
python3 -m venv venv
. venv/bin/activate
pip install numpy

# ZSH 似乎和 cargo 不兼容，加到 zshrc 里也没有生效
# 需要手工执行 export
export PATH="$HOME/.cargo/bin:$PATH"

# [orjson, PyO3] 需要 nighly 版本的 Rust
rustup install nightly
rustup override set nightly

# 运行 demo
cargo run
```

## References

* https://doc.rust-lang.org/book/appendix-07-nightly-rust.html#rustup-and-the-role-of-rust-nightly
* https://pyo3.rs/v0.9.0-alpha.1/
* https://github.com/PyO3/pyo3
* https://github.com/PyO3/rust-numpy
* https://github.com/rust-ndarray/ndarray
