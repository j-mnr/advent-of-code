use std::io::BufRead;

fn main() {
    let part1 = aoc2025::new_reader()
        .lines()
        .map(|line| {
            let binding = line.unwrap();
            let s = binding.trim();
            let mut first: u32 = 0;
            let mut fidx: usize = 0;
            let mut it = s.char_indices().peekable();
            while let Some((i, c)) = it.next() {
                let d = c.to_digit(10).unwrap();
                if d > first && !it.peek().is_none() {
                    first = d;
                    fidx = i;
                }
            }

            let mut second: u32 = 0;
            for (_, c) in s.char_indices().skip_while(|(i, _)| i <= &fidx) {
                let d = c.to_digit(10).unwrap();
                if d > second {
                    second = d
                }
            }
            first * 10 + second
        })
        .fold(0, |sum, n| sum + n);
    println!("Part 1: {part1}")
}
