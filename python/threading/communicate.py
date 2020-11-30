#!/usr/bin/env python3

import threading
import time
from queue import Queue

globalTaskQueue = Queue(maxsize=10)


def producer(taskQueue):
    while True:
        if taskQueue.empty():
            taskQueue.put(int(time.time()))
        time.sleep(0.1)


def consumer(taskQueue, consumer_id: int):
    while True:
        if taskQueue.empty():
            time.sleep(1)
            continue
        data = taskQueue.get()
        print(consumer_id, data)
        time.sleep(0.2)


t0 = threading.Thread(target=producer, args=(globalTaskQueue,))
t1 = threading.Thread(target=consumer, args=(globalTaskQueue, 1,), daemon=True)
t2 = threading.Thread(target=consumer, args=(globalTaskQueue, 2,), daemon=True)
t3 = threading.Thread(target=consumer, args=(globalTaskQueue, 3,), daemon=True)

t1.start()
t2.start()
t3.start()

t0.start()
