use std::{env, fs::File, io::BufReader};

pub fn new_reader() -> BufReader<File> {
    let filename = env::args()
        .nth(1)
        .expect("Porivde filename as first argument");
    let file = File::open(filename).unwrap();
    BufReader::new(file)
}
