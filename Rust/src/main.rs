use std::io;

#[derive(Debug, PartialOrd, Ord, PartialEq, Eq)]
struct Person {
    name: String,
    birth_date: [u32; 3],
    country: String,
}

fn main() {
    let mut people: Vec<Person> = Vec::new();
    let mut n_str = String::new();

    println!("Введите количество людей (N):");
    io::stdin().read_line(&mut n_str).expect("Не удалось прочитать строку");
    let n: usize = match n_str.trim().parse() {
        Ok(num) => num,
        Err(_) => {
            println!("Некорректный ввод числа.");
            return;
        }
    };

    for _ in 0..n {
        println!("Введите имя: ");
        let mut name = String::new();
        io::stdin().read_line(&mut name).expect("Не удалось прочитать строку");
        let name = name.trim().to_string();

        println!("Введите дату рождения (через пробел):");
        let mut birth_date = String::new();
        io::stdin().read_line(&mut birth_date).expect("Не удалось прочитать строку");
        let birth_date: Vec<u32> = birth_date
            .trim()
            .split_whitespace()
            .map(|s| s.parse().expect("Ожидалось целое число"))
            .collect();

        if birth_date.len() != 3 {
            println!("Некорректный ввод даты рождения.");
            return;
        }

        println!("Введите страну проживания:");
        let mut country = String::new();
        io::stdin().read_line(&mut country).expect("Не удалось прочитать строку");
        let country = country.trim().to_string();

        let person = Person {
            name,
            birth_date: [birth_date[0], birth_date[1], birth_date[2]],
            country,
        };

        people.push(person);
    }

    println!("Все люди: {:?}", people);

    println!("Введите значение для поиска (например, 'Россия'):");
    let mut search_value = String::new();
    io::stdin().read_line(&mut search_value).expect("Не удалось прочитать строку");
    let search_value = search_value.trim();

    let filtered_people: Vec<&Person> = people
        .iter()
        .filter(|&person| person.country == search_value)
        .collect();

    println!("Найденные люди: {:?}", filtered_people);

    people.sort_by(|a, b| a.name.cmp(&b.name));
    println!("Список людей по имени: {:?}", people);

    people.sort_by(|a, b| a.birth_date.cmp(&b.birth_date));
    println!("Список людей по дате рождения: {:?}", people);
}
