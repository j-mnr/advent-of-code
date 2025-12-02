use std::io::BufRead;

fn parse_line(line: &str) -> i32 {
    let (dir, rest) = line.split_at(1);
    let n: i32 = rest.parse().expect("Invalid number");
    match dir {
        "L" => -n,
        "R" => n,
        _ => panic!("Invalid line: {}", line),
    }
}

fn turn(a: i32, b: i32) -> i32 {
    (a + b).rem_euclid(100)
}

struct ByOnes {
    step: i32,
    remaining: usize,
}

impl Iterator for ByOnes {
    type Item = i32;
    fn next(&mut self) -> Option<i32> {
        if self.remaining == 0 {
            None
        } else {
            self.remaining -= 1;
            Some(self.step)
        }
    }
}

fn by_ones(x: i32) -> ByOnes {
    ByOnes {
        step: if x < 0 { -1 } else { 1 },
        remaining: x.abs() as usize,
    }
}

fn main() {
    let reader = aoc2025::new_reader();
    let (_, c1, _, c2) = reader.lines().map(|l| parse_line(&l.unwrap())).fold(
        (50, 0, 50, 0),
        |(pos1, count1, pos2, count2), rot| {
            let new_pos1 = turn(pos1, rot);
            let new_count1 = count1 + (new_pos1 == 0) as i32;
            let (new_pos2, new_count2) = by_ones(rot).fold((pos2, count2), |(p, c), step| {
                let np = turn(p, step);
                (np, c + (np == 0) as i32)
            });
            (new_pos1, new_count1, new_pos2, new_count2)
        },
    );
    println!("Part 1: {}", c1);
    println!("Part 2: {}", c2);
}
