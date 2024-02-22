import csv
import random
import string

class User:
    def __init__(self, id, name, age, email, city, country):
        self.id = id
        self.name = name
        self.age = age
        self.email = email
        self.city = city
        self.country = country

def generate_random_string(length):
    letters = string.ascii_lowercase
    return ''.join(random.choice(letters) for _ in range(length))

def generate_users(num_users):
    users = []
    for i in range(1, num_users + 1):
        name = generate_random_string(8)
        age = random.randint(18, 65)
        email = f"{name}@example.com"
        city = generate_random_string(6)
        country = generate_random_string(5)
        users.append(User(i, name, age, email, city, country))
    return users

def write_csv(filename, users):
    with open(filename, 'w', newline='') as csvfile:
        fieldnames = ['ID', 'Name', 'Age', 'Email', 'City', 'Country']
        writer = csv.DictWriter(csvfile, fieldnames=fieldnames)
        
        writer.writeheader()
        for user in users:
            writer.writerow({'ID': user.id, 'Name': user.name, 'Age': user.age, 'Email': user.email, 'City': user.city, 'Country': user.country})

if __name__ == "__main__":
    num_users = 10000
    users = generate_users(num_users)
    write_csv('users.csv', users)
    print("CSV file generated successfully.")
