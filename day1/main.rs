use std::fs;

fn main() {

  let contents = fs::read_to_string("input.txt")
      .expect("Should have been able to read the file");
  let x = contents.split('\n');
  let mut sum: i128 = 0;
  
  for str in x{
    let arr: Vec<char> = str.chars().filter(|f| f.is_numeric()).collect();
    let mut s= String::from(arr.first().unwrap().to_string() + & arr.last().unwrap().to_string());

    let z: i128 =s.parse().unwrap();
    sum += z;  
  }
   println!("{sum}");
}
