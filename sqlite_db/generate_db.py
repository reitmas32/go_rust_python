import sqlite3
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

def create_table(cursor):
    cursor.execute('''CREATE TABLE IF NOT EXISTS users (
                        ID INTEGER PRIMARY KEY,
                        Name TEXT,
                        Age INTEGER,
                        Email TEXT,
                        City TEXT,
                        Country TEXT)''')

def insert_users(cursor, users):
    for user in users:
        cursor.execute('''INSERT INTO users (ID, Name, Age, Email, City, Country)
                          VALUES (?, ?, ?, ?, ?, ?)''', (user.id, user.name, user.age, user.email, user.city, user.country))

if __name__ == "__main__":
    num_users = 1000000
    users = generate_users(num_users)

    # Conectar a la base de datos SQLite
    conn = sqlite3.connect('users.db')
    cursor = conn.cursor()

    # Crear la tabla si no existe
    create_table(cursor)

    # Insertar los usuarios en la base de datos
    insert_users(cursor, users)

    # Confirmar los cambios y cerrar la conexión
    conn.commit()
    conn.close()

    print("Datos insertados en la base de datos SQLite.")
