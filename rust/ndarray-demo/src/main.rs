use ndarray::*;
use ndarray_linalg::*;
// use num_traits::Float;

fn main() {
    let a1 = Array::range(1., 10., 1.);
    println!("{}", a1);

    let a2 = arr1(&[1.0, 1.0, -3.0]);
    println!("a2.norm_l2: {}", a2.norm_l2());
    println!("{} {} {}", a2.norm_max(), 3.0, 1e-7);
    let b2 = arr2(&[[1.0, 3.0], [1.0, -4.0]]);
    println!("{} {} {}", b2.norm_max(), 4.0, 1e-7);
}
