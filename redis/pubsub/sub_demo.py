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
