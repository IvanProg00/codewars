pub fn run() {
    println!("{}", reverse_words("hello world!"));
}

fn reverse_words(str: &str) -> String {
    str.split(' ').rev().collect::<Vec<_>>().join(" ")
}
