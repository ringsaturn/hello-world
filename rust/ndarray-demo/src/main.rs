use ndarray::*;
use ndarray_linalg::*;
// use num_traits::Float;

fn compute_distance_demo() {
    let point_a = array![[0.,0.]];
    let point_b = array![[4.,3.]];
    println!("{}", (point_a - point_b).norm_l2());
}

fn main() {
    compute_distance_demo();
}
