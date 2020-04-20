# Redis pubsub demo

Redis 的 pubsub 是实时的，如果先执行 pub，再执行 sub，则消息是丢失的。
所以需要先启动 sub，再启动 pub

```sh
python sub_demo.py
```

新开一个 shell:

```sh
python pub_demo.py
```
