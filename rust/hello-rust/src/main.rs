
fn compute_l2_norm_two_points(x1: f64, y1:f64, x2:f64, y2:f64) -> f64{
    let distance_square = (x1-x2)*(x1-x2) + (y1-y2)*(y1-y2);
    let distance = distance_square.sqrt();
    return distance;
}

fn two_points_demo(){
    let x1: f64 = 119.0;
    let y1: f64 = 39.0;
    println!("(x1, y1): {},{}", x1, y1);

    let x2: f64 = 120.0;
    let y2: f64 = 40.0;
    println!("(x2, y2): {},{}", x2, y2);

    let distance = compute_l2_norm_two_points(x1, y1, x2, y2);
    println!("distance={}", distance);
}

fn main(){
    two_points_demo();
}

