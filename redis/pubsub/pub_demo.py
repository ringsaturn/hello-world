import redis

client = redis.Redis()

client.publish(channel="test-channel-01", message="hello")
client.publish(channel="test-channel-01", message="world")
