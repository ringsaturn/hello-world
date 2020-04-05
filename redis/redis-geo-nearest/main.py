import redis

redis = redis.Redis()

test_geo_location_data = (
    (116.3883, 39.9289, "test_geo_location_a")
    + (116.3581, 39.5646, "test_geo_location_b")
    + (116.3279, 40.0089, "test_geo_location_c")
)
redis.geoadd("test_location_member", *test_geo_location_data)

redis.geopos("test_location_member", "test_geo_location_a")
redis.geopos("test_geo_location_c")

resp = redis.georadiusbymember(
    "test_location_member",
    "test_geo_location_a",
    1000,
    withdist=True,
    withcoord=True,
    withhash=True,
)
print(resp)

resp = redis.georadiusbymember("test_location_member", "test_geo_location_b", 10)
print(resp)
