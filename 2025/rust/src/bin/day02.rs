use std::io::BufRead;

fn main() {
    let reader = aoc2025::new_reader();
    let x = reader
        .split(b',')
        .flat_map(|range| {
            let binding = String::from_utf8(range.unwrap()).unwrap();
            let (a, b) = binding.split_once('-').unwrap();
            a.parse::<i64>().unwrap()..=b.parse::<i64>().unwrap()
        })
        .map(|n| n.to_string())
        .filter(|n| n.len() % 2 == 0)
        .inspect(|n| println!("n: {n}"))
        .count();
    // .split(",")
    // .flat_map(|range| {
    //     let (a, b) = range.split_once('-').expect("Invalid range given");
    //     let start: i64 = a.parse().unwrap();
    //     let end: i64 = b.parse().unwrap();
    //     start..=end
    // });
}
