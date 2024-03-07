# from memory_profiler import profile
import time

class Person:
    def __init__(self, name, birth_date, country):
        self.name = name
        self.birth_date = birth_date
        self.country = country

# @profile
def main():
    people = []
    n = int(input("Введите количество людей (N): "))
    
    for _ in range(n):
        name = input("Введите имя: ")
        birth_date = list(map(int, input("Введите дату рождения (через пробел): ").split()))
        country = input("Введите страну проживания: ")

        person = Person(name, birth_date, country)
        people.append(person)

    print("Все люди:", [vars(person) for person in people])

    search_value = input("Введите значение для поиска (например, 'Россия'): ")

    filtered_people = [vars(person) for person in people if person.country == search_value]
    print("Найденные люди:", filtered_people)

    people_sorted_by_name = sorted(people, key=lambda x: x.name)
    print("Список людей по имени:", [vars(person) for person in people_sorted_by_name])

    people_sorted_by_birth_date = sorted(people, key=lambda x: x.birth_date)
    print("Список людей по дате рождения:", [vars(person) for person in people_sorted_by_birth_date])

if __name__ == "__main__":
    start_time = time.time()
    main()
    elapsed_time = time.time() - start_time
    print(f"Время выполнения: {elapsed_time} сек.")
