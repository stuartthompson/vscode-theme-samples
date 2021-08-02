#[allow(dead_code)]
struct Thing<'a> {
    name: &'a str
}

fn main() {
    let mine = Thing { name: 'test' };
    let other = vec![1, 2, 3];

    // Say Hello
    println!("Hello, World!");
}




