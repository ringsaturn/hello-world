"""Compute nearest point.

References:

- https://stackoverflow.com/a/19794741/6713916
- https://github.com/numpy/numpy/pull/3387
"""

import numpy as np

points = np.array(
    [[114.39, 39.490], [135.98, 12.31], [98.12, 12.98], [50, 60], [119, 45]]
)
target = np.array([45.67, 87.98])

distance_list = np.linalg.norm(points - target, axis=1)
print(distance_list)
min_index = np.argmin(distance_list)
print(min_index)
