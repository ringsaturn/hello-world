# Redis pubsub demo

Redis 的 pubsub 是实时的，如果先执行 pub，再执行 sub，则消息是丢失的。
所以需要先启动 sub，再启动 pub

Pub:

```py
import redis
import time

client = redis.Redis()

p = client.pubsub()
p.psubscribe("test-channel-01")
while True:
    message = p.get_message()
    if message:
        print(message)
    else:
        time.sleep(1)
```

Sub:

```py
import redis

client = redis.Redis()

client.publish(channel="test-channel-01", message="hello")
client.publish(channel="test-channel-01", message="world")
```

