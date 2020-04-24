# Redis push&pop demo

```python
import redis

client = redis.Redis()

# Append to list
res = client.rpush("test_list_01", "hello")
print(res)
res = client.rpush("test_list_01", "world")
print(res)

# Show list
res = client.lrange("test_list_01", 0, -1)
print(res)

# Get from start
res = client.lpop("test_list_01")
print(res)

res = client.lpop("test_list_01")
print(res)
```

